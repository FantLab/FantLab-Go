package endpoints

import (
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) DeleteBlogArticleComment(r *http.Request) (int, proto.Message) {
	var params struct {
		// id комментария
		CommentId uint64 `http:"id,path"`
	}

	api.bindParams(&params, r)

	if params.CommentId == 0 {
		return api.badParam("id")
	}

	comment, err := api.services.DB().FetchBlogTopicComment(r.Context(), params.CommentId)

	if err != nil {
		if db.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.CommentId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	article, err := api.services.DB().FetchBlogTopic(r.Context(), comment.TopicId)

	if err != nil {
		if db.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: "Статья не существует",
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

	userId := api.getUserId(r)

	// NOTE Пропущен весь хардкод касательно id отдельных юзеров, обработка is_referee (заданы в Auth.pm) и
	// can_link_blogarticle_to_work (из main.cfg). Все они считаются модераторами любых блогов.

	userIsCommunityModerator, err := api.services.DB().FetchUserIsCommunityModerator(r.Context(), userId, blog.BlogId, article.TopicId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// В отличие от форума, нет ограничения на время редактирования сообщения
	if !(comment.UserId == userId || blog.UserId == userId || userIsCommunityModerator) {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Вы не можете удалить данный комментарий",
		}
	}

	err = api.services.DB().DeleteBlogTopicComment(r.Context(), comment.MessageId, comment.ParentMessageId, article.TopicId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	helpers.DeleteBlogCommentTextCache(comment.MessageId)

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
