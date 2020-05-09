package endpoints

import (
	"fantlab/core/converters"
	"fantlab/pb"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func (api *API) ShowCommunities(r *http.Request) (int, proto.Message) {
	dbCommunities, err := api.services.DB().FetchCommunities(r.Context())

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	communities := converters.GetCommunities(dbCommunities, api.services.AppConfig())
	return http.StatusOK, communities
}
