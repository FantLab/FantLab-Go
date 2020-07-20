package endpoints

import (
	"fantlab/core/db"
	"fantlab/pb"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (api *API) AddResponse(r *http.Request) (int, proto.Message) {
	var params struct {
		// id произведения
		WorkId uint64 `http:"id,path"`
		// текст отзыва
		Response string `http:"response,form"`
	}

	api.bindParams(&params, r)

	if params.WorkId == 0 {
		return api.badParam("id")
	}

	// TODO: В Perl-бэке баг. Удаление черновика отзыва происходит до любых проверок

	// TODO: В Perl-бэке этой проверки нет. Так что можно добавить отзыв несуществующему произведению
	dbWork, err := api.services.DB().FetchWork(r.Context(), params.WorkId)

	if err != nil {
		if db.IsNotFoundError(err) {
			return http.StatusNotFound, &pb.Error_Response{
				Status:  pb.Error_NOT_FOUND,
				Context: strconv.FormatUint(params.WorkId, 10),
			}
		}

		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// TODO: В Perl-бэке два бага. Во-первых, не обрезаются пробельные символы по краям текста (в отличие от
	//  редактирования). Во-вторых, длина текста проверяется до вырезания смайлов. Это значит, что, к примеру, отзыв,
	//  состоящий из одних смайлов, будет успешно добавлен как пустой
	formattedResponse := api.services.AppConfig().Smiles.RemoveFromString(strings.TrimSpace(params.Response))

	if uint64(len(formattedResponse)) < api.services.AppConfig().MinResponseLength {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: fmt.Sprintf("Текст сообщения слишком короткий (меньше %d символов после удаления смайлов)", api.services.AppConfig().MinResponseLength),
		}
	}

	userId := api.getUserId(r)

	userResponseCountForWork, err := api.services.DB().FetchUserResponseCountForWork(r.Context(), userId, dbWork.WorkId)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if userResponseCountForWork >= api.services.AppConfig().MaxUserResponseCountPerWork {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: fmt.Sprintf("На одно произведение можно написать не больше %d отзывов", api.services.AppConfig().MaxUserResponseCountPerWork),
		}
	}

	suchUserResponseCountForWork, err := api.services.DB().FetchSuchUserResponseCountForWork(r.Context(), userId, dbWork.WorkId, formattedResponse)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	if suchUserResponseCountForWork > 0 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_PERMITTED,
			Context: "У вас уже есть такой отзыв на данное произведение",
		}
	}

	var workAuthorIds []uint64
	if dbWork.AutorId != 0 {
		workAuthorIds = append(workAuthorIds, dbWork.AutorId)
	}
	if dbWork.Autor2Id != 0 {
		workAuthorIds = append(workAuthorIds, dbWork.Autor2Id)
	}
	if dbWork.Autor3Id != 0 {
		workAuthorIds = append(workAuthorIds, dbWork.Autor3Id)
	}
	if dbWork.Autor4Id != 0 {
		workAuthorIds = append(workAuthorIds, dbWork.Autor4Id)
	}
	if dbWork.Autor5Id != 0 {
		workAuthorIds = append(workAuthorIds, dbWork.Autor5Id)
	}

	err = api.services.DB().InsertResponse(r.Context(), userId, dbWork.WorkId, workAuthorIds, formattedResponse)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	_ = api.services.SetUserResponseCache(r.Context(), userId, dbWork.WorkId)
	_ = api.services.DeleteUserCache(r.Context(), userId)
	_ = api.services.DeleteWorkStatCache(r.Context(), dbWork.WorkId)
	_ = api.services.DeleteHomepageResponsesCache(r.Context())

	// NOTE Код, эквивалентный https://github.com/parserpro/fantlab/blob/8b1e361f4d0379369aaf8aae6063c558b768ec01/lib/BD/Response.pm#L447-L448
	// и https://github.com/parserpro/fantlab/blob/8b1e361f4d0379369aaf8aae6063c558b768ec01/pm/Functions.pm#L3181-L3201,
	// не перенесен, поскольку на проде нет реальных следов кеширования в виде непустых каталогов /cache/autors,
	// /cache/works, /cache/xml, /cache/json

	// NOTE В Perl-бэке есть код, добавляющий в таблицу tasks задачу на генерацию RSS для нового отзыва. Перенести его
	// в Go проблематично, поскольку это требует имплементации аналога Storable::nfreeze из Perl

	return http.StatusOK, &pb.Common_SuccessResponse{}
}
