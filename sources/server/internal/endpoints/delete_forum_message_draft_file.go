package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/app"
	"fantlab/server/internal/converters"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
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

	if len(params.FileName) == 0 {
		return api.badParam("file_name")
	}

	availableForums := api.getAvailableForums(r)

	dbTopic, err := api.services.DB().FetchForumTopic(r.Context(), availableForums, params.TopicId)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.TopicId, 10),
			}
		}

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
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.TopicId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	files, err := api.services.GetFiles(r.Context(), app.ForumMessageDraftFileGroup, dbMessageDraft.DraftId)

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
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Не удалось найти аттач с таким именем",
		}
	}

	// Удаляем файл, ошибку игнорим
	_ = api.services.DeleteFile(r.Context(), app.ForumMessageDraftFileGroup, dbMessageDraft.DraftId, params.FileName)

	messageDraftResponse := converters.GetForumTopicMessageDraft(dbMessageDraft, api.config)

	return http.StatusOK, messageDraftResponse
}
