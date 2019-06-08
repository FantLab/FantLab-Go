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
	dbTopics := c.db.getTopics(uint16(forumId), uint32(limit), uint32(offset))
	topics := getTopics(dbTopics)
	ctx.JSON(http.StatusOK, topics)
}
