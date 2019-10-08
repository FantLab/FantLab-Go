package endpoints

import (
	"fantlab/api/internal/endpoints/internal/datahelpers"
	"fantlab/pb"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func (api *API) ShowGenres(r *http.Request) (int, proto.Message) {
	dbResponse, err := api.services.DB().FetchGenres(r.Context())

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	response := datahelpers.GetGenres(dbResponse)

	return http.StatusOK, response
}
