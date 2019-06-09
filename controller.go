package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	db *FDB
}

func NewController(db *FDB) *Controller {
	return &Controller{
		db: db,
	}
}

func (c *Controller) ShowForums(ctx *gin.Context) {
	dbForums := c.db.getForums()
	dbModerators := c.db.getModerators()
	forumBlocks := getForumBlocks(dbForums, dbModerators)
	ctx.JSON(http.StatusOK, forumBlocks)
}

func (c *Controller) showForumTopics(ctx *gin.Context) {
	forumId, err := strconv.ParseUint(ctx.Param("id"), 10, 16)
	if err != nil {
		//noinspection GoUnhandledErrorResult
		ctx.Error(err)
	}
	page, err := strconv.ParseUint(ctx.Query("page"), 10, 32)
	if err != nil {
		//noinspection GoUnhandledErrorResult
		ctx.Error(err)
	}
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", "20"), 10, 32) // todo get from config
	if err != nil {
		//noinspection GoUnhandledErrorResult
		ctx.Error(err)
	}
	if len(ctx.Errors) != 0 {
		if gin.IsDebugging() {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Errors.JSON())
		} else {
			ctx.AbortWithStatus(http.StatusBadRequest)
		}
		return
	}
	offset := limit * (page - 1)
	dbForumTopics := c.db.getForumTopics(uint16(forumId), uint32(limit), uint32(offset))
	forumTopics := getForumTopics(dbForumTopics)
	ctx.JSON(http.StatusOK, forumTopics)
}

func (c *Controller) showTopicMessages(ctx *gin.Context) {
	topicId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		//noinspection GoUnhandledErrorResult
		ctx.Error(err)
	}
	page, err := strconv.ParseUint(ctx.Query("page"), 10, 32)
	if err != nil {
		//noinspection GoUnhandledErrorResult
		ctx.Error(err)
	}
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", "20"), 10, 32) // todo get from config
	if err != nil {
		//noinspection GoUnhandledErrorResult
		ctx.Error(err)
	}
	if len(ctx.Errors) != 0 {
		if gin.IsDebugging() {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Errors.JSON())
		} else {
			ctx.AbortWithStatus(http.StatusBadRequest)
		}
		return
	}
	offset := limit * (page - 1)
	dbTopicMessages := c.db.getTopicMessages(uint32(topicId), uint32(limit), uint32(offset))
	topicMessages := getTopicMessages(dbTopicMessages)
	ctx.JSON(http.StatusOK, topicMessages)
}
