package endpoints

import (
	"fantlab/server/internal/pb"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func (api *API) UnsubscribeArticle(r *http.Request) (int, proto.Message) {
	var params struct {
		// айди статьи
		ArticleId uint64 `http:"id,path"`
	}

	api.bindParams(&params, r)

	if params.ArticleId == 0 {
		return api.badParam("id")
	}

	userId := api.getUserId(r)

	isDbTopicSubscribed, err := api.services.DB().FetchBlogTopicSubscribed(r.Context(), params.ArticleId, userId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if !isDbTopicSubscribed {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "already unsubscribed",
		}
	}

	err = api.services.DB().UpdateBlogTopicUnsubscribed(r.Context(), params.ArticleId, userId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
