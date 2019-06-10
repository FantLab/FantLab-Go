package main

import (
	"fantlab/config"
	"fantlab/routing"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// db, err := gorm.Open("mysql", "root:root@/fantlab?charset=utf8&parseTime=True&loc=Europe%2FMoscow")
	db, err := gorm.Open("mysql", "root@/fl?charset=utf8&parseTime=True&loc=Local")
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
	fldb := &config.FLDB{db}

	router := routing.SetupWith(fldb)

	if err := router.Run(":4242"); err != nil {
		log.Fatal(err)
	}
}
