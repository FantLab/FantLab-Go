package endpoints

import (
	"net/http"
	"strconv"
	"strings"

	"fantlab/dbtools"
	"fantlab/helpers"
	"fantlab/pb"

	"github.com/golang/protobuf/proto"
)

func (api *API) UnsubscribeForumTopic(r *http.Request) (int, proto.Message) {
	topicId, err := uintURLParam(r, "id")

	if err != nil {
		return http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "id",
		}
	}

	userId := api.getUserId(r)

	availableForumsString, err := api.services.DB().FetchAvailableForums(r.Context(), userId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	availableForums, err := helpers.ParseUints(strings.Split(availableForumsString, ","), 10, 64)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	_, err = api.services.DB().FetchForumTopic(r.Context(), availableForums, topicId)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(topicId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	isDbTopicSubscribed, err := api.services.DB().FetchForumTopicSubscribed(r.Context(), topicId, userId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if !isDbTopicSubscribed {
		return http.StatusUnauthorized, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "already unsubscribed",
		}
	}

	err = api.services.DB().UpdateForumTopicUnsubscribed(r.Context(), topicId, userId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Forum_ForumTopicSubscriptionResponse{
		IsSubscribed: false,
	}
}
