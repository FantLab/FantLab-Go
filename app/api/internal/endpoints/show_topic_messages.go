package endpoints

import (
	"fantlab/api/internal/endpoints/internal/datahelpers"
	"fantlab/dbtools"
	"fantlab/helpers"
	"fantlab/pb"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
)

func (api *API) ShowTopicMessages(r *http.Request) (int, proto.Message) {
	topicID, err := uintURLParam(r, "id")

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

	limit, err := uintQueryParam(r, "limit", api.config.ForumMessagesInPage)

	if err != nil || !helpers.IsValidLimit(limit) {
		return http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "limit",
		}
	}

	sortDirection := strings.ToUpper(queryParam(r, "order", "asc"))
	if sortDirection != "DESC" {
		sortDirection = "ASC"
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

		availableForums, err = helpers.ParseUints(strings.Split(availableForumsString, ","), 10, 32)

		if err != nil {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			}
		}
	}

	offset := limit * (page - 1)

	dbResponse, err := api.services.DB().FetchTopicMessages(
		r.Context(),
		availableForums,
		uint32(topicID),
		uint32(limit),
		uint32(offset),
		sortDirection,
	)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(topicID, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	topicMessages := datahelpers.GetTopic(dbResponse, uint32(page), uint32(limit), api.config)
	return http.StatusOK, topicMessages
}
