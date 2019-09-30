package endpoints

import (
	"fantlab/pb"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func (api *API) UnsubscribeArticle(r *http.Request) (int, proto.Message) {
	userId := api.getUserId(r)

	articleId, err := uintURLParam(r, "id")

	if err != nil {
		return http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "id",
		}
	}

	isDbTopicSubscribed, err := api.services.DB().FetchBlogTopicSubscribed(uint32(articleId), uint32(userId))

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

	_, err = api.services.DB().UpdateBlogTopicUnsubscribed(uint32(articleId), uint32(userId))

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Blog_BlogSubscriptionResponse{
		IsSubscribed: false,
	}
}
