package genresapi

import (
	"fantlab/pb"
	"fantlab/shared"
	"fantlab/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	services *shared.Services
}

func NewController(services *shared.Services) *Controller {
	return &Controller{services: services}
}

func (c *Controller) ShowGenres(ctx *gin.Context) {
	dbResponse, err := c.services.DB.FetchGenres()

	if err != nil {
		utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		})
		return
	}

	response := getGenres(dbResponse)

	utils.ShowProto(ctx, http.StatusOK, response)
}
