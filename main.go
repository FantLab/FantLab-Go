package main

import (
	"fantlab/routing"
	"fantlab/shared"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@/fantlab?charset=utf8&parseTime=True&loc=Europe%2FMoscow")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	db.LogMode(true)

	services := &shared.Services{DB: db}

	router := routing.SetupWith(services)

	if err := router.Run(":4242"); err != nil {
		log.Fatal(err)
	}
}
