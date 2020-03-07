package endpoints

import (
	"fantlab/pb"
	"fantlab/server/internal/converters"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func (api *API) GetUserWorkGenres(r *http.Request) (int, proto.Message) {
	var params struct {
		// айди произведения
		WorkId uint64 `http:"id,path"`
	}

	api.bindParams(&params, r)

	if params.WorkId == 0 {
		return api.badParam("id")
	}

	userId := api.getUserId(r)

	// получаем список всех жанров

	genreTree := api.services.GetGenreTree(r.Context())

	if genreTree == nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// получаем айдишки жанров, выбранных пользователем

	genreIds, err := api.services.DB().GetUserWorkGenreIds(r.Context(), params.WorkId, userId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// формируем ответ

	response := converters.GetGenres(genreTree, genreIds, nil)

	// успех

	return http.StatusOK, response
}
