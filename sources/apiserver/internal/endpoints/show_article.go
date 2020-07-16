package endpoints

import (
	"fantlab/core/app"
	"fantlab/core/converters"
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
	"net/http"
	"sort"
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

	attachments, _ := helpers.GetBlogArticleAttachments(dbTopic.TopicId)
	files, _ := api.services.GetFiles(r.Context(), app.BlogArticleFileGroup, dbTopic.TopicId)
	attachments = append(attachments, files...)
	sort.Slice(attachments, func(i, j int) bool {
		return attachments[i].Name < attachments[j].Name
	})

	article := converters.GetArticle(dbTopic, viewCounts[0], attachments, api.services.AppConfig())
	return http.StatusOK, article
}
