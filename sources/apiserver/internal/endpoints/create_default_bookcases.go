package endpoints

import (
	"fantlab/core/converters"
	"fantlab/pb"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func (api *API) CreateDefaultBookcases(r *http.Request) (int, proto.Message) {
	user := api.getUser(r)

	dbBookcases, err := api.services.DB().FetchAllUserBookcases(r.Context(), user.UserId, true)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if len(dbBookcases) > 0 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Создать первичные книжные полки можно только в том случае, если не создано еще ни одной",
		}
	}

	dbBookcases, err = api.services.DB().InsertDefaultBookcases(r.Context(), user.UserId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	bookcasesResponse := converters.GetBookcases(dbBookcases)

	return http.StatusOK, bookcasesResponse
}
