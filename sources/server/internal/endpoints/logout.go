package endpoints

import (
	"fantlab/server/internal/pb"
	"net/http"

	"github.com/golang/protobuf/proto"
)

// Удаляет текущую сессию пользователя
func (api *API) Logout(r *http.Request) (int, proto.Message) {
	sid := api.getSession(r)

	err := api.services.DB().DeleteSession(r.Context(), sid)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	_ = api.services.Cache().DeleteSession(r.Context(), sid)

	return http.StatusOK, &pb.Common_SuccessResponse{}
}