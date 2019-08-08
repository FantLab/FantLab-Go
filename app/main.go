package main

import (
	"database/sql"
	"fantlab/logger"
	"log"
	"os"

	"fantlab/cache"
	"fantlab/db"
	"fantlab/routing"
	"fantlab/shared"
	"fantlab/sqlr"

	"github.com/bradfitz/gomemcache/memcache"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mysql, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer mysql.Close()

	mc := memcache.New(os.Getenv("MC_ADDRESS"))

	services := &shared.Services{
		Config: makeConfig(os.Getenv("IMAGES_BASE_URL")),
		Cache:  &cache.MemCache{Client: mc},
		DB:     &db.DB{R: sqlr.New(mysql, logger.Sqlr)},
	}

	router := routing.SetupWith(services)

	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
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
