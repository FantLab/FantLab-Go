package endpoints

import (
	"fantlab/api/internal/endpoints/internal/datahelpers"
	"fantlab/dbtools"
	"fantlab/helpers"
	"fantlab/pb"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
)

func (api *API) SetWorkGenres(r *http.Request) (int, proto.Message) {
	var params struct {
		// айди произведения
		WorkId uint64 `http:"id,path"`
		// айди жанров, разделённые запятыми
		GenredIds string `http:"genres,form"`
	}

	api.bindParams(&params, r)

	if params.WorkId == 0 {
		return api.badParam("id")
	}

	genreIds := helpers.ParseUints(strings.Split(params.GenredIds, ","))

	if genreIds == nil {
		return api.badParam("genres")
	}

	// проверяем что ворк существует

	_, err := api.services.DB().WorkExists(r.Context(), params.WorkId)

	{
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.WorkId, 10),
			}
		}

		if err != nil {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			}
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

	userId := api.getUserId(r)

	err = api.services.DB().GenreVote(r.Context(), params.WorkId, userId, genreIdsWithParents)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// возвращаем OK

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
