package blogsapi

import (
	"fmt"
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
	return &Controller{services: services}
}

func (c *Controller) ShowCommunities(ctx *gin.Context) {
	dbCommunities := fetchCommunities(c.services.DB)
	communities := getCommunities(dbCommunities, c.services.UrlFormatter)
	utils.ShowProto(ctx, http.StatusOK, communities)
}

func (c *Controller) ShowCommunity(ctx *gin.Context) {
	communityId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		utils.ShowError(ctx, http.StatusBadRequest, fmt.Sprintf("incorrect community id: %s", ctx.Param("id")))
		return
	}

	page, err := strconv.ParseUint(ctx.DefaultQuery("page", "1"), 10, 32)

	if err != nil {
		utils.ShowError(ctx, http.StatusBadRequest, fmt.Sprintf("incorrect page: %s", ctx.Query("page")))
		return
	}

	defaultLimit := strconv.Itoa(int(c.services.Config.BlogTopicsInPage))
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", defaultLimit), 10, 32)

	if err != nil || !utils.IsValidLimit(limit) {
		utils.ShowError(ctx, http.StatusBadRequest, fmt.Sprintf("incorrect limit: %s", ctx.Query("limit")))
		return
	}

	offset := limit * (page - 1)

	dbCommunity, dbModerators, dbAuthors, dbTopics, err :=
		fetchCommunity(c.services.DB, uint32(communityId), uint32(limit), uint32(offset))

	if err != nil {
		utils.ShowError(ctx, http.StatusNotFound, err.Error())
		return
	}

	community := getCommunity(dbCommunity, dbModerators, dbAuthors, dbTopics, c.services.UrlFormatter)
	utils.ShowProto(ctx, http.StatusOK, community)
}

func (c *Controller) ShowBlogs(ctx *gin.Context) {
	page, err := strconv.ParseUint(ctx.DefaultQuery("page", "1"), 10, 32)

	if err != nil {
		utils.ShowError(ctx, http.StatusBadRequest, fmt.Sprintf("incorrect page: %s", ctx.Query("page")))
		return
	}

	defaultLimit := strconv.Itoa(int(c.services.Config.BlogsInPage))
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", defaultLimit), 10, 32)

	if err != nil || !utils.IsValidLimit(limit) {
		utils.ShowError(ctx, http.StatusBadRequest, fmt.Sprintf("incorrect limit: %s", ctx.Query("limit")))
		return
	}

	sort := ctx.DefaultQuery("sort", "update")
	offset := limit * (page - 1)

	dbBlogs := fetchBlogs(c.services.DB, uint32(limit), uint32(offset), sort)
	blogs := getBlogs(dbBlogs, c.services.UrlFormatter)
	utils.ShowProto(ctx, http.StatusOK, blogs)
}

func (c *Controller) ShowBlog(ctx *gin.Context) {
	blogID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		utils.ShowError(ctx, http.StatusBadRequest, fmt.Sprintf("incorrect blog id: %s", ctx.Param("id")))
		return
	}

	page, err := strconv.ParseUint(ctx.DefaultQuery("page", "1"), 10, 32)

	if err != nil {
		utils.ShowError(ctx, http.StatusBadRequest, fmt.Sprintf("incorrect page: %s", ctx.Query("page")))
		return
	}

	defaultLimit := strconv.Itoa(int(c.services.Config.BlogTopicsInPage))
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", defaultLimit), 10, 32)

	if err != nil || !utils.IsValidLimit(limit) {
		utils.ShowError(ctx, http.StatusBadRequest, fmt.Sprintf("incorrect limit: %s", ctx.Query("limit")))
		return
	}

	offset := limit * (page - 1)

	dbBlogTopics, err := fetchBlog(c.services.DB, uint32(blogID), uint32(limit), uint32(offset))

	if err != nil {
		utils.ShowError(ctx, http.StatusNotFound, err.Error())
		return
	}

	blog := getBlog(dbBlogTopics, c.services.UrlFormatter)
	utils.ShowProto(ctx, http.StatusOK, blog)
}
