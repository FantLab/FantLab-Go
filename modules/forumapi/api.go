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
	utils.ShowJson(ctx, http.StatusOK, forumBlocks)
}

func (c *Controller) ShowForumTopics(ctx *gin.Context) {
	forumID, err := strconv.ParseUint(ctx.Param("id"), 10, 16)

	if err != nil {
		utils.ShowError(ctx, http.StatusBadRequest, "incorrect forum id")
		return
	}

	page, err := strconv.ParseUint(ctx.Query("page"), 10, 32)

	if err != nil {
		utils.ShowError(ctx, http.StatusBadRequest, "incorrect page")
		return
	}

	defaultLimit := strconv.Itoa(int(c.services.Config.ForumTopicsInPage))
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", defaultLimit), 10, 32)

	if err != nil {
		utils.ShowError(ctx, http.StatusBadRequest, "incorrect limit")
		return
	}

	offset := limit * (page - 1)

	dbForumTopics, err := fetchForumTopics(
		c.services.DB,
		c.services.Config.DefaultAccessToForums,
		uint16(forumID),
		uint32(limit),
		uint32(offset))

	if err != nil {
		utils.ShowError(ctx, http.StatusNotFound, err.Error())
		return
	}

	forumTopics := getForumTopics(dbForumTopics)
	utils.ShowJson(ctx, http.StatusOK, forumTopics)
}

func (c *Controller) ShowTopicMessages(ctx *gin.Context) {
	topicID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		utils.ShowError(ctx, http.StatusBadRequest, "incorrect topic id")
		return
	}

	page, err := strconv.ParseUint(ctx.Query("page"), 10, 32)

	if err != nil {
		utils.ShowError(ctx, http.StatusBadRequest, "incorrect page")
		return
	}

	defaultLimit := strconv.Itoa(int(c.services.Config.ForumMessagesInPage))
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", defaultLimit), 10, 32)

	if err != nil {
		utils.ShowError(ctx, http.StatusBadRequest, "incorrect limit")
		return
	}

	offset := limit * (page - 1)

	dbTopicMessages, err := fetchTopicMessages(
		c.services.DB,
		c.services.Config.DefaultAccessToForums,
		uint32(topicID),
		uint32(limit),
		uint32(offset))

	if err != nil {
		utils.ShowError(ctx, http.StatusNotFound, err.Error())
		return
	}

	topicMessages := getTopicMessages(dbTopicMessages, c.services.Config.ImageUrl)
	utils.ShowJson(ctx, http.StatusOK, topicMessages)
}
