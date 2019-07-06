package forumapi

import (
	"fmt"
	"net/http"
	"strconv"

	"fantlab/shared"
	"fantlab/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Controller struct {
	services *shared.Services
}

func NewController(services *shared.Services) *Controller {
	return &Controller{services: services}
}

func (c *Controller) ShowForums(ctx *gin.Context) {
	dbForums, err := fetchForums(c.services.DB, c.services.Config.DefaultAccessToForums)

	if err != nil {
		utils.ShowError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	dbModerators, err := fetchModerators(c.services.DB)

	if err != nil {
		utils.ShowError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	forumBlocks := getForumBlocks(dbForums, dbModerators, c.services.UrlFormatter)
	utils.ShowProto(ctx, http.StatusOK, forumBlocks)
}

func (c *Controller) ShowForumTopics(ctx *gin.Context) {
	forumID, err := strconv.ParseUint(ctx.Param("id"), 10, 16)

	if err != nil {
		utils.ShowError(ctx, http.StatusBadRequest, fmt.Sprintf("incorrect forum id: %s", ctx.Param("id")))
		return
	}

	page, err := strconv.ParseUint(ctx.DefaultQuery("page", "1"), 10, 32)

	if err != nil {
		utils.ShowError(ctx, http.StatusBadRequest, fmt.Sprintf("incorrect page: %s", ctx.Query("page")))
		return
	}

	defaultLimit := strconv.Itoa(int(c.services.Config.ForumTopicsInPage))
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", defaultLimit), 10, 32)

	if err != nil || !utils.IsValidLimit(limit) {
		utils.ShowError(ctx, http.StatusBadRequest, fmt.Sprintf("incorrect limit: %s", ctx.Query("limit")))
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
		if gorm.IsRecordNotFoundError(err) {
			utils.ShowError(ctx, http.StatusNotFound, fmt.Sprintf("incorrect forum id: %d", forumID))
		} else {
			utils.ShowError(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}

	forumTopics := getForumTopics(dbForumTopics, c.services.UrlFormatter)
	utils.ShowProto(ctx, http.StatusOK, forumTopics)
}

func (c *Controller) ShowTopicMessages(ctx *gin.Context) {
	topicID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		utils.ShowError(ctx, http.StatusBadRequest, fmt.Sprintf("incorrect topic id: %s", ctx.Param("id")))
		return
	}

	page, err := strconv.ParseUint(ctx.DefaultQuery("page", "1"), 10, 32)

	if err != nil {
		utils.ShowError(ctx, http.StatusBadRequest, fmt.Sprintf("incorrect page: %s", ctx.Query("page")))
		return
	}

	defaultLimit := strconv.Itoa(int(c.services.Config.ForumMessagesInPage))
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", defaultLimit), 10, 32)

	if err != nil || !utils.IsValidLimit(limit) {
		utils.ShowError(ctx, http.StatusBadRequest, fmt.Sprintf("incorrect limit: %s", ctx.Query("limit")))
		return
	}

	offset := limit * (page - 1)

	dbShortForumTopic, dbTopicMessages, err := fetchTopicMessages(
		c.services.DB,
		c.services.Config.DefaultAccessToForums,
		uint32(topicID),
		uint32(limit),
		uint32(offset))

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ShowError(ctx, http.StatusNotFound, fmt.Sprintf("incorrect topic id: %d", topicID))
		} else {
			utils.ShowError(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}

	topicMessages := getTopic(dbShortForumTopic, dbTopicMessages, c.services.UrlFormatter)
	utils.ShowProto(ctx, http.StatusOK, topicMessages)
}
