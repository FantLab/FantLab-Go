package endpoints

import (
	"net/http"
	"strings"

	"fantlab/api/internal/endpoints/internal/datahelpers"
	"fantlab/helpers"
	"fantlab/pb"

	"github.com/golang/protobuf/proto"
)

func (api *API) ShowForums(r *http.Request) (int, proto.Message) {
	availableForums := api.config.DefaultAccessToForums

	userId := api.getUserId(r)

	if userId > 0 {
		availableForumsString, err := api.services.DB().FetchAvailableForums(r.Context(), userId)

		if err != nil {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			}
		}

		availableForums = helpers.ParseUints(strings.Split(availableForumsString, ","))

		if availableForums == nil {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			}
		}
	}

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

	forumBlocks := datahelpers.GetForumBlocks(dbForums, dbModerators, api.config)
	return http.StatusOK, forumBlocks
}
