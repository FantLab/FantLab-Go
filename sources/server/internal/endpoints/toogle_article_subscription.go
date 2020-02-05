package endpoints

import (
	"fantlab/pb"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func (api *API) ToggleArticleSubscription(r *http.Request) (int, proto.Message) {
	var params struct {
		// айди статьи
		ArticleId uint64 `http:"id,path"`
		// подписаться - true, отписаться - false
		Subscribe bool `http:"subscribe,form"`
	}

	api.bindParams(&params, r)

	if params.ArticleId == 0 {
		return api.badParam("id")
	}

	userId := api.getUserId(r)

	var err error

	if params.Subscribe {
		err = api.services.DB().UpdateBlogTopicSubscribed(r.Context(), params.ArticleId, userId)
	} else {
		err = api.services.DB().UpdateBlogTopicUnsubscribed(r.Context(), params.ArticleId, userId)
	}

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
