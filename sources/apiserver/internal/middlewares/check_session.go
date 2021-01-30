package middlewares

import (
	"fantlab/core/app"
	"fantlab/core/logs"
	"fantlab/pb"
	"net/http"
	"strconv"

	"go.elastic.co/apm"
	"google.golang.org/protobuf/proto"
)

func unpackClaims(services *app.Services, sid string) (*pb.Auth_Claims, error) {
	raw, err := services.CryptoCoder().Decode([]byte(sid))
	if err != nil {
		return nil, err
	}
	claims := new(pb.Auth_Claims)
	err = proto.Unmarshal(raw, claims)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func CheckSession(services *app.Services) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sid := r.Header.Get("X-Session")
			if len(sid) > 0 {
				claims, err := unpackClaims(services, sid)
				if err == nil {
					ctx := app.SetUserAuth(claims, r.Context())
					if t := apm.TransactionFromContext(ctx); t != nil {
						t.Context.SetUserID(strconv.FormatUint(claims.User.UserId, 10))
					}
					r = r.WithContext(ctx)
				} else {
					logs.WithAPM(r.Context()).Error(err.Error())
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}
