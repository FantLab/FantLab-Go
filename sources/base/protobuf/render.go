package protobuf

import (
	"net/http"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

const (
	pbContentType   = "application/x-protobuf"
	jsonContentType = "application/json; charset=utf-8"
)

func marshal(msg proto.Message, usePB bool) ([]byte, error) {
	if usePB {
		return proto.Marshal(msg)
	}
	return protojson.Marshal(msg)
}

func render(w http.ResponseWriter, r *http.Request, code int, msg proto.Message) error {
	usePB := r.Header.Get("Accept") == pbContentType

	if usePB {
		w.Header().Set("Content-Type", pbContentType)
	} else {
		w.Header().Set("Content-Type", jsonContentType)
	}

	w.Header().Set("X-Content-Type-Options", "nosniff")

	w.WriteHeader(code)

	data, err := marshal(msg, usePB)
	if err == nil {
		_, err = w.Write(data)
	}
	return err
}
