package forumapi

import (
	"net/http"
	"strconv"

	"fantlab/pb"
	"fantlab/shared"
	"fantlab/utils"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	services *shared.Services
}

func NewController(services *shared.Services) *Controller {
	return &Controller{services: services}
}

func (c *Controller) ShowForums(ctx *gin.Context) {
	dbForums, err := c.services.DB.FetchForums(c.services.Config.DefaultAccessToForums)

	if err != nil {
		utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		})
		return
	}

	dbModerators, err := c.services.DB.FetchModerators()

	if err != nil {
		utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		})
		return
	}

	forumBlocks := getForumBlocks(dbForums, dbModerators, c.services.Config)
	utils.ShowProto(ctx, http.StatusOK, forumBlocks)
}

func (c *Controller) ShowForumTopics(ctx *gin.Context) {
	forumID, err := strconv.ParseUint(ctx.Param("id"), 10, 16)

	if err != nil {
		utils.ShowProto(ctx, http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "id",
		})
		return
	}

	page, err := strconv.ParseUint(ctx.DefaultQuery("page", "1"), 10, 32)

	if err != nil {
		utils.ShowProto(ctx, http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "page",
		})
		return
	}

	defaultLimit := strconv.Itoa(int(c.services.Config.ForumTopicsInPage))
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", defaultLimit), 10, 32)

	if err != nil || !utils.IsValidLimit(limit) {
		utils.ShowProto(ctx, http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "limit",
		})
		return
	}

	offset := limit * (page - 1)

	dbResponse, err := c.services.DB.FetchForumTopics(
		c.services.Config.DefaultAccessToForums,
		uint16(forumID),
		uint32(limit),
		uint32(offset),
	)

	if err != nil {
		if utils.IsRecordNotFoundError(err) {
			utils.ShowProto(ctx, http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(forumID, 10),
			})
		} else {
			utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			})
		}
		return
	}

	forumTopics := getForumTopics(dbResponse, uint32(page), uint32(limit), c.services.Config)
	utils.ShowProto(ctx, http.StatusOK, forumTopics)
}

func (c *Controller) ShowTopicMessages(ctx *gin.Context) {
	topicID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		utils.ShowProto(ctx, http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "id",
		})
		return
	}

	page, err := strconv.ParseUint(ctx.DefaultQuery("page", "1"), 10, 32)

	if err != nil {
		utils.ShowProto(ctx, http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "page",
		})
		return
	}

	defaultLimit := strconv.Itoa(int(c.services.Config.ForumMessagesInPage))
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", defaultLimit), 10, 32)

	if err != nil || !utils.IsValidLimit(limit) {
		utils.ShowProto(ctx, http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "limit",
		})
		return
	}

	offset := limit * (page - 1)

	dbResponse, err := c.services.DB.FetchTopicMessages(
		c.services.Config.DefaultAccessToForums,
		uint32(topicID),
		uint32(limit),
		uint32(offset),
	)

	if err != nil {
		if utils.IsRecordNotFoundError(err) {
			utils.ShowProto(ctx, http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(topicID, 10),
			})
		} else {
			utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			})
		}
		return
	}

	topicMessages := getTopic(dbResponse, uint32(page), uint32(limit), c.services.Config)
	utils.ShowProto(ctx, http.StatusOK, topicMessages)
}
