package endpoints

import (
	"fantlab/core/converters"
	"fantlab/pb"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func (api *API) GetWorkSubWorks(r *http.Request) (int, proto.Message) {
	params := struct {
		// айди произведения
		WorkId uint64 `http:"id,path"`
		// глубина дерева (1 - 5, по умолчанию - 4)
		Depth uint8 `http:"depth,query"`
	}{
		Depth: 4,
	}

	api.bindParams(&params, r)

	if params.WorkId == 0 {
		return api.badParam("id")
	}
	if params.Depth < 1 || params.Depth > 5 {
		return api.badParam("depth")
	}

	// получаем иерархию всех дочерних произведений в виде списка

	children, err := api.services.DB().GetWorkChildren(r.Context(), params.WorkId, params.Depth)

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
