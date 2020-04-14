package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/converters"
	"fmt"
	"google.golang.org/protobuf/proto"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

func (api *API) UploadForumMessageDraftFile(r *http.Request) (int, proto.Message) {
	var params struct {
		// id темы
		TopicId uint64 `http:"id,path"`
		// локальный путь к файлу
		FilePath string `http:"file_path,form"`
	}

	api.bindParams(&params, r)

	if params.TopicId == 0 {
		return api.badParam("id")
	}

	if len(params.FilePath) == 0 {
		return api.badParam("file_path")
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

	uploadedFileCount, err := api.services.DB().FetchForumMessageDraftFileCount(r.Context(), dbMessageDraft.DraftId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if uploadedFileCount >= api.config.MaxAttachCountPerMessage {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: fmt.Sprintf("К сообщению уже приаттачено %d файлов, это максимум", uploadedFileCount),
		}
	}

	now := time.Now()

	fileSize, err := api.services.UploadFile(r.Context(), params.FilePath, now)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	_, fileName := filepath.Split(params.FilePath)

	// Снова читаем из базы количество приаттаченных файлов на случай, если параллельный запрос уже успел залить файл
	uploadedFileCount, err = api.services.DB().FetchForumMessageFileCount(r.Context(), dbMessageDraft.DraftId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if uploadedFileCount >= api.config.MaxAttachCountPerMessage {
		// Удаляем только что залитый файл, потенциальную ошибку игнорируем
		_ = api.services.DeleteFile(r.Context(), fileName, now)

		return http.StatusInternalServerError, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: fmt.Sprintf("К сообщению уже приаттачено %d файлов, это максимум", uploadedFileCount),
		}
	}

	dbMessageDraft, err = api.services.DB().InsertForumMessageDraftFile(r.Context(), dbTopic.TopicId, dbMessageDraft.DraftId, fileName, fileSize, now, user.UserId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	messageDraftResponse := converters.GetForumTopicMessageDraft(dbMessageDraft, api.config)

	return http.StatusOK, messageDraftResponse
}
