package endpoints

import (
	"fantlab/core/converters"
	"fantlab/pb"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func (api *API) ShowForums(r *http.Request) (int, proto.Message) {
	availableForums := api.getAvailableForums(r)

	dbForums, err := api.services.DB().FetchForums(r.Context(), availableForums)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	dbModerators, err := api.services.DB().FetchModerators(r.Context())

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	forumBlocks := converters.GetForumBlocks(dbForums, dbModerators, api.services.AppConfig())
	return http.StatusOK, forumBlocks
}
