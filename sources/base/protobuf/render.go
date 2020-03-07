package protobuf

import (
	"net/http"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

const protoContentType = "application/x-protobuf"

func render(w http.ResponseWriter, r *http.Request, code int, pb proto.Message) {
	acceptProto := r.Header.Get("Accept") == protoContentType

	if acceptProto {
		w.Header().Set("Content-Type", protoContentType)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}

	w.WriteHeader(code)

	var data []byte
	var err error

	if acceptProto {
		data, err = proto.Marshal(pb)
	} else {
		data, err = protojson.Marshal(pb)
	}

	if err == nil {
		_, err = w.Write(data)
	}

	if err != nil {
		panic(err)
	}
}
