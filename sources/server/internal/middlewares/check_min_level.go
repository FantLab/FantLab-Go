package middlewares

import (
	"fantlab/base/protobuf"
	"fantlab/pb"
	"fantlab/server/internal/app"
	"net/http"

	"github.com/golang/protobuf/proto"
)

type UserClass uint8

const (
	UserClass_Beginner    UserClass = 0
	UserClass_Activist    UserClass = 1
	UserClass_Authority   UserClass = 2
	UserClass_Philosopher UserClass = 3
	UserClass_Master      UserClass = 4
	UserClass_GrandMaster UserClass = 5
	UserClass_PeaceKeeper UserClass = 6
	UserClass_PeaceMaker  UserClass = 7
)

var levelName = map[UserClass]string{
	UserClass_Beginner:    "новичок",
	UserClass_Activist:    "активист",
	UserClass_Authority:   "авторитет",
	UserClass_Philosopher: "философ",
	UserClass_Master:      "магистр",
	UserClass_GrandMaster: "гранд-мастер",
	UserClass_PeaceKeeper: "миродержец",
	UserClass_PeaceMaker:  "миротворец",
}

func CheckMinLevel(services *app.Services, minUserClass UserClass) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			uid := app.GetUserId(r.Context())

			if uid > 0 {
				rawUserClass, err := services.DB().FetchUserClass(r.Context(), uid)

				if err != nil {
					protobuf.Handle(func(r *http.Request) (int, proto.Message) {
						return http.StatusInternalServerError, &pb.Error_Response{
							Status: pb.Error_SOMETHING_WENT_WRONG,
						}
					}).ServeHTTP(w, r)

					return
				}

				if UserClass(rawUserClass) >= minUserClass {
					next.ServeHTTP(w, r)

					return
				}
			}

			protobuf.Handle(func(r *http.Request) (int, proto.Message) {
				return http.StatusForbidden, &pb.Error_Response{
					Status:  pb.Error_ACTION_PERMITTED,
					Context: "Вы ещё не достигли класса «" + levelName[minUserClass] + "»",
				}
			}).ServeHTTP(w, r)
		})
	}
}
