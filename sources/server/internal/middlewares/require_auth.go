package middlewares

import (
	"fantlab/base/protobuf"
	"fantlab/pb"
	"fantlab/server/internal/app"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid := app.GetUserId(r.Context())

		if uid > 0 {
			next.ServeHTTP(w, r)
		} else {
			protobuf.Handle(func(r *http.Request) (int, proto.Message) {
				return http.StatusUnauthorized, &pb.Error_Response{
					Status: pb.Error_INVALID_SESSION,
				}
			}).ServeHTTP(w, r)
		}
	})
}
