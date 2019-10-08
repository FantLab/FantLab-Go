package endpoints

import (
	"fantlab/api/internal/endpoints/internal/datahelpers"
	"fantlab/pb"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func (api *API) ShowCommunities(r *http.Request) (int, proto.Message) {
	dbCommunities, err := api.services.DB().FetchCommunities(r.Context(), )

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	communities := datahelpers.GetCommunities(dbCommunities, api.config)
	return http.StatusOK, communities
}
