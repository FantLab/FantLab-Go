package endpoints

import (
	"fantlab/pb"
	"fantlab/server/internal/converters"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func (api *API) GetWorkClassification(r *http.Request) (int, proto.Message) {
	var params struct {
		// айди произведения
		WorkId uint64 `http:"id,path"`
	}

	api.bindParams(&params, r)

	if params.WorkId == 0 {
		return api.badParam("id")
	}

	// получаем список всех жанров

	genreTree := api.services.GetGenreTree(r.Context())

	if genreTree == nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// получаем кол-во классификаций

	classificationCount, err := api.services.DB().GetWorkClassificationCount(r.Context(), params.WorkId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// получаем распределение голосов по жанрам

	genreVotes, err := api.services.DB().FetchWorkGenreVotes(r.Context(), params.WorkId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// формируем ответ

	response := converters.GetWorkClassification(genreTree, classificationCount, genreVotes)

	// успех

	return http.StatusOK, response
}
