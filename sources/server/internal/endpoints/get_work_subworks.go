package endpoints

import (
	"fantlab/pb"
	"fantlab/server/internal/converters"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func (api *API) GetWorkSubWorks(r *http.Request) (int, proto.Message) {
	var params struct {
		// айди произведения
		WorkId uint64 `http:"id,path"`
	}

	api.bindParams(&params, r)

	if params.WorkId == 0 {
		return api.badParam("id")
	}

	// получаем иерархию всех дочерних произведений в виде списка

	children, err := api.services.DB().GetWorkChildren(r.Context(), params.WorkId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// формируем ответ

	response := converters.GetSubWorks(params.WorkId, children)

	// успех

	return http.StatusOK, response
}
