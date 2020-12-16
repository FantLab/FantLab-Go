package endpoints

import (
	"fantlab/core/app"
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
	"fmt"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) GetForumMessageDraftFileUploadUrl(r *http.Request) (int, proto.Message) {
	var params struct {
		// id темы
		TopicId uint64 `http:"id,path"`
		// полное имя файла (с расширением)
		FileName string `http:"file_name,query"`
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
			Status:  pb.Error_ACTION_PERMITTED,
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

	attachments, err := helpers.GetForumMessageDraftAttachments(user.UserId, dbTopic.TopicId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	for _, attachment := range attachments {
		if attachment.Name == params.FileName {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status:  pb.Error_ACTION_PERMITTED,
				Context: "К сообщению уже приаттачен файл с таким именем",
			}
		}
	}

	files, err := api.services.GetFiles(r.Context(), app.ForumMessageDraftFileGroup, dbMessageDraft.DraftId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	for _, file := range files {
		if file.Name == params.FileName {
			return http.StatusInternalServerError, &pb.Error_Response{
				Status:  pb.Error_ACTION_PERMITTED,
				Context: "К сообщению уже приаттачен файл с таким именем",
			}
		}
	}

	fileCount := uint64(len(attachments) + len(files))

	if fileCount >= api.services.AppConfig().MaxAttachCountPerMessage {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: fmt.Sprintf("К сообщению уже приаттачено %d файлов, это максимум", api.services.AppConfig().MaxAttachCountPerMessage),
		}
	}

	uploadUrl, err := api.services.GetFileUploadUrl(r.Context(), app.ForumMessageDraftFileGroup, dbMessageDraft.DraftId, params.FileName)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	return http.StatusOK, &pb.Common_FileUploadResponse{
		Url: uploadUrl,
	}
}
