package main

import (
	"log"
	"os"

	"fantlab/config"
	"fantlab/db"
	"fantlab/logger"
	"fantlab/routing"
	"fantlab/shared"
	"fantlab/utils"

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

	configuration := config.ParseConfig(os.Getenv("CONFIG_FILE"))

	services := &shared.Services{
		Config:       configuration,
		DB:           &db.DB{ORM: orm},
		UrlFormatter: utils.UrlFormatter{Config: &configuration},
	}

	router := routing.SetupWith(services)

	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
