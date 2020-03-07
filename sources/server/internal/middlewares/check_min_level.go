package middlewares

import (
	"fantlab/base/protobuf"
	"fantlab/pb"
	"fantlab/server/internal/app"
	"fantlab/server/internal/helpers"
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
						Status:  pb.Error_ACTION_PERMITTED,
						Context: fmt.Sprintf("Вы ещё не достигли класса «%s»", helpers.GetUserClassDescription(minUserClass)),
					}
				}).ServeHTTP(w, r)
			}
		})
	}
}
