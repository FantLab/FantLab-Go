package server

import (
	"database/sql"
	"fantlab/base/logs/logger"
	"fantlab/docs"
	"fantlab/server/internal/app"
	"fantlab/server/internal/config"
	"fantlab/server/internal/router"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GenerateDocs() {
	_ = docs.Generate(os.Stdout, router.Routes(nil, nil, nil), router.BasePath)
}

func Start() {
	mysql, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := mysql.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	isDebug := os.Getenv("DEBUG") == "1"

	router := router.MakeRouter(
		makeConfig(os.Getenv("IMAGES_BASE_URL")),
		app.MakeServices(isDebug, mysql, os.Getenv("MC_ADDRESS")),
		logFunc(isDebug),
		isDebug,
	)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}

func logFunc(isDebug bool) logger.ToString {
	if isDebug {
		return logger.Console
	}
	return logger.JSON
}

func makeConfig(imagesBaseURL string) *config.AppConfig {
	return &config.AppConfig{
		ImagesBaseURL:       imagesBaseURL,
		BlogsInPage:         50,
		BlogTopicsInPage:    5,
		ForumTopicsInPage:   20,
		ForumMessagesInPage: 20,
		// https://github.com/parserpro/fantlab/blob/ea456f3e8b8f9e02ab13ca2cdb9c335d36884d93/config/main.cfg#L402
		// 20 (один из первоапрельских форумов) убрал из списка
		DefaultAccessToForums: []uint64{1, 2, 3, 5, 6, 7, 8, 10, 12, 13, 14, 15, 16, 17, 22},
		// https://github.com/parserpro/fantlab/blob/ce769f66c5eacd59f487de840eb4bf62cac733a2/config/misc.cfg#L71
		CensorshipText: "Сообщение изъято модератором",
	}
}
