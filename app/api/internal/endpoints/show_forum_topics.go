package endpoints

import (
	"net/http"
	"strconv"
	"strings"

	"fantlab/api/internal/endpoints/internal/datahelpers"
	"fantlab/dbtools"
	"fantlab/helpers"
	"fantlab/pb"

	"github.com/golang/protobuf/proto"
)

func (api *API) ShowForumTopics(r *http.Request) (int, proto.Message) {
	forumID, err := uintURLParam(r, "id")

	if err != nil {
		return http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "id",
		}
	}

	page, err := uintQueryParam(r, "page", 1)

	if err != nil {
		return http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "page",
		}
	}

	limit, err := uintQueryParam(r, "limit", api.config.ForumTopicsInPage)

	if err != nil || !helpers.IsValidLimit(limit) {
		return http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "limit",
		}
	}

	availableForums := api.config.DefaultAccessToForums

	userId := api.getUserId(r)

	if userId > 0 {
		availableForumsString, err := api.services.DB().FetchAvailableForums(r.Context(), userId)

		if err != nil {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			}
		}

		availableForums, err = helpers.ParseUints(strings.Split(availableForumsString, ","), 10, 64)

		if err != nil {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			}
		}
	}

	offset := limit * (page - 1)

	dbResponse, err := api.services.DB().FetchForumTopics(
		r.Context(),
		availableForums,
		forumID,
		limit,
		offset,
	)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(forumID, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	forumTopics := datahelpers.GetForumTopics(dbResponse, page, limit, api.config)
	return http.StatusOK, forumTopics
}
