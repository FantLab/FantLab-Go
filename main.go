package main

import (
	"log"
	"os"

	"fantlab/config"
	"fantlab/logger"
	"fantlab/routing"
	"fantlab/shared"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", os.Getenv("MYSQL_CS"))
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

	isDebug := os.Getenv("DEBUG") == "1"
	if !isDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	configuration := config.ParseConfig()
	configuration.IsDebug = isDebug
	services := &shared.Services{
		Config: configuration,
		DB:     db,
	}

	router := routing.SetupWith(services)

	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
