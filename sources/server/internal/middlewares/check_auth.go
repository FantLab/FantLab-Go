package middlewares

import (
	"fantlab/base/protobuf"
	"fantlab/pb"
	"fantlab/server/internal/app"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := app.GetUserAuth(r.Context())

		if auth != nil {
			next.ServeHTTP(w, r)
		} else {
			protobuf.Handle(func(r *http.Request) (int, proto.Message) {
				return http.StatusUnauthorized, &pb.Error_Response{
					Status: pb.Error_AUTH_REQUIRED,
				}
			}).ServeHTTP(w, r)
		}
	})
}
