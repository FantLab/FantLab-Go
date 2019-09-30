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
	defer mysql.Close()

	mc := memcache.New(os.Getenv("MC_ADDRESS"))

	services := shared.MakeServices(
		db.NewDB(sqlr.Log(sqldb.New(mysql), logger.DB)),
		cache.New(caches.NewMemcache(mc)),
	)

	config := makeConfig(os.Getenv("IMAGES_BASE_URL"))

	handler := api.MakeRouter(config, services)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))
}

func makeConfig(imagesBaseURL string) *shared.AppConfig {
	return &shared.AppConfig{
		ImagesBaseURL:         imagesBaseURL,
		BlogsInPage:           50,
		BlogTopicsInPage:      5,
		ForumTopicsInPage:     20,
		ForumMessagesInPage:   20,
		DefaultAccessToForums: []uint16{1, 2, 3, 5, 6, 7, 8, 10, 12, 13, 14, 15, 16, 17, 20, 22},
	}
}
