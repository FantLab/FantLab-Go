package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"github.com/golang/protobuf/proto"
	"net/http"
	"strconv"
)

func (api *API) ToggleArticleLike(r *http.Request) (int, proto.Message) {
	var params struct {
		// айди статьи
		ArticleId uint64 `http:"id,path"`
		// лайк - true, dislike - false
		Like bool `http:"like,form"`
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
			Context: "Нельзя лайкнуть собственную статью",
		}
	}

	if params.Like {
		err = api.services.DB().LikeBlogTopic(r.Context(), params.ArticleId, userId)
	} else {
		err = api.services.DB().DislikeBlogTopic(r.Context(), params.ArticleId, userId)
	}

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	dbTopicLikeCount, err := api.services.DB().FetchBlogTopicLikeCount(r.Context(), params.ArticleId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	err = api.services.SetBlogArticleLikeCountCache(r.Context(), params.ArticleId, dbTopicLikeCount)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Blog_BlogArticleLikeResponse{
		LikeCount: dbTopicLikeCount,
	}
}
