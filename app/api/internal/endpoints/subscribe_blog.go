package endpoints

import (
	"fantlab/dbtools"
	"fantlab/pb"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"
)

func (api *API) SubscribeBlog(r *http.Request) (int, proto.Message) {
	var params struct {
		// айди блога
		BlogId uint64 `http:"id,path"`
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

	isDbBlogSubscribed, err := api.services.DB().FetchBlogSubscribed(r.Context(), params.BlogId, userId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if isDbBlogSubscribed {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "already subscribed",
		}
	}

	err = api.services.DB().UpdateBlogSubscribed(r.Context(), params.BlogId, userId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
