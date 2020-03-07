package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) ToggleBlogSubscription(r *http.Request) (int, proto.Message) {
	var params struct {
		// айди блога
		BlogId uint64 `http:"id,path"`
		// подписаться - true, отписаться - false
		Subscribe bool `http:"subscribe,form"`
	}

	api.bindParams(&params, r)

	if params.BlogId == 0 {
		return api.badParam("id")
	}

	dbBlog, err := api.services.DB().FetchBlog(r.Context(), params.BlogId)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.BlogId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	userId := api.getUserId(r)

	if dbBlog.UserId == userId {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "your own blog",
		}
	}

	if params.Subscribe {
		err = api.services.DB().UpdateBlogSubscribed(r.Context(), params.BlogId, userId)
	} else {
		err = api.services.DB().UpdateBlogUnsubscribed(r.Context(), params.BlogId, userId)
	}

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
