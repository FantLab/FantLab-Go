package endpoints

import (
	"fantlab/pb"
	"fantlab/server/internal/converters"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func (api *API) ShowGenres(r *http.Request) (int, proto.Message) {
	// получаем список всех жанров

	genreTree := api.services.GetGenreTree(r.Context())

	if genreTree == nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// получаем распределение произведений по жанрам

	workCounts, err := api.services.DB().FetchGenreWorkCounts(r.Context())

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// формируем ответ

	response := converters.GetGenres(genreTree, nil, workCounts)

	// успех

	return http.StatusOK, response
}
