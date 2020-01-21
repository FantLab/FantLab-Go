package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"
)

func (api *API) SubscribeForumTopic(r *http.Request) (int, proto.Message) {
	var params struct {
		// айди темы
		TopicId uint64 `http:"id,path"`
	}

	api.bindParams(&params, r)

	if params.TopicId == 0 {
		return api.badParam("id")
	}

	userId := api.getUserId(r)

	availableForums := api.getAvailableForums(r)

	_, err := api.services.DB().FetchForumTopic(r.Context(), availableForums, params.TopicId)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.TopicId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	isDbTopicSubscribed, err := api.services.DB().FetchForumTopicSubscribed(r.Context(), params.TopicId, userId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if isDbTopicSubscribed {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "already subscribed",
		}
	}

	err = api.services.DB().UpdateForumTopicSubscribed(r.Context(), params.TopicId, userId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
