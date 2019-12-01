package endpoints

import (
	"fantlab/dbtools"
	"fantlab/pb"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"
)

func (api *API) UnsubscribeCommunity(r *http.Request) (int, proto.Message) {
	var params struct {
		// айди сообщества
		CommunityId uint64 `http:"id,path"`
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

	isDbCommunitySubscribed, err := api.services.DB().FetchBlogSubscribed(r.Context(), params.CommunityId, userId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if !isDbCommunitySubscribed {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "already unsubscribed",
		}
	}

	err = api.services.DB().UpdateBlogUnsubscribed(r.Context(), params.CommunityId, userId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Blog_BlogSubscriptionResponse{
		IsSubscribed: false,
	}
}
