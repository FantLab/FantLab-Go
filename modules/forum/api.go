package forumapi

import (
	"fantlab/config"
	"fantlab/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	db *config.FLDB
}

func NewController(db *config.FLDB) *Controller {
	return &Controller{
		db: db,
	}
}

func (c *Controller) ShowForums(ctx *gin.Context) {
	dbForums := fetchForums(c.db)
	dbModerators := fetchModerators(c.db)
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
	dbForumTopics := fetchForumTopics(c.db, uint16(forumID), uint32(limit), uint32(offset))
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
	dbTopicMessages := fetchTopicMessages(c.db, uint32(topicID), uint32(limit), uint32(offset))
	topicMessages := getTopicMessages(dbTopicMessages)
	ctx.JSON(http.StatusOK, topicMessages)
}
