package server

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fantlab/base/anyserver"
	"fantlab/base/codeflow"
	"fantlab/base/edsign"
	"fantlab/base/logs/logger"
	"fantlab/base/memcached"
	"fantlab/base/redisco"
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
	var redisClient redisco.Client
	var memcacheClient memcached.Client
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
		func() error { // редис
			client, close := redisco.NewPool(os.Getenv("RDS_ADDRESS"), 8)
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
		func() error { // мемкэш
			client := memcached.New(os.Getenv("MC_ADDRESS"))
			err := client.Ping()
			if err != nil {
				return fmt.Errorf("Memcache setup error: %v", err)
			}
			memcacheClient = client
			return nil
		},
		func() error { // криптокодер для jwt-like токенов
			pubKey, err := base64.StdEncoding.DecodeString(os.Getenv("SIGN_PUB_KEY"))
			if err != nil {
				return fmt.Errorf("Invalid JWT signer public key: %v", err)
			}
			privKey, err := base64.StdEncoding.DecodeString(os.Getenv("SIGN_PRIV_KEY"))
			if err != nil {
				return fmt.Errorf("Invalid JWT signer private key: %v", err)
			}
			cryptoCoder = edsign.NewCoder(pubKey, privKey)
			return nil
		},
		func() error { // конфигурация бизнес-логики
			appConfig = &config.AppConfig{
				ImagesBaseURL:         os.Getenv("IMAGES_BASE_URL"),
				ForumTopicsInPage:     20,
				ForumMessagesInPage:   20,
				MaxForumMessageLength: 20000,
				// https://github.com/parserpro/fantlab/blob/ea456f3e8b8f9e02ab13ca2cdb9c335d36884d93/config/main.cfg#L402
				// 20 (один из первоапрельских форумов) убрал из списка
				DefaultAccessToForums: []uint64{1, 2, 3, 5, 6, 7, 8, 10, 12, 13, 14, 15, 16, 17, 22},
				// https://github.com/parserpro/fantlab/blob/ce769f66c5eacd59f487de840eb4bf62cac733a2/config/misc.cfg#L71
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

	isDebug := os.Getenv("DEBUG") != ""

	var requestToString func(r *logger.Request) string
	if isDebug {
		requestToString = logger.Console
	} else {
		requestToString = logger.JSON
	}

	httpServer := &http.Server{
		Addr: ":" + os.Getenv("PORT"),
		Handler: routes.MakeHandler(
			appConfig,
			app.MakeServices(isDebug, mysqlDB, redisClient, memcacheClient, cryptoCoder),
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
