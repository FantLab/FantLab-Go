package main

import (
	"log"
	"net/http"
	"strconv"

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
	router.GET("/forums/:id", forumTopicsEndpoint)

	router.Run()
}

func forumsEndpoint(c *gin.Context) {
	dbForums := fdb.getForums()
	dbModerators := fdb.getModerators()
	forumBlocks := getForumBlocks(dbForums, dbModerators)
	c.JSON(http.StatusOK, forumBlocks)
}

func forumTopicsEndpoint(c *gin.Context) {
	forumId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20")) // todo get from config
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	offset := limit * (page - 1)
	dbTopics := fdb.getTopics(uint16(forumId), uint16(limit), uint16(offset))
	topics := getTopics(dbTopics)
	c.JSON(http.StatusOK, topics)
}
