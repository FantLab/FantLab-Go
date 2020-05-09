package endpoints

import (
	"fantlab/core/converters"
	"fantlab/core/db"
	"fantlab/pb"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) EditBlogArticleComment(r *http.Request) (int, proto.Message) {
	var params struct {
		// id комментария
		CommentId uint64 `http:"id,path"`
		// текст комментария (непустой)
		Comment string `http:"comment,form"`
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
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Блог закрыт",
		}
	}

	user := api.getUser(r)

	// TODO: Пропущен весь хардкод касательно id отдельных юзеров, обработка is_referee и can_link_blogarticle_to_work
	//  (заданы через main.cfg, но зачем так сделано - неясно). Все они считаются модераторами любых блогов.

	userIsCommunityModerator, err := api.services.DB().FetchUserIsCommunityModerator(r.Context(), user.UserId, blog.BlogId, article.TopicId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// В отличие от форума, нет ограничения на время редактирования сообщения
	if !(comment.UserId == user.UserId || blog.UserId == user.UserId || userIsCommunityModerator) {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Вы не можете отредактировать данный комментарий",
		}
	}

	// В отличие от форума, здесь нет предварительного форматирования. Это позволяет не только навтыкать массу пробельных
	// символов (мелочь), но и написать модераторское сообщение, будучи самым обычным пользователем (достаточно заключить
	// текст в теги `moder`). https://github.com/parserpro/fantlab/issues/976
	commentLength := uint64(len(params.Comment))

	if commentLength == 0 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Текст комментария пустой",
		}
	}

	dbComment, err := api.services.DB().UpdateBlogTopicComment(r.Context(), comment.MessageId, params.Comment)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// TODO Удалить файловый кеш

	commentResponse := converters.GetBlogArticleComment(dbComment, api.services.AppConfig())

	return http.StatusOK, commentResponse
}
