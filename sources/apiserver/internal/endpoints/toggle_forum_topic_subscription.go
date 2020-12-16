package endpoints

import (
	"fantlab/pb"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) ToggleForumTopicSubscription(r *http.Request) (int, proto.Message) {
	var params struct {
		// айди темы
		TopicId uint64 `http:"id,path"`
		// подписаться - true, отписаться - false
		Subscribe bool `http:"subscribe,form"`
	}

	api.bindParams(&params, r)

	if params.TopicId == 0 {
		return api.badParam("id")
	}

	availableForums := api.getAvailableForums(r)

	isTopicExists, err := api.services.DB().FetchForumTopicExists(r.Context(), params.TopicId, availableForums)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if !isTopicExists {
		return http.StatusNotFound, &pb.Error_Response{
			Status:  pb.Error_NOT_FOUND,
			Context: strconv.FormatUint(params.TopicId, 10),
		}
	}

	userId := api.getUserId(r)

	if params.Subscribe {
		err = api.services.DB().UpdateForumTopicSubscribed(r.Context(), params.TopicId, userId)
	} else {
		err = api.services.DB().UpdateForumTopicUnsubscribed(r.Context(), params.TopicId, userId)
	}

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
