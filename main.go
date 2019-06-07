package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var fdb *FDB

func main() {
	db, err := gorm.Open("mysql", "root:root@/fantlab?charset=utf8&parseTime=True&loc=Europe%2FMoscow")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.LogMode(true)
	fdb = &FDB{db}

	router := gin.Default()

	router.GET("/forums", forumsEndpoint)

	router.Run()
}

func forumsEndpoint(c *gin.Context) {
	forums := fdb.getForums()
	moderators := fdb.getModerators()
	forumBlocks := getForumBlocks(forums, moderators)
	c.IndentedJSON(http.StatusOK, forumBlocks)
}
