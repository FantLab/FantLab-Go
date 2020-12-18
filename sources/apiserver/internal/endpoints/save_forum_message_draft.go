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

	// NOTE Здесь мы отходим от логики Perl-бэка. В нем нет этой проверки, поэтому можно создать черновик сообщения в
	// закрытой теме, который принципиально нельзя подтвердить (facepalm)
	if dbTopic.IsClosed == 1 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Тема закрыта",
		}
	}

	user := api.getUser(r)

	messageLength := uint64(len(params.Message))

	if messageLength == 0 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Текст сообщения пустой",
		}
	}

	// В Perl-бэке стоит ограничение в 100_000 символов. Это приводит к странному эффекту: можно сохранить в черновик
	// сообщение длиннее 20_000 символов, но подтвердить его нельзя, не урезав
	if messageLength > api.services.AppConfig().MaxForumMessageLength {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: fmt.Sprintf("Текст сообщения слишком длинный (больше %d символов)", api.services.AppConfig().MaxForumMessageLength),
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

	messageDraftResponse := converters.GetForumTopicMessageDraft(dbMessageDraft, attaches, api.services.AppConfig())

	return http.StatusOK, messageDraftResponse
}
