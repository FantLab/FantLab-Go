package blogsapi

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

func (c *Controller) ShowCommunities(ctx *gin.Context) {
	dbCommunities := fetchCommunities(c.services.DB)
	communities := getCommunities(dbCommunities)
	ctx.JSON(http.StatusOK, communities)
}

func (c *Controller) ShowBlogs(ctx *gin.Context) {
	page, err := strconv.ParseUint(ctx.Query("page"), 10, 32)
	if err != nil {
		//noinspection GoUnhandledErrorResult
		ctx.Error(err)
	}
	defaultLimit := strconv.Itoa(int(c.services.Config.BlogsInPage))
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", defaultLimit), 10, 32)
	if err != nil {
		//noinspection GoUnhandledErrorResult
		ctx.Error(err)
	}
	sort := ctx.DefaultQuery("sort", "update")
	if len(ctx.Errors) != 0 {
		utils.ShowErrors(ctx)
		return
	}
	offset := limit * (page - 1)
	dbBlogs := fetchBlogs(c.services.DB, uint32(limit), uint32(offset), sort)
	blogs := getBlogs(dbBlogs)
	ctx.JSON(http.StatusOK, blogs)
}

func (c *Controller) ShowBlogArticles(ctx *gin.Context) {
	blogID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		//noinspection GoUnhandledErrorResult
		ctx.Error(err)
	}
	page, err := strconv.ParseUint(ctx.Query("page"), 10, 32)
	if err != nil {
		//noinspection GoUnhandledErrorResult
		ctx.Error(err)
	}
	defaultLimit := strconv.Itoa(int(c.services.Config.BlogTopicsInPage))
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", defaultLimit), 10, 32)
	if err != nil {
		//noinspection GoUnhandledErrorResult
		ctx.Error(err)
	}
	offset := limit * (page - 1)
	dbBlogTopics := fetchBlogTopics(c.services.DB, uint32(blogID), uint32(limit), uint32(offset))
	articles := getBlogArticles(dbBlogTopics, c.services.Config)
	ctx.JSON(http.StatusOK, articles)
}
