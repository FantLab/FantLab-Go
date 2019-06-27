package forumapi

import (
	"fantlab/shared"
	"fantlab/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	services *shared.Services
}

func NewController(services *shared.Services) *Controller {
	return &Controller{services: services}
}

func (c *Controller) ShowForums(ctx *gin.Context) {
	dbForums := fetchForums(c.services.DB, c.services.Config.DefaultAccessToForums)
	dbModerators := fetchModerators(c.services.DB)
	forumBlocks := getForumBlocks(dbForums, dbModerators)
	ctx.JSON(http.StatusOK, forumBlocks)
}

func (c *Controller) ShowForumTopics(ctx *gin.Context) {
	forumID, err := strconv.ParseUint(ctx.Param("id"), 10, 16)
	if err != nil {
		//noinspection GoUnhandledErrorResult
		ctx.Error(err)
	}
	page, err := strconv.ParseUint(ctx.Query("page"), 10, 32)
	if err != nil {
		//noinspection GoUnhandledErrorResult
		ctx.Error(err)
	}
	defaultLimit := strconv.Itoa(int(c.services.Config.ForumTopicsInPage))
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", defaultLimit), 10, 32)
	if err != nil {
		//noinspection GoUnhandledErrorResult
		ctx.Error(err)
	}
	if len(ctx.Errors) != 0 {
		utils.ShowErrors(ctx)
		return
	}
	offset := limit * (page - 1)
	dbForumTopics := fetchForumTopics(
		c.services.DB,
		c.services.Config.DefaultAccessToForums,
		uint16(forumID),
		uint32(limit),
		uint32(offset))
	forumTopics := getForumTopics(dbForumTopics)
	ctx.JSON(http.StatusOK, forumTopics)
}

func (c *Controller) ShowTopicMessages(ctx *gin.Context) {
	topicID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		//noinspection GoUnhandledErrorResult
		ctx.Error(err)
	}
	page, err := strconv.ParseUint(ctx.Query("page"), 10, 32)
	if err != nil {
		//noinspection GoUnhandledErrorResult
		ctx.Error(err)
	}
	defaultLimit := strconv.Itoa(int(c.services.Config.ForumMessagesInPage))
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", defaultLimit), 10, 32)
	if err != nil {
		//noinspection GoUnhandledErrorResult
		ctx.Error(err)
	}
	if len(ctx.Errors) != 0 {
		utils.ShowErrors(ctx)
		return
	}
	offset := limit * (page - 1)
	dbTopicMessages := fetchTopicMessages(
		c.services.DB,
		c.services.Config.DefaultAccessToForums,
		uint32(topicID),
		uint32(limit),
		uint32(offset))
	topicMessages := getTopicMessages(dbTopicMessages, c.services.Config)
	ctx.JSON(http.StatusOK, topicMessages)
}
