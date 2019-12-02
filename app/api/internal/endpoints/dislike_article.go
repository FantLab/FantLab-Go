package endpoints

import (
	"fantlab/dbtools"
	"fantlab/pb"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"
)

func (api *API) DislikeArticle(r *http.Request) (int, proto.Message) {
	var params struct {
		// айди статьи
		ArticleId uint64 `http:"id,path"`
	}

	api.bindParams(&params, r)

	if params.ArticleId == 0 {
		return api.badParam("id")
	}

	dbTopic, err := api.services.DB().FetchBlogTopic(r.Context(), params.ArticleId)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.ArticleId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	userId := api.getUserId(r)

	if dbTopic.UserId == userId {
		return http.StatusUnauthorized, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "your own article",
		}
	}

	ok, err := api.services.DB().IsBlogTopicLiked(r.Context(), params.ArticleId, userId)

	if err != nil && !dbtools.IsNotFoundError(err) {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if !ok {
		return http.StatusUnauthorized, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "already disliked",
		}
	}

	err = api.services.DB().DislikeBlogTopic(r.Context(), params.ArticleId, userId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	likeCount, err := api.services.DB().FetchBlogTopicLikeCount(r.Context(), params.ArticleId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Blog_BlogArticleLikeResponse{
		LikeCount: likeCount,
	}
}
