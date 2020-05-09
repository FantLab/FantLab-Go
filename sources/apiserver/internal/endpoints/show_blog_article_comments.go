package endpoints

import (
	"fantlab/core/converters"
	"fantlab/core/db"
	"fantlab/pb"
	"net/http"
	"time"

	"google.golang.org/protobuf/proto"
)

func (api *API) BlogArticleComments(r *http.Request) (int, proto.Message) {
	params := struct {
		// id статьи
		ArticleId uint64 `http:"id,path"`
		// дата, после которой искать сообщения (в формате RFC3339)
		After string `http:"after,query"`
		// кол-во комментариев верхнего уровня (по умолчанию - 10, [5, 20])
		Count uint64 `http:"count,query"`
		// порядок выдачи (0 - от новых к старым, 1 - наоборот; по умолчанию - 0)
		SortAsc uint8 `http:"sortAsc,query"`
	}{
		Count:   api.services.AppConfig().BlogArticleCommentsInPage,
		SortAsc: 0,
	}

	api.bindParams(&params, r)

	if params.ArticleId == 0 {
		return api.badParam("id")
	}
	if params.Count < 5 || params.Count > 20 {
		return api.badParam("count")
	}
	if !(params.SortAsc == 0 || params.SortAsc == 1) {
		return api.badParam("sortAsc")
	}

	var err error
	after := time.Unix(0, 0)

	if params.After != "" {
		after, err = time.Parse(time.RFC3339, params.After)
	}

	if err != nil {
		return api.badParam("after")
	}

	count, err := api.services.DB().FetchBlogTopicCommentsCount(r.Context(), params.ArticleId)

	if err != nil && !db.IsNotFoundError(err) {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if count == 0 {
		return http.StatusOK, &pb.Blog_BlogArticleCommentsResponse{}
	}

	comments, err := api.services.DB().FetchBlogTopicComments(
		r.Context(),
		params.ArticleId,
		after,
		params.SortAsc == 1,
		uint8(params.Count),
	)

	if err != nil && !db.IsNotFoundError(err) {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	response := converters.GetBlogArticleComments(comments, count, api.services.AppConfig())

	return http.StatusOK, response
}
