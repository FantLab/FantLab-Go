package endpoints

import (
	"fantlab/core/app"
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (api *API) CancelForumMessageDraft(r *http.Request) (int, proto.Message) {
	var params struct {
		// id темы
		TopicId uint64 `http:"id,path"`
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
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "Тема закрыта",
		}
	}

	userId := api.getUserId(r)

	dbMessageDraft, err := api.services.DB().FetchForumMessageDraft(r.Context(), dbTopic.TopicId, userId)

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

	err = api.services.DB().DeleteForumMessageDraft(r.Context(), dbTopic.TopicId, userId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	helpers.DeleteForumMessageDraftAttachments(userId, dbTopic.TopicId)
	api.services.DeleteFiles(r.Context(), app.ForumMessageDraftFileGroup, dbMessageDraft.DraftId)

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
