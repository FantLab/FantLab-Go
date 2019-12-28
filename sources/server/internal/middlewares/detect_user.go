package middlewares

import (
	"fantlab/server/internal/app"
	"net/http"
)

func DetectUser(services *app.Services) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sid := r.Header.Get(app.SessionHeader)

			if len(sid) > 0 {
				uid := services.GetUserIdBySessionId(r.Context(), sid)

				if uid > 0 {
					ctx := app.SetUserId(uid, r.Context())
					next.ServeHTTP(w, r.WithContext(ctx))

					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
