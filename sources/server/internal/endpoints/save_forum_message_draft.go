package endpoints

import (
	"fantlab/base/dbtools"
	"fantlab/pb"
	"fantlab/server/internal/converters"
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

	// Здесь мы отходим от логики Perl-бэка. В нем нет этой проверки, поэтому можно создать черновик сообщения в
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
	if messageLength > api.config.MaxForumMessageLength {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: fmt.Sprintf("Текст сообщения слишком длинный (больше %d символов)", api.config.MaxForumMessageLength),
		}
	}

	dbMessageDraft, err := api.services.DB().InsertForumMessageDraft(r.Context(), params.Message, dbTopic.TopicId, user.UserId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// TODO: загрузка аттачей

	messageDraftResponse := converters.GetForumTopicMessageDraft(dbMessageDraft, api.config)

	return http.StatusOK, messageDraftResponse
}
