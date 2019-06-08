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
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	db.LogMode(true)
	fdb = &FDB{db}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/forums", forumsEndpoint)
	router.GET("/forums/:id", forumTopicsEndpoint)

	if err := router.Run(":4242"); err != nil {
		log.Fatal(err)
	}
}

func forumsEndpoint(c *gin.Context) {
	dbForums := fdb.getForums()
	dbModerators := fdb.getModerators()
	forumBlocks := getForumBlocks(dbForums, dbModerators)
	c.JSON(http.StatusOK, forumBlocks)
}

func forumTopicsEndpoint(c *gin.Context) {
	forumId, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		//noinspection GoUnhandledErrorResult
		c.Error(err)
	}
	page, err := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 32)
	if err != nil {
		//noinspection GoUnhandledErrorResult
		c.Error(err)
	}
	limit, err := strconv.ParseUint(c.DefaultQuery("limit", "20"), 10, 32) // todo get from config
	if err != nil {
		//noinspection GoUnhandledErrorResult
		c.Error(err)
	}
	if len(c.Errors) != 0 {
		if gin.IsDebugging() {
			c.AbortWithStatusJSON(http.StatusBadRequest, c.Errors.JSON())
		} else {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		return
	}
	offset := limit * (page - 1)
	dbTopics := fdb.getTopics(uint16(forumId), uint32(limit), uint32(offset))
	topics := getTopics(dbTopics)
	c.JSON(http.StatusOK, topics)
}
