package api

import (
	"fantlab/api/internal/render"
	"net/http"

	"github.com/golang/protobuf/proto"
)

type apiFunc func(*http.Request) (int, proto.Message)

func httpHandler(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code, pb := fn(r)

		render.Proto(w, r, code, pb)
	}
}
