package server

import (
	"context"
	"database/sql"
	"fantlab/base/anyserver"
	"fantlab/base/codeflow"
	"fantlab/base/edsign"
	"fantlab/base/logs/logger"
	"fantlab/base/memcacheclient"
	"fantlab/base/redisclient"
	"fantlab/base/sharedconfig"
	"fantlab/docs"
	"fantlab/server/internal/app"
	"fantlab/server/internal/config"
	"fantlab/server/internal/routes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GenerateDocs() {
	_ = docs.Generate(os.Stdout, routes.Tree(nil, nil, nil), "/"+routes.BasePath)
}

func Start() {
	logsChan := make(chan string)

	go func() {
		mainLogger := log.New(os.Stdout, "$ ", 0)
		for {
			mainLogger.Println(<-logsChan)
		}
	}()

	apiServer := makeAPIServer(func(s string) {
		logsChan <- s
	})

	anyserver.RunWithGracefulShutdown(func(err error) {
		logsChan <- err.Error()
	}, apiServer)

	time.Sleep(1 * time.Second)
}

func makeAPIServer(logFunc func(string)) (server *anyserver.Server) {
	server = new(anyserver.Server)

	var mysqlDB *sql.DB
	var redisClient redisclient.Client
	var memcacheClient memcacheclient.Client
	var cryptoCoder *edsign.Coder
	var appConfig *config.AppConfig

	server.SetupError = codeflow.Try(
		func() error { // мускуль
			db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
			if err == nil {
				err = db.Ping()
			}
			if err != nil {
				return fmt.Errorf("MySQL setup error: %v", err)
			}
			mysqlDB = db
			server.DisposeBag = append(server.DisposeBag, db.Close)
			return nil
		},
		func() error { // редис (опционально)
			serverAddr := os.Getenv("RDS_ADDRESS")
			if len(serverAddr) == 0 {
				return nil
			}

			client, close := redisclient.NewPool(serverAddr, 8)
			err := client.Perform(context.Background(), func(conn redisco.Conn) error {
				_, err := conn.Do("PING")
				return err
			})
			if err != nil {
				return fmt.Errorf("Redis setup error: %v", err)
			}
			redisClient = client
			server.DisposeBag = append(server.DisposeBag, close)
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
			memcacheClient = client
			return nil
		},
		func() error { // криптокодер для jwt-like токенов
			coder, err := edsign.NewCoder64(os.Getenv("SIGN_PUB_KEY"), os.Getenv("SIGN_PRIV_KEY"))
			if err != nil {
				return fmt.Errorf("JWT setup error: %v", err)
			}
			cryptoCoder = coder
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
				BlogTopicsInPage:          5,
				BlogArticleCommentsInPage: 10,
				CensorshipText:            "Сообщение изъято модератором",
				BotUserId:                 2, // Р. Букашка
			}
			return nil
		},
	)

	if server.SetupError != nil {
		return
	}

	var requestToString func(r *logger.Request) string
	if sharedconfig.IsDebug() {
		requestToString = logger.Console
	} else {
		requestToString = logger.JSON
	}

	httpServer := &http.Server{
		Addr: ":" + os.Getenv("PORT"),
		Handler: routes.MakeHandler(
			appConfig,
			app.MakeServices(mysqlDB, redisClient, memcacheClient, cryptoCoder),
			func(r *logger.Request) {
				logFunc(requestToString(r))
			},
		),
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
