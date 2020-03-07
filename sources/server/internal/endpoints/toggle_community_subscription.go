package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) ToggleCommunitySubscription(r *http.Request) (int, proto.Message) {
	var params struct {
		// айди сообщества
		CommunityId uint64 `http:"id,path"`
		// подписаться - true, отписаться - false
		Subscribe bool `http:"subscribe,form"`
	}

	api.bindParams(&params, r)

	if params.CommunityId == 0 {
		return api.badParam("id")
	}

	_, err := api.services.DB().FetchCommunity(r.Context(), params.CommunityId)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.CommunityId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	userId := api.getUserId(r)

	if params.Subscribe {
		err = api.services.DB().UpdateBlogSubscribed(r.Context(), params.CommunityId, userId)
	} else {
		err = api.services.DB().UpdateBlogUnsubscribed(r.Context(), params.CommunityId, userId)
	}

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
