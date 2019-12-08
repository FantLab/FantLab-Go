package protobuf

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

func render(w http.ResponseWriter, r *http.Request, code int, pb proto.Message) {
	acceptProto := r.Header.Get("Accept") == protoContentType

	if acceptProto {
		w.Header().Set("Content-Type", protoContentType)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}

	w.WriteHeader(code)

	var err error

	if acceptProto {
		err = writePB(w, pb)
	} else {
		err = writeJSON(w, pb)
	}

	if err != nil {
		panic(err)
	}
}
