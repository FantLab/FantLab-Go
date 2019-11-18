package endpoints

import (
	"fantlab/api/internal/endpoints/internal/datahelpers"
	"fantlab/dbtools"
	"fantlab/helpers"
	"fantlab/pb"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"
)

func (api *API) SetWorkGenres(r *http.Request) (int, proto.Message) {
	userId := api.getUserId(r)

	// валидируем идентификатор ворка

	workId, err := api.uintURLParam(r, "id")

	if err != nil {
		return http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "id",
		}
	}

	// проверяем что ворк существует

	_, err = api.services.DB().WorkExists(r.Context(), workId)

	{
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(workId, 10),
			}
		}

		if err != nil {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			}
		}
	}

	// валидируем идентификаторы жанров

	genreIds, err := helpers.ParseUints(r.PostForm["genres"], 10, 32)

	if err != nil || len(genreIds) == 0 {
		return http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "genres",
		}
	}

	// получаем все жанры из базы

	dbResponse, err := api.services.DB().FetchGenreIds(r.Context())

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// создаем дерево жанров

	genreTree := datahelpers.MakeGenreTree(dbResponse)

	// проверяем что выбраны жанры из обязательных групп

	err = datahelpers.CheckRequiredGroupsForGenreIds(genreIds, genreTree)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status:  pb.Error_VALIDATION_FAILED,
			Context: err.Error(),
		}
	}

	// получаем идентификаторы всех выбранных жанров + родительские

	genreIdsWithParents := datahelpers.SelectGenreIdsWithParents(genreIds, genreTree)

	// сохраняем выбор в базе

	err = api.services.DB().GenreVote(r.Context(), workId, userId, genreIdsWithParents)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// возвращаем OK

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
