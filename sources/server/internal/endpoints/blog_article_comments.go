package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/converters"
	"net/http"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
)

func (api *API) BlogArticleComments(r *http.Request) (int, proto.Message) {
	params := struct {
		// айди статьи
		ArticleId uint64 `http:"id,path"`
		// дата, после которой искать сообщения (в формате RFC3339)
		After string `http:"after,query"`
		// кол-во комментариев верхнего уровня (по умолчанию - 10, [5, 20])
		Count uint64 `http:"count,query"`
		// Сортировка (asc, dec, по умолчанию - asc)
		Sort string `http:"sort,query"`
	}{
		Count: api.config.BlogArticleCommentsInPage,
		Sort:  "asc",
	}

	api.bindParams(&params, r)

	if params.ArticleId == 0 {
		return api.badParam("id")
	}
	if params.Count < 5 || params.Count > 20 {
		return api.badParam("count")
	}

	sort := "ASC"

	if strings.ToLower(params.Sort) == "desc" {
		sort = "DESC"
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

	if err != nil && !dbtools.IsNotFoundError(err) {
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
		sort,
		uint8(params.Count),
	)

	if err != nil && !dbtools.IsNotFoundError(err) {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	response := converters.GetBlogArticleComments(comments, count, api.config)

	return http.StatusOK, response
}
