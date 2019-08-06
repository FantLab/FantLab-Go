package main

import (
	"database/sql"
	"log"
	"os"

	"fantlab/cache"
	"fantlab/db"
	"fantlab/logger"
	"fantlab/routing"
	"fantlab/shared"
	"fantlab/sqlr"

	"github.com/bradfitz/gomemcache/memcache"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	orm, err := gorm.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := orm.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	orm.SetLogger(logger.GormLogger)
	orm.LogMode(true)

	mysql, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer mysql.Close()

	mc := memcache.New(os.Getenv("MC_ADDRESS"))

	services := &shared.Services{
		Config: makeConfig(os.Getenv("IMAGES_BASE_URL")),
		Cache:  &cache.MemCache{Client: mc},
		DB: &db.DB{
			ORM: orm,
			R:   sqlr.New(mysql),
		},
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
