package middlewares

import (
	"fantlab/base/protobuf"
	"fantlab/core/app"
	"fantlab/core/helpers"
	"fantlab/pb"
	"fmt"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func CheckMinLevel(minUserClass pb.Common_UserClass) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := app.GetUserAuth(r.Context())

			if auth.User.Class >= minUserClass {
				next.ServeHTTP(w, r)
			} else {
				protobuf.Handle(func(r *http.Request) (int, proto.Message) {
					return http.StatusForbidden, &pb.Error_Response{
						Status:  pb.Error_ACTION_FORBIDDEN,
						Context: fmt.Sprintf("Вы ещё не достигли класса «%s»", helpers.UserClassDescriptionMap[minUserClass]),
					}
				}).ServeHTTP(w, r)
			}
		})
	}
}
