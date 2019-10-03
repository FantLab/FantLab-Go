package endpoints

import (
	"fantlab/api/internal/endpoints/internal/datahelpers"
	"fantlab/pb"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func (api *API) ShowForums(r *http.Request) (int, proto.Message) {
	dbForums, err := api.services.DB().FetchForums(api.config.DefaultAccessToForums)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	dbModerators, err := api.services.DB().FetchModerators()

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	forumBlocks := datahelpers.GetForumBlocks(dbForums, dbModerators, api.config)
	return http.StatusOK, forumBlocks
}
