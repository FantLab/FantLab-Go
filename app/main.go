package main

import (
	"database/sql"
	"fantlab/api"
	"fantlab/cache"
	"fantlab/caches"
	"fantlab/db"
	"fantlab/dbtools/sqldb"
	"fantlab/dbtools/sqlr"
	"fantlab/logger"
	"fantlab/shared"
	"log"
	"net/http"
	"os"

	"github.com/bradfitz/gomemcache/memcache"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	startServer()
}

func startServer() {
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

	mc := memcache.New(os.Getenv("MC_ADDRESS"))

	services := shared.MakeServices(
		db.NewDB(sqlr.Log(sqldb.New(mysql), logger.DB)),
		cache.New(caches.Log("Memcached", caches.NewMemcache(mc), logger.Cache)),
	)

	config := makeConfig(os.Getenv("IMAGES_BASE_URL"))

	handler := api.MakeRouter(config, services)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))
}

func makeConfig(imagesBaseURL string) *shared.AppConfig {
	return &shared.AppConfig{
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
