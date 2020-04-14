package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/converters"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
)

func (api *API) DeleteForumMessageDraftFile(r *http.Request) (int, proto.Message) {
	var params struct {
		// id темы
		TopicId uint64 `http:"id,path"`
		// id файла
		FileId uint64 `http:"file_id,form"`
	}

	api.bindParams(&params, r)

	if params.TopicId == 0 {
		return api.badParam("id")
	}

	if params.FileId == 0 {
		return api.badParam("file_id")
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

	dbFile, err := api.services.DB().FetchForumMessageDraftFile(r.Context(), dbMessageDraft.DraftId, params.FileId)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.FileId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	dbMessageDraft, err = api.services.DB().DeleteForumMessageDraftFile(r.Context(), dbMessageDraft.DraftId, dbFile.FileId, dbTopic.TopicId, user.UserId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// Удаляем файл, ошибку игнорим
	_ = api.services.DeleteFile(r.Context(), dbFile.FileName, dbFile.DateOfAdd)

	messageDraftResponse := converters.GetForumTopicMessageDraft(dbMessageDraft, api.config)

	return http.StatusOK, messageDraftResponse
}
