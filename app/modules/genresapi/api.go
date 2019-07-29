package genresapi

import (
	"fantlab/pb"
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

func (c *Controller) SetWorkGenres(ctx *gin.Context) {
	// проверяем что юзер может голосовать

	userId := ctx.GetInt64(gin.AuthUserKey)

	if userId == 0 {
		utils.ShowProto(ctx, http.StatusUnauthorized, &pb.Error_Response{
			Status: pb.Error_INVALID_SESSION,
		})
		return
	}

	// валидируем идентификатор ворка

	workId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		utils.ShowProto(ctx, http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "id",
		})
		return
	}

	// проверяем что ворк существует

	err = c.services.DB.WorkExists(workId)

	if err != nil {
		if utils.IsRecordNotFoundError(err) {
			utils.ShowProto(ctx, http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(workId, 10),
			})
		} else {
			utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			})
		}
		return
	}

	// валидируем идентификаторы жанров

	genreIds, err := utils.ParseUints(ctx.PostFormArray("genres"), 10, 32)

	if err != nil || len(genreIds) == 0 {
		utils.ShowProto(ctx, http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "genres",
		})
		return
	}

	// получаем все жанры из базы

	dbResponse, err := c.services.DB.FetchGenreIds()

	if err != nil {
		utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		})
		return
	}

	// создаем дерево жанров

	genreTree := makeGenreTree(dbResponse)

	// проверяем что выбраны жанры из обязательных групп

	err = checkRequiredGroupsForGenreIds(genreIds, genreTree)

	if err != nil {
		utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
			Status:  pb.Error_VALIDATION_FAILED,
			Context: err.Error(),
		})
		return
	}

	// получаем идентификаторы всех выбранных жанров + родительские

	genreIdsWithParents := selectGenreIdsWithParents(genreIds, genreTree)

	// сохраняем выбор в базе

	err = c.services.DB.GenreVote(workId, userId, genreIdsWithParents)

	if err != nil {
		utils.ShowProto(ctx, http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		})
		return
	}

	// возвращаем OK

	utils.ShowProto(ctx, http.StatusOK, &pb.Common_SuccessResponse{})
}
