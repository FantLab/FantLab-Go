package endpoints

import (
	"fantlab/pb"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func (api *API) Logout(r *http.Request) (int, proto.Message) {
	sid := api.getSession(r)

	_, err := api.services.DB().DeleteSession(r.Context(), sid)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	_ = api.services.Cache().DeleteSession(sid)

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
