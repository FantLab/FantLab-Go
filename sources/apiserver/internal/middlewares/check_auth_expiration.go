package middlewares

import (
	"fantlab/base/protobuf"
	"fantlab/base/protobuf/pbutils"
	"fantlab/core/app"
	"fantlab/core/config"
	"fantlab/pb"
	"net/http"
	"time"

	"google.golang.org/protobuf/proto"
)

func CheckAuthExpiration(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := app.GetUserAuth(r.Context())

		dt := time.Until(pbutils.Timestamp(auth.Issued).Add(config.AuthTokenTimeout))

		if dt > 0 {
			w.Header().Set("X-Auth-Expired", dt.String())

			next.ServeHTTP(w, r)
		} else {
			protobuf.Handle(func(r *http.Request) (int, proto.Message) {
				return http.StatusUnauthorized, &pb.Error_Response{
					Status: pb.Error_AUTH_TOKEN_EXPIRED,
				}
			}).ServeHTTP(w, r)
		}
	})
}
