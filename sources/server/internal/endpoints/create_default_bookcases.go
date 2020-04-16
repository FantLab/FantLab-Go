package endpoints

import (
	"fantlab/pb"
	"fantlab/server/internal/converters"
	"google.golang.org/protobuf/proto"
	"net/http"
)

func (api *API) CreateDefaultBookcases(r *http.Request) (int, proto.Message) {
	user := api.getUser(r)

	dbBookcases, err := api.services.DB().FetchBookcases(r.Context(), user.UserId, true)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if len(dbBookcases) > 0 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
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
