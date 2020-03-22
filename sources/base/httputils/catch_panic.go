package httputils

import (
	"net/http"
)

func CatchPanic(panicHandler func(http.ResponseWriter, *http.Request, interface{})) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if x := recover(); x != nil {
					panicHandler(w, r, x)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
