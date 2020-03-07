package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/base/utils"
	"fantlab/pb"
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/protobuf/proto"
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

	genreIds := utils.ParseUints(strings.Split(params.GenredIds, ","))

	if genreIds == nil {
		return api.badParam("genres")
	}

	userId := api.getUserId(r)

	// проверяем что произведение существует
	{
		workExists, err := api.services.DB().WorkExists(r.Context(), params.WorkId)
		if err != nil && !dbtools.IsNotFoundError(err) {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			}
		}
		if !workExists {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: fmt.Sprintf("Произведение с идентификатором %d не найдено", params.WorkId),
			}
		}
	}

	// проверяем что пользователь выставил оценку произведению
	{
		mark, err := api.services.DB().GetWorkUserMark(r.Context(), params.WorkId, userId)
		if err != nil && !dbtools.IsNotFoundError(err) {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			}
		}
		if mark == 0 {
			return http.StatusForbidden, &pb.Error_Response{
				Status:  pb.Error_ACTION_PERMITTED,
				Context: "Вы еще не оценили это произведение",
			}
		}
	}

	// получаем дерево жанры

	genreTree := api.services.GetGenreTree(r.Context())

	if genreTree == nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// проверяем что выбраны жанры из обязательных групп
	{
		err := genreTree.CheckRequiredGroupsForGenreIds(genreIds)

		if err != nil {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status:  pb.Error_VALIDATION_FAILED,
				Context: err.Error(),
			}
		}
	}

	// получаем идентификаторы всех выбранных жанров + родительские

	genreIdTable := genreTree.SelectGenreIdsWithParents(genreIds)

	// сохраняем выбор в базе
	{
		genreIds = nil

		for id := range genreIdTable {
			genreIds = append(genreIds, id)
		}

		err := api.services.DB().GenreVote(r.Context(), params.WorkId, userId, genreIds)

		if err != nil {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			}
		}
	}

	// сбрасываем кэш юзера (для перла)

	api.services.InvalidateUserCache(r.Context(), userId)

	// успех

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
