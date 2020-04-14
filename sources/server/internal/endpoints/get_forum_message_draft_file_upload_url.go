package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/app"
	"fmt"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
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

	fileCount := uint64(len(files))

	if fileCount == api.config.MaxAttachCountPerMessage {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: fmt.Sprintf("К сообщению уже приаттачено %d файлов, это максимум", fileCount),
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
