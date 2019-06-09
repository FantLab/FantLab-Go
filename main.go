package main

import (
	"log"

	"github.com/gin-gonic/gin"
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
	fdb := &FDB{db}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	controller := NewController(fdb)

	router.GET("/forums", controller.ShowForums)
	router.GET("/forums/:id", controller.ShowForumTopics)
	router.GET("/topics/:id", controller.ShowTopicMessages)

	if err := router.Run(":4242"); err != nil {
		log.Fatal(err)
	}
}
