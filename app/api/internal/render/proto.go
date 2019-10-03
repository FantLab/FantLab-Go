package render

import (
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

const protoContentType = "application/x-protobuf"

func writePB(w http.ResponseWriter, pb proto.Message) error {
	data, err := proto.Marshal(pb)

	if err == nil {
		_, err = w.Write(data)
	}

	return err
}

func writeJSON(w http.ResponseWriter, pb proto.Message) error {
	m := jsonpb.Marshaler{
		OrigName: true,
	}

	return m.Marshal(w, pb)
}

func Proto(w http.ResponseWriter, r *http.Request, code int, pb proto.Message) {
	var err error

	if r.Header.Get("Accept") == protoContentType {
		w.Header().Set("Content-Type", protoContentType)
		err = writePB(w, pb)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		err = writeJSON(w, pb)
	}

	if err == nil {
		w.WriteHeader(code)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
