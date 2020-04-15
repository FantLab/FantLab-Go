package server

import (
	"context"
	"errors"
	"expvar"
	"fantlab/base/anyserver"
	"fantlab/base/codeflow"
	"fantlab/base/dbtools"
	"fantlab/base/dbtools/sqldb"
	"fantlab/base/dbtools/sqlr"
	"fantlab/base/edsign"
	"fantlab/base/httprouter"
	"fantlab/base/memcacheclient"
	"fantlab/base/redisclient"
	"fantlab/docs"
	"fantlab/server/internal/app"
	"fantlab/server/internal/config"
	"fantlab/server/internal/logs"
	"fantlab/server/internal/routes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/pprof"
	"os"
	"time"

	"go.uber.org/zap"

	"github.com/gomodule/redigo/redis"
	"github.com/minio/minio-go/v6"
	"go.elastic.co/apm/module/apmredigo"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/mysql"
)

func GenerateDocs() {
	_ = docs.Generate(os.Stdout, routes.Tree(nil, nil, nil), "/"+routes.BasePath)
}

func Start() {
	apiServer := makeAPIServer()

	var monitoringServer *anyserver.Server

	if apiServer.SetupError == nil {
		monitoringServer = makeMonitoringServer()
	}

	anyserver.RunWithGracefulShutdown(func(err error) {
		logs.Logger().Error(err.Error())
	}, apiServer, monitoringServer)

	time.Sleep(1 * time.Second)
}

func makeAPIServer() (server *anyserver.Server) {
	server = new(anyserver.Server)

	var mysqlDB sqlr.DB
	var redisClient redisclient.Client
	var memcacheClient memcacheclient.Client
	var cryptoCoder *edsign.Coder
	var minioClient *minio.Client
	var minioBucket string
	var appConfig *config.AppConfig

	server.SetupError = codeflow.Try(
		func() error { // мускуль
			db, err := apmsql.Open("mysql", os.Getenv("MYSQL_URL"))
			if err == nil {
				err = db.Ping()
			}
			if err != nil {
				return fmt.Errorf("MySQL setup error: %v", err)
			}

			server.DisposeBag = append(server.DisposeBag, db.Close)

			mysqlDB = sqlr.Log(sqldb.New(db), func(ctx context.Context, entry sqlr.LogEntry) {
				logger := logs.WithAPM(ctx)
				logger.Info(
					entry.Query(),
					zap.Duration("duration", entry.Duration),
					zap.Int64("rows", entry.Rows),
				)
				if entry.Err != nil && !dbtools.IsNotFoundError(entry.Err) {
					logger.Error(entry.Err.Error())
				}
			})

			return nil
		},
		func() error { // редис (опционально)
			serverAddr := os.Getenv("RDS_ADDRESS")
			if len(serverAddr) == 0 {
				return nil
			}

			client, close := redisclient.NewPoolClient(serverAddr, 8, func(pool *redis.Pool, ctx context.Context) (redis.Conn, error) {
				return apmredigo.Wrap(pool.Get()).WithContext(ctx), nil
			})
			err := client.Perform(context.Background(), func(conn redisclient.Conn) error {
				_, err := conn.Do("PING")
				return err
			})
			if err != nil {
				return fmt.Errorf("Redis setup error: %v", err)
			}

			server.DisposeBag = append(server.DisposeBag, close)

			redisClient = redisclient.Log(client, func(ctx context.Context, err error) {
				logs.WithAPM(ctx).Error(err.Error())
			})

			return nil
		},
		func() error { // мемкэш (опционально)
			serverAddr := os.Getenv("MC_ADDRESS")
			if len(serverAddr) == 0 {
				return nil
			}

			client := memcacheclient.New(serverAddr)
			err := client.Ping()
			if err != nil {
				return fmt.Errorf("Memcache setup error: %v", err)
			}

			memcacheClient = memcacheclient.APM(memcacheclient.Log(client, func(ctx context.Context, entry *memcacheclient.LogEntry) {
				logger := logs.WithAPM(ctx)
				logger.Info(fmt.Sprintf("%s %s", entry.Operation, entry.Key))
				if entry.Err != nil && !memcacheclient.IsNotFoundError(entry.Err) {
					logger.Error(entry.Err.Error())
				}
			}), "db.memcache")

			return nil
		},
		func() error { // криптокодер для jwt-like токенов
			coder, err := edsign.NewFileCoder64(os.Getenv("JWT_PUBLIC_KEY_FILE"), os.Getenv("JWT_PRIVATE_KEY_FILE"))
			if err != nil {
				coder, err = edsign.NewCoder64(os.Getenv("JWT_PUB_KEY"), os.Getenv("JWT_PRIV_KEY"))
				if err != nil {
					return fmt.Errorf("JWT setup error: %v", err)
				}
			}
			cryptoCoder = coder
			return nil
		},
		func() error { // minio
			server := os.Getenv("MINIO_SERVER")
			if len(server) == 0 {
				return errors.New("Min.io setup error: server not found")
			}

			client, err := makeMinioWithSecrets(server, os.Getenv("MINIO_ACCESS_KEY_FILE"), os.Getenv("MINIO_SECRET_KEY_FILE"))
			if err != nil {
				client, err = minio.New(server, os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), false)
				if err != nil {
					return fmt.Errorf("Min.io setup error: %v", err)
				}
			}

			minioBucket = os.Getenv("MINIO_BUCKET")
			bucketExists, err := client.BucketExists(minioBucket)
			if err != nil {
				return fmt.Errorf("Min.io setup error: %v", err)
			}
			if !bucketExists {
				return fmt.Errorf("Min.io setup error: bucket %s not found", minioBucket)
			}

			minioClient = client

			return nil
		},
		func() error { // конфигурация бизнес-логики
			// Все параметры заданы в config/main.cfg и config/misc.cfg Perl-бэка
			appConfig = &config.AppConfig{
				ImagesBaseURL:         os.Getenv("IMAGES_BASE_URL"),
				ForumTopicsInPage:     20,
				ForumMessagesInPage:   20,
				MaxForumMessageLength: 20000,
				// В Perl-бэке указаны разные значения: при редактировании - 2_000с., там же в комментарии - 3_600с.,
				// при удалении - 1_800c. Остановимся на часе.
				MaxForumMessageEditTimeout: 3600,
				// Первоапрельские форумы, в отличие от Perl-бэка, недоступны для любых действий (поскольку доступ к ним
				// реализован хардкодом в Auth.pm)
				DefaultAccessToForums:     []uint64{1, 2, 3, 5, 6, 7, 8, 10, 12, 13, 14, 15, 16, 17, 22},
				BlogsInPage:               50,
				BlogTopicsInPage:          20,
				BlogArticleCommentsInPage: 10,
				CensorshipText:            "Сообщение изъято модератором",
				BotUserId:                 2, // Р. Букашка
				MaxAttachCountPerMessage:  10,
			}
			return nil
		},
	)

	if server.SetupError != nil {
		return
	}

	httpHandler, markAsUnavailable := routes.MakeHandler(
		appConfig,
		app.MakeServices(mysqlDB, redisClient, memcacheClient, cryptoCoder, minioClient, minioBucket),
	)

	httpServer := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: httpHandler,
	}

	server.Start = func() error {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	}
	server.Stop = func(ctx context.Context) error {
		markAsUnavailable()
		httpServer.SetKeepAlivesEnabled(false)
		return httpServer.Shutdown(ctx)
	}
	server.ShutdownTimeout = 5 * time.Second

	return
}

func makeMinioWithSecrets(server, accessKeyFile, secretKeyFile string) (*minio.Client, error) {
	accessKey, err := ioutil.ReadFile(accessKeyFile)
	if err != nil {
		return nil, err
	}
	secretKey, err := ioutil.ReadFile(secretKeyFile)
	if err != nil {
		return nil, err
	}
	return minio.New(server, string(accessKey), string(secretKey), false)
}

func makeMonitoringServer() (server *anyserver.Server) {
	port := os.Getenv("MONITORING_PORT")
	if port == "" {
		return nil
	}

	server = new(anyserver.Server)

	var httpHandler http.Handler

	{
		rootGroup := new(httprouter.Group)
		{
			rootGroup.Endpoint(http.MethodGet, "/pprof/:index", http.HandlerFunc(pprof.Index))
			rootGroup.Endpoint(http.MethodGet, "/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
			rootGroup.Endpoint(http.MethodGet, "/pprof/profile", http.HandlerFunc(pprof.Profile))
			rootGroup.Endpoint(http.MethodGet, "/pprof/symbol", http.HandlerFunc(pprof.Symbol))
			rootGroup.Endpoint(http.MethodGet, "/pprof/trace", http.HandlerFunc(pprof.Trace))
		}
		{
			rootGroup.Endpoint(http.MethodGet, "/expvar", expvar.Handler())
		}

		routerConfig := &httprouter.Config{
			RootGroup: rootGroup,
			NotFoundHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			}),
			CommonPrefix: "debug",
		}

		httpHandler, _ = httprouter.NewRouter(routerConfig)
	}

	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: httpHandler,
	}

	server.Start = func() error {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	}
	server.Stop = func(ctx context.Context) error {
		httpServer.SetKeepAlivesEnabled(false)

		return httpServer.Shutdown(ctx)
	}
	server.ShutdownTimeout = 5 * time.Second

	return
}
