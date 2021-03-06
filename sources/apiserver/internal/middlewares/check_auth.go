package middlewares

import (
	"fantlab/base/protobuf"
	"fantlab/core/app"
	"fantlab/pb"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := app.GetUserAuth(r.Context())

		if auth != nil {
			next.ServeHTTP(w, r)
		} else {
			protobuf.Handle(func(r *http.Request) (int, proto.Message) {
				return http.StatusUnauthorized, &pb.Error_Response{
					Status:  pb.Error_AUTH_REQUIRED,
					Context: "Требуется авторизация",
				}
			}).ServeHTTP(w, r)
		}
	})
}
