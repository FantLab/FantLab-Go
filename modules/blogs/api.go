package blogsapi

import (
	"fantlab/shared"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	services *shared.Services
}

func NewController(services *shared.Services) *Controller {
	return &Controller{
		services: services,
	}
}

func (c *Controller) ShowCommunities(ctx *gin.Context) {
	dbCommunities := fetchCommunities(c.services.DB)
	communities := getCommunities(dbCommunities)
	ctx.JSON(http.StatusOK, communities)
}
