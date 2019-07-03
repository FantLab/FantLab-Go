package main

import (
	"log"
	"os"

	"fantlab/config"
	"fantlab/logger"
	"fantlab/routing"
	"fantlab/shared"
	"fantlab/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	db.SetLogger(logger.GormLogger)
	db.LogMode(true)

	configuration := config.ParseConfig(os.Getenv("CONFIG_FILE"))

	services := &shared.Services{
		Config:       configuration,
		DB:           db,
		UrlFormatter: utils.UrlFormatter{Config: &configuration},
	}

	router := routing.SetupWith(services)

	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
