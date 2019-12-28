package endpoints

import (
	"fantlab/pb"
	"net/http"

	"github.com/golang/protobuf/proto"
)

// Удаляет текущую сессию пользователя
func (api *API) Logout(r *http.Request) (int, proto.Message) {
	sid := api.getSession(r)

	err := api.services.DeleteSessionById(r.Context(), sid)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
