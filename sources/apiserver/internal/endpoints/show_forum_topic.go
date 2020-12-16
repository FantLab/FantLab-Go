package endpoints

import (
	"fantlab/core/app"
	"fantlab/core/converters"
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) ShowTopicMessages(r *http.Request) (int, proto.Message) {
	params := struct {
		// id темы
		TopicId uint64 `http:"id,path"`
		// номер страницы (по умолчанию - 1)
		Page uint64 `http:"page,query"`
		// кол-во записей на странице (по умолчанию - 20)
		Limit uint64 `http:"limit,query"`
		// порядок выдачи (0 - от новых к старым, 1 - наоборот; по умолчанию - 0)
		SortAsc uint8 `http:"sortAsc,query"`
	}{
		Page:    1,
		Limit:   api.services.AppConfig().ForumMessagesInPage,
		SortAsc: 0,
	}

	api.bindParams(&params, r)

	if params.TopicId == 0 {
		return api.badParam("id")
	}
	if params.Page == 0 {
		return api.badParam("page")
	}
	if !helpers.IsValidLimit(params.Limit) {
		return api.badParam("limit")
	}
	if !(params.SortAsc == 0 || params.SortAsc == 1) {
		return api.badParam("sortAsc")
	}

	availableForums := api.getAvailableForums(r)

	isTopicExists, err := api.services.DB().FetchForumTopicExists(r.Context(), params.TopicId, availableForums)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if !isTopicExists {
		return http.StatusNotFound, &pb.Error_Response{
			Status:  pb.Error_NOT_FOUND,
			Context: strconv.FormatUint(params.TopicId, 10),
		}
	}

	dbTopic, err := api.services.DB().FetchForumTopic(r.Context(), params.TopicId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	user := api.getUser(r)
	var userId uint64
	var userIsForumModerator bool

	if user != nil {
		userId = user.UserId

		userIsForumModerator, err = api.services.DB().FetchUserIsForumModerator(r.Context(), userId, dbTopic.ForumId)

		if err != nil {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			}
		}

		// Если тема не отмодерирована, юзер - не автор темы и не модератор, возвращаем 404
		if dbTopic.Moderated != 1 && dbTopic.UserId != userId && !userIsForumModerator {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(dbTopic.TopicId, 10),
			}
		}
	}

	userCanPerformAdminActions := api.isPermissionGranted(r, pb.Auth_Claims_PERMISSION_CAN_PERFORM_ADMIN_ACTIONS)
	userCanEditOwnForumMessages := api.isPermissionGranted(r, pb.Auth_Claims_PERMISSION_CAN_EDIT_OWN_FORUM_MESSAGES_WITHOUT_TIME_RESTRICTION)

	dbResponse, err := api.services.DB().FetchTopicMessages(r.Context(), dbTopic.TopicId, params.Limit,
		params.Limit*(params.Page-1), params.SortAsc == 1, userId)

	if err != nil {
		if db.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.TopicId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// TODO Скорее всего, на сайте будет неверно отображаться количество новых сообщений в форуме. Это значение опирается
	//  на Profile->new_forum_answers, а Profile нам из Go совершенно недоступен, ни для чтения, ни для изменения

	_ = api.services.DeleteUserCache(r.Context(), userId)

	attachmentsMap := map[uint64][]helpers.File{}
	for _, attachment := range dbResponse.Attachments {
		attachmentsMap[attachment.MessageId] = append(attachmentsMap[attachment.MessageId], helpers.File{
			Name: attachment.FileName,
			Size: attachment.FileSize,
		})
	}
	for _, message := range dbResponse.Messages {
		messageFiles, _ := api.services.GetFiles(r.Context(), app.ForumMessageFileGroup, message.MessageId)
		attachmentsMap[message.MessageId] = append(attachmentsMap[message.MessageId], messageFiles...)
	}

	attaches := map[uint64][]*pb.Common_Attachment{}
	for messageId, attachment := range attachmentsMap {
		for _, attach := range attachment {
			attaches[messageId] = append(attaches[messageId], &pb.Common_Attachment{
				Name: attach.Name,
				Size: attach.Size,
			})
		}
	}

	var draftAttaches []*pb.Common_Attachment

	if dbResponse.MessageDraft != (db.ForumMessageDraft{}) {
		draftAttachments, _ := helpers.GetForumMessageDraftAttachments(userId, dbTopic.TopicId)
		draftFiles, _ := api.services.GetFiles(r.Context(), app.ForumMessageDraftFileGroup, dbResponse.MessageDraft.DraftId)
		draftAttachments = append(draftAttachments, draftFiles...)

		for _, draftAttachment := range draftAttachments {
			draftAttaches = append(draftAttaches, &pb.Common_Attachment{
				Name: draftAttachment.Name,
				Size: draftAttachment.Size,
			})
		}
	}

	// NOTE Пропущена следующая логика Perl-бэка:
	// 1. workgroup_referee - модераторы (поскольку задается хардкодом в Auth.pm)
	// 2. workgroup_referee могут редактировать цензурированные сообщения
	// 3. хардкод конкретных юзеров как модераторов
	// 4. возможность правки сообщений в FAQ
	// 5. возможность ответа на опрос (пока нет endpoint-а)
	// 6. возможность цензурирования сообщений (пока нет endpoint-а)
	// 7. возможность утверждения находящихся на премодерации пользователей (пока нет endpoint-а)
	// 8. возможность вызова модератора (пока нет endpoint-а)
	// 9. открытие/правка/закрытие темы (пока нет endpoint-ов)
	// 10. если у юзера в настройках отключены смайлы, они просто вырезаются (должны заменяться на алиасы)

	topicMessages := converters.GetTopic(dbResponse, attaches, draftAttaches, params.Page, params.Limit,
		api.services.AppConfig(), user, userCanPerformAdminActions, userCanEditOwnForumMessages)

	return http.StatusOK, topicMessages
}
