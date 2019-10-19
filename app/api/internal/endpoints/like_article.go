package endpoints

import (
	"fantlab/dbtools"
	"fantlab/pb"
	"net/http"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
)

func (api *API) LikeArticle(r *http.Request) (int, proto.Message) {
	userId := api.getUserId(r)

	articleId, err := uintURLParam(r, "id")

	if err != nil {
		return http.StatusBadRequest, &pb.Error_Response{
			Status:  pb.Error_INVALID_PARAMETER,
			Context: "id",
		}
	}

	dbTopic, err := api.services.DB().FetchBlogTopic(r.Context(), articleId)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(articleId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if dbTopic.UserId == userId {
		return http.StatusUnauthorized, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "your own article",
		}
	}

	ok, err := api.services.DB().IsBlogTopicLiked(r.Context(), articleId, userId)

	if err != nil && !dbtools.IsNotFoundError(err) {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if ok {
		return http.StatusUnauthorized, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "already liked",
		}
	}

	err = api.services.DB().LikeBlogTopic(r.Context(), time.Now(), articleId, userId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	dbTopicLikeCount, err := api.services.DB().FetchBlogTopicLikeCount(r.Context(), articleId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Blog_BlogArticleLikeResponse{
		LikeCount: dbTopicLikeCount,
	}
}
