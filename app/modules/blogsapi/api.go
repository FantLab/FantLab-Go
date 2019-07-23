package blogsapi

import (
	"fantlab/pb"
	"net/http"
	"strconv"
	"strings"

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

func (c *Controller) ShowCommunities(ctx *gin.Context) {
	dbCommunities, err := c.services.DB.FetchCommunities()

	if err != nil {
		utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		})
		return
	}

	communities := getCommunities(dbCommunities, c.services.Config)
	utils.ShowProto(ctx, http.StatusOK, communities)
}

func (c *Controller) ShowCommunity(ctx *gin.Context) {
	communityId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

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

	defaultLimit := strconv.Itoa(int(c.services.Config.BlogTopicsInPage))
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", defaultLimit), 10, 32)

	if err != nil || !utils.IsValidLimit(limit) {
		utils.ShowProto(ctx, http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "limit",
		})
		return
	}

	offset := limit * (page - 1)

	dbResponse, err := c.services.DB.FetchCommunity(uint32(communityId), uint32(limit), uint32(offset))

	if err != nil {
		if utils.IsRecordNotFoundError(err) {
			utils.ShowProto(ctx, http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(communityId, 10),
			})
		} else {
			utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			})
		}
		return
	}

	community := getCommunity(dbResponse, uint32(page), uint32(limit), c.services.Config)
	utils.ShowProto(ctx, http.StatusOK, community)
}

func (c *Controller) ShowBlogs(ctx *gin.Context) {
	page, err := strconv.ParseUint(ctx.DefaultQuery("page", "1"), 10, 32)

	if err != nil {
		utils.ShowProto(ctx, http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "page",
		})
		return
	}

	defaultLimit := strconv.Itoa(int(c.services.Config.BlogsInPage))
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", defaultLimit), 10, 32)

	if err != nil || !utils.IsValidLimit(limit) {
		utils.ShowProto(ctx, http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "limit",
		})
		return
	}

	sort := strings.ToLower(ctx.DefaultQuery("sort", "update"))
	offset := limit * (page - 1)

	dbResponse, err := c.services.DB.FetchBlogs(uint32(limit), uint32(offset), sort)

	if err != nil {
		utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		})
		return
	}

	blogs := getBlogs(dbResponse, uint32(page), uint32(limit), c.services.Config)
	utils.ShowProto(ctx, http.StatusOK, blogs)
}

func (c *Controller) ShowBlog(ctx *gin.Context) {
	blogID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

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

	defaultLimit := strconv.Itoa(int(c.services.Config.BlogTopicsInPage))
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", defaultLimit), 10, 32)

	if err != nil || !utils.IsValidLimit(limit) {
		utils.ShowProto(ctx, http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "limit",
		})
		return
	}

	offset := limit * (page - 1)

	dbResponse, err := c.services.DB.FetchBlog(uint32(blogID), uint32(limit), uint32(offset))

	if err != nil {
		if utils.IsRecordNotFoundError(err) {
			utils.ShowProto(ctx, http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(blogID, 10),
			})
		} else {
			utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			})
		}
		return
	}

	blog := getBlog(dbResponse, uint32(page), uint32(limit), c.services.Config)
	utils.ShowProto(ctx, http.StatusOK, blog)
}

func (c *Controller) ShowArticle(ctx *gin.Context) {
	articleId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		utils.ShowProto(ctx, http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "id",
		})
		return
	}

	dbTopic, err := c.services.DB.FetchBlogTopic(uint32(articleId))

	if err != nil {
		if utils.IsRecordNotFoundError(err) {
			utils.ShowProto(ctx, http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(articleId, 10),
			})
		} else {
			utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			})
		}
		return
	}

	article := getArticle(dbTopic, c.services.Config)
	utils.ShowProto(ctx, http.StatusOK, article)
}

func (c *Controller) LikeArticle(ctx *gin.Context) {
	userId := ctx.GetInt64(gin.AuthUserKey)

	if userId == 0 {
		utils.ShowProto(ctx, http.StatusUnauthorized, &pb.Error_Response{
			Status: pb.Error_INVALID_SESSION,
		})
		return
	}

	articleId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		utils.ShowProto(ctx, http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "id",
		})
		return
	}

	dbTopicCreatorId, err := c.services.DB.FetchBlogTopicCreatorId(uint32(articleId))

	if err != nil {
		if utils.IsRecordNotFoundError(err) {
			utils.ShowProto(ctx, http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(articleId, 10),
			})
		} else {
			utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			})
		}
		return
	}

	if *dbTopicCreatorId == uint32(userId) {
		utils.ShowProto(ctx, http.StatusUnauthorized, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "your own article",
		})
		return
	}

	isDbTopicLiked, err := c.services.DB.FetchBlogTopicLiked(uint32(articleId), uint32(userId))

	if err != nil {
		utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		})
		return
	}

	if *isDbTopicLiked == true {
		utils.ShowProto(ctx, http.StatusUnauthorized, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "already liked",
		})
		return
	}

	err = c.services.DB.UpdateBlogTopicLiked(uint32(articleId), uint32(userId))

	if err != nil {
		utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		})
		return
	}

	dbTopicLikeCount, err := c.services.DB.FetchBlogTopicLikeCount(uint32(articleId))

	if err != nil {
		utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		})
		return
	}

	utils.ShowProto(ctx, http.StatusOK, &pb.Blog_BlogArticleLikeResponse{
		LikeCount: *dbTopicLikeCount,
	})
}

func (c *Controller) DislikeArticle(ctx *gin.Context) {
	userId := ctx.GetInt64(gin.AuthUserKey)

	if userId == 0 {
		utils.ShowProto(ctx, http.StatusUnauthorized, &pb.Error_Response{
			Status: pb.Error_INVALID_SESSION,
		})
		return
	}

	articleId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		utils.ShowProto(ctx, http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "id",
		})
		return
	}

	dbTopicCreatorId, err := c.services.DB.FetchBlogTopicCreatorId(uint32(articleId))

	if err != nil {
		if utils.IsRecordNotFoundError(err) {
			utils.ShowProto(ctx, http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(articleId, 10),
			})
		} else {
			utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			})
		}
		return
	}

	if *dbTopicCreatorId == uint32(userId) {
		utils.ShowProto(ctx, http.StatusUnauthorized, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "your own article",
		})
		return
	}

	isDbTopicLiked, err := c.services.DB.FetchBlogTopicLiked(uint32(articleId), uint32(userId))

	if err != nil {
		utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		})
		return
	}

	if *isDbTopicLiked == false {
		utils.ShowProto(ctx, http.StatusUnauthorized, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "already disliked",
		})
		return
	}

	err = c.services.DB.UpdateBlogTopicDisliked(uint32(articleId), uint32(userId))

	if err != nil {
		utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		})
		return
	}

	dbTopicLikeCount, err := c.services.DB.FetchBlogTopicLikeCount(uint32(articleId))

	if err != nil {
		utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		})
		return
	}

	utils.ShowProto(ctx, http.StatusOK, &pb.Blog_BlogArticleLikeResponse{
		LikeCount: *dbTopicLikeCount,
	})
}
