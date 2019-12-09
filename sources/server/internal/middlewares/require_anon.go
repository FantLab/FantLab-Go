package middlewares

import (
	"fantlab/base/protobuf"
	"fantlab/server/internal/keys"
	"fantlab/server/internal/pb"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func RequireAnon(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid := keys.GetUserId(r.Context())

		if uid > 0 {
			protobuf.Handle(func(r *http.Request) (int, proto.Message) {
				return http.StatusMethodNotAllowed, &pb.Error_Response{
					Status: pb.Error_LOG_OUT_FIRST,
				}
			}).ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
