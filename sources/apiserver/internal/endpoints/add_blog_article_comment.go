package endpoints

import (
	"fantlab/core/converters"
	"fantlab/core/db"
	"fantlab/pb"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) AddBlogArticleComment(r *http.Request) (int, proto.Message) {
	var params struct {
		// id статьи
		ArticleId uint64 `http:"id,path"`
		// текст комментария (непустой)
		Comment string `http:"comment,form"`
		// id родительского комментария (0, если комментарий 1-го уровня вложенности)
		ParentCommentId uint64 `http:"parent_comment_id,form"`
	}

	api.bindParams(&params, r)

	if params.ArticleId == 0 {
		return api.badParam("id")
	}

	article, err := api.services.DB().FetchBlogTopic(r.Context(), params.ArticleId)

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

	blog, err := api.services.DB().FetchBlog(r.Context(), article.BlogId)

	if err != nil {
		if db.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: "Блог не существует",
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if blog.IsClose == 1 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Блог закрыт",
		}
	}

	// В отличие от форума, здесь нет предварительного форматирования. Это позволяет не только навтыкать массу пробельных
	// символов (мелочь), но и написать модераторское сообщение, будучи самым обычным пользователем (достаточно заключить
	// текст в теги `moder`). https://github.com/parserpro/fantlab/issues/976
	commentLength := uint64(len(params.Comment))

	if commentLength == 0 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Текст комментария пустой",
		}
	}

	userId := api.getUserId(r)

	isUserReadOnly, err := api.services.DB().FetchIsUserReadOnly(r.Context(), userId, article.TopicId, article.BlogId)

	if err != nil && !db.IsNotFoundError(err) {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if isUserReadOnly {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Вам запрещено писать комментарии в этот блог/рубрику",
		}
	}

	articleUserId := article.UserId

	var parentCommentUserId uint64

	if params.ParentCommentId > 0 {
		parentComment, err := api.services.DB().FetchBlogTopicComment(r.Context(), params.ParentCommentId)

		if err != nil {
			if db.IsNotFoundError(err) {
				return http.StatusNotFound, &pb.Error_Response{
					Status:  pb.Error_NOT_FOUND,
					Context: strconv.FormatUint(params.ParentCommentId, 10),
				}
			}

			return http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			}
		}

		// В Perl-бэке такой проверки нет, так что через прямой вызов API можно написать ответ на собственный
		// комментарий. https://github.com/parserpro/fantlab/issues/977
		if parentComment.UserId == userId {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status:  pb.Error_ACTION_FORBIDDEN,
				Context: "Нельзя написать ответ на собственный комментарий",
			}
		}

		parentCommentUserId = parentComment.UserId
	}

	dbComment, err := api.services.DB().InsertBlogTopicComment(r.Context(), userId, article.TopicId, articleUserId,
		params.ParentCommentId, parentCommentUserId, params.Comment, api.services.AppConfig().BlogArticleCommentsInPage)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// Инвалидируем кеши подписчиков, автора родительского комментария и автора статьи. В случае ошибки запрос не фейлим.

	excludedUserIds := []uint64{userId, parentCommentUserId, articleUserId}
	subscribers, _ := api.services.DB().FetchBlogTopicSubscribers(r.Context(), article.TopicId, excludedUserIds)

	for _, subscriber := range subscribers {
		_ = api.services.DeleteUserCache(r.Context(), subscriber)
	}

	_ = api.services.DeleteUserCache(r.Context(), parentCommentUserId)
	_ = api.services.DeleteUserCache(r.Context(), articleUserId)

	commentResponse := converters.GetBlogArticleComment(dbComment, api.services.AppConfig())

	return http.StatusOK, commentResponse
}
