package protobuf

import (
	"context"
	"net/http"

	"google.golang.org/protobuf/proto"
)

// jfyi: setup once before server start
var HandleError func(context.Context, error)

type HandlerFunc func(*http.Request) (int, proto.Message)

func Handle(fn HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code, pb := fn(r)
		err := render(w, r, code, pb)
		if HandleError != nil && err != nil {
			HandleError(r.Context(), err)
		}
	}
}
