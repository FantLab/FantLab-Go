package protobuf

import (
	"net/http"

	"google.golang.org/protobuf/proto"
)

type HandlerFunc func(*http.Request) (int, proto.Message)

func Handle(fn HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code, pb := fn(r)

		render(w, r, code, pb)
	}
}
