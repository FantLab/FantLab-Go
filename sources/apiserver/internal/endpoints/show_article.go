package endpoints

import (
	"fantlab/core/app"
	"fantlab/core/converters"
	"fantlab/core/db"
	"fantlab/pb"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
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
		if db.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.ArticleId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	err = api.services.SetBlogArticleLikeCountCache(r.Context(), dbTopic.TopicId, dbTopic.LikesCount)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	viewCounts := api.services.BlogTopicsViewCount(r.Context(), []uint64{dbTopic.TopicId})

	// TODO Переделать
	attachments, _ := app.GetBlogArticleAttachments(dbTopic.TopicId)
	files, _ := api.services.GetMinioFiles(r.Context(), app.BlogArticleFileGroup, dbTopic.TopicId)
	attachments = append(attachments, files...)

	article := converters.GetArticle(dbTopic, viewCounts[0], attachments, api.services.AppConfig())
	return http.StatusOK, article
}
