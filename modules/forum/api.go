package forumapi

import (
	"net/http"
	"strconv"

	"fantlab/shared"
	"fantlab/utils"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	services *shared.Services
}

func NewController(services *shared.Services) *Controller {
	return &Controller{
		services: services,
	}
}

func (c *Controller) ShowForums(ctx *gin.Context) {
	dbForums := fetchForums(c.services.DB)
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
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", "20"), 10, 32) // todo get from config
	if err != nil {
		//noinspection GoUnhandledErrorResult
		ctx.Error(err)
	}
	if len(ctx.Errors) != 0 {
		utils.ShowErrors(ctx)
		return
	}
	offset := limit * (page - 1)
	dbForumTopics := fetchForumTopics(c.services.DB, uint16(forumID), uint32(limit), uint32(offset))
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
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", "20"), 10, 32) // todo get from config
	if err != nil {
		//noinspection GoUnhandledErrorResult
		ctx.Error(err)
	}
	if len(ctx.Errors) != 0 {
		utils.ShowErrors(ctx)
		return
	}
	offset := limit * (page - 1)
	dbTopicMessages := fetchTopicMessages(c.services.DB, uint32(topicID), uint32(limit), uint32(offset))
	topicMessages := getTopicMessages(dbTopicMessages)
	ctx.JSON(http.StatusOK, topicMessages)
}
