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

func (api *API) DeleteForumMessageDraftFile(r *http.Request) (int, proto.Message) {
	var params struct {
		// id темы
		TopicId uint64 `http:"id,path"`
		// полное имя файла (с расширением)
		FileName string `http:"file_name,form"`
	}

	api.bindParams(&params, r)

	if params.TopicId == 0 {
		return api.badParam("id")
	}

	if !helpers.IsValidFileName(params.FileName) {
		return api.badParam("file_name")
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

	dbMessageDraft, err := api.services.DB().FetchForumMessageDraft(r.Context(), params.TopicId, user.UserId)

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

	attachments, err := app.GetForumMessageDraftAttachments(user.UserId, dbTopic.TopicId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	attachmentExist := false

	for _, attachment := range attachments {
		if attachment.Name == params.FileName {
			attachmentExist = true
			break
		}
	}

	if !attachmentExist {
		files, err := api.services.GetMinioFiles(r.Context(), app.ForumMessageDraftFileGroup, dbMessageDraft.DraftId)

		if err != nil {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status: pb.Error_SOMETHING_WENT_WRONG,
			}
		}

		fileExist := false

		for _, file := range files {
			if file.Name == params.FileName {
				fileExist = true
				break
			}
		}

		if !fileExist {
			return http.StatusForbidden, &pb.Error_Response{
				Status:  pb.Error_ACTION_FORBIDDEN,
				Context: "Не удалось найти аттач с таким именем",
			}
		}
	}

	if attachmentExist {
		app.DeleteForumMessageDraftAttachment(user.UserId, dbTopic.TopicId, params.FileName)
	} else { // Minio file exist
		api.services.DeleteMinioFile(r.Context(), app.ForumMessageDraftFileGroup, dbMessageDraft.DraftId, params.FileName)
	}

	var attaches []*pb.Common_Attachment
	attachments, _ = app.GetForumMessageDraftAttachments(user.UserId, dbTopic.TopicId)
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
