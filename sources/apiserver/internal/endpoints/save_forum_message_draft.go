package endpoints

import (
	"fantlab/core/app"
	"fantlab/core/converters"
	"fantlab/pb"
	"fmt"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) SaveForumMessageDraft(r *http.Request) (int, proto.Message) {
	var params struct {
		// id темы
		TopicId uint64 `http:"id,path"`
		// текст сообщения
		Message string `http:"message,form"`
	}

	api.bindParams(&params, r)

	if params.TopicId == 0 {
		return api.badParam("id")
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

	if dbTopic.IsClosed == 1 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Тема закрыта",
		}
	}

	user := api.getUser(r)

	messageLength := uint64(len(params.Message))

	if messageLength == 0 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Текст сообщения пустой",
		}
	}

	if messageLength > api.services.AppConfig().MaxMessageLength {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: fmt.Sprintf("Текст сообщения слишком длинный (больше %d символов)", api.services.AppConfig().MaxMessageLength),
		}
	}

	dbMessageDraft, err := api.services.DB().InsertForumMessageDraft(r.Context(), params.Message, dbTopic.TopicId, user.UserId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	var attaches []*pb.Common_Attachment
	attachments, _ := app.GetForumMessageDraftAttachments(user.UserId, dbTopic.TopicId)
	for _, attachment := range attachments {
		attaches = append(attaches, &pb.Common_Attachment{
			Url:  api.services.GetFSForumMessageDraftAttachmentUrl(user.UserId, dbTopic.TopicId, attachment.Name),
			Size: attachment.Size,
		})
	}
	files, _ := api.services.GetMinioFiles(r.Context(), app.ForumMessageDraftFileGroup, dbMessageDraft.DraftId)
	for _, file := range files {
		attaches = append(attaches, &pb.Common_Attachment{
			Url:  api.services.GetMinioForumMessageDraftAttachmentUrl(dbMessageDraft.DraftId, file.Name),
			Size: file.Size,
		})
	}

	// NOTE Пропущено достаточно много логики в сравнении с самими сообщениями. Например, нет проверок, что
	// пользователь находится в readonly в данном форуме. То есть он сможет создать черновик, но не сможет
	// подтвердить его

	messageDraftResponse := converters.GetForumTopicMessageDraft(dbMessageDraft, attaches, api.services.AppConfig())

	return http.StatusOK, messageDraftResponse
}
