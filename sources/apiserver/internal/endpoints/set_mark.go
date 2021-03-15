package endpoints

import (
	"fantlab/core/db"
	"fantlab/core/helpers"
	"fantlab/pb"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
)

func (api *API) SetMark(r *http.Request) (int, proto.Message) {
	var params struct {
		// id произведения
		WorkId uint64 `http:"id,path"`
		// оценка (1-10, 0 - удалить)
		Mark uint8 `http:"mark,form"`
	}

	api.bindParams(&params, r)

	if params.WorkId == 0 {
		return api.badParam("id")
	}

	if params.Mark > 10 {
		return api.badParam("mark")
	}

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

	if dbWork.Published == 0 {
		return http.StatusForbidden, &pb.Error_Response{
			Status:  pb.Error_ACTION_FORBIDDEN,
			Context: "Нельзя поставить оценку неопубликованному произведению",
		}
	}

	userId := api.getUserId(r)

	isFlContestWork := api.services.AppConfig().FlContestInProgress && dbWork.AutorId == api.services.AppConfig().FlContestAuthorId
	correlationUserMarkCountThreshold := api.services.AppConfig().CorrelationUserMarkCountThreshold

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

	workStats, err := api.services.DB().UpsertMark(r.Context(), userId, dbWork.WorkId, workAuthorIds, isFlContestWork,
		params.Mark, correlationUserMarkCountThreshold)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	_ = api.services.DeleteWorkStatCache(r.Context(), dbWork.WorkId)
	_ = api.services.DeleteUserCache(r.Context(), userId)
	_ = api.services.DeleteUserMarksCache(r.Context(), userId)
	helpers.DeleteWorkRatingImageCache(dbWork.WorkId)

	// NOTE Код, эквивалентный https://github.com/parserpro/fantlab/blob/8b1e361f4d0379369aaf8aae6063c558b768ec01/pm/Functions.pm#L885-L890
	// https://github.com/parserpro/fantlab/blob/8b1e361f4d0379369aaf8aae6063c558b768ec01/pm/Functions.pm#L3181-L3201
	// и https://github.com/parserpro/fantlab/blob/8b1e361f4d0379369aaf8aae6063c558b768ec01/pm/Functions.pm#L3147-L3151,
	// не перенесен, поскольку на проде нет реальных следов кеширования в виде непустых каталогов /cache/autors,
	// /cache/works, /cache/xml, /cache/json, /cache/recoms

	var averageMark float64
	var markCount int64

	if isFlContestWork {
		// NOTE Непонятно, в чем суть, поскольку, во-первых, flcontest_is_going в конфиге толком не меняется, во-вторых,
		// на странице автора эти рейтинг/количество оценок все равно отображаются. Похоже на баг в Perl-бэке
		averageMark = -1
		markCount = -1
	} else {
		averageMark = workStats.AverageMarkByWeight
		markCount = int64(workStats.MarkCount)
	}

	return http.StatusOK, &pb.Mark_Response{
		AverageMark: averageMark,
		MarkCount:   markCount,
	}
}
