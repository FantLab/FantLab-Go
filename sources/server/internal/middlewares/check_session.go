package middlewares

import (
	"fantlab/pb"
	"fantlab/server/internal/app"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func CheckSession(services *app.Services) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sid := r.Header.Get("X-Session")
			if len(sid) > 0 {
				raw, err := services.CryptoCoder().Decode([]byte(sid))
				if err == nil {
					claims := new(pb.Auth_Claims)
					err = proto.Unmarshal(raw, claims)
					if err == nil {
						ctx := app.SetUserAuth(claims, r.Context())
						next.ServeHTTP(w, r.WithContext(ctx))
						return
					}
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}
