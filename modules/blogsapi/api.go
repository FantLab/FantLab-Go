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
	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", "50"), 10, 32) // todo get from config
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
