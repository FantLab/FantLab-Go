package server

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fantlab/base/edsign"
	"fantlab/base/logs/logger"
	"fantlab/base/redisco"
	"fantlab/docs"
	"fantlab/server/internal/app"
	"fantlab/server/internal/config"
	"fantlab/server/internal/routes"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GenerateDocs() {
	_ = docs.Generate(os.Stdout, routes.Tree(nil, nil, nil), "/"+routes.BasePath)
}

func Start() {
	isDebug := os.Getenv("DEBUG") == "1"

	mysqlDB, closeDB := makeDB(os.Getenv("MYSQL_URL"))
	defer func() {
		if err := closeDB(); err != nil {
			log.Println(err)
		}
	}()

	redisClient, closeRedis := makeRedis(os.Getenv("RDS_ADDRESS"))
	defer func() {
		if err := closeRedis(); err != nil {
			log.Println(err)
		}
	}()

	handler := routes.MakeHandler(
		makeConfig(os.Getenv("IMAGES_BASE_URL")),
		app.MakeServices(
			isDebug,
			makeCryptoCoder(),
			mysqlDB,
			redisClient,
			os.Getenv("MC_ADDRESS"),
		),
		logFunc(isDebug),
		isDebug,
	)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))
}

func logFunc(isDebug bool) logger.ToString {
	if isDebug {
		return logger.Console
	}
	return logger.JSON
}

func makeDB(cs string) (*sql.DB, func() error) {
	db, err := sql.Open("mysql", cs)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db, db.Close
}

func makeRedis(server string) (redisco.Client, func() error) {
	client, close := redisco.NewPool(server, 8)
	err := client.Perform(context.Background(), func(conn redisco.Conn) error {
		_, err := conn.Do("PING")
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	return client, close
}

func makeCryptoCoder() *edsign.Coder {
	pubKey, err := base64.StdEncoding.DecodeString(os.Getenv("SIGN_PUB_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	privKey, err := base64.StdEncoding.DecodeString(os.Getenv("SIGN_PRIV_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	return edsign.NewCoder(pubKey, privKey)
}

func makeConfig(imagesBaseURL string) *config.AppConfig {
	return &config.AppConfig{
		ImagesBaseURL:         imagesBaseURL,
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
}
