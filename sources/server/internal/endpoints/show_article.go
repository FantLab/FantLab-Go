package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/converters"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"
)

func (api *API) ShowArticle(r *http.Request) (int, proto.Message) {
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

	err = api.services.SetBlogArticleLikeCountCache(r.Context(), params.ArticleId, dbTopic.LikesCount)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	viewCount := api.services.BlogTopicsViewCount(r.Context(), []uint64{params.ArticleId})

	article := converters.GetArticle(dbTopic, viewCount[0], api.config)
	return http.StatusOK, article
}
