package endpoints

import (
	"fantlab/core/converters"
	"fantlab/pb"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func (api *API) ShowForums(r *http.Request) (int, proto.Message) {
	userId := api.getUserId(r)

	availableForums := api.getAvailableForums(r)

	dbForumsResponse, err := api.services.DB().FetchForums(r.Context(), userId, availableForums)

	if err != nil {
		return http.StatusInternalServerError, &pb.Error_Response{
			Status: pb.Error_SOMETHING_WENT_WRONG,
		}
	}

	// NOTE В сравнении с Perl-бэком, пропущен следующий код:
	// 1. Все, что касается блоков ниже списка форумов (Online посетители, Статистика форумов и Дни рождения)
	// 2. 2 UPDATE-запроса к базе: на обновление рекорда online-посещаемости (для этого необходимо было бы адаптировать
	//   MemcachedFunctions::MC_GetLastUsersActions, оперирующий сложными хешами из Memcached) и на сброс счетчика
	//   topics_to_moderate_count у пользователя (поскольку я не понимаю стоящей за этим логики)

	forumBlocks := converters.GetForumBlocks(dbForumsResponse, userId, api.services.AppConfig())
	return http.StatusOK, forumBlocks
}
