package utils

import (
	"fantlab/protobuf/generated/fantlab/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func ShowError(ctx *gin.Context, code int, message string) {
	ShowProto(ctx, code, &pb.ErrorResponse{
		ErrorCode: int32(code),
		Message:   message,
	})
}

func ShowProto(ctx *gin.Context, code int, pb proto.Message) {
	if ctx.GetHeader("Accept") == "application/x-protobuf" {
		ctx.ProtoBuf(code, pb)
		return
	}

	ctx.Header("Content-Type", "application/json; charset=utf-8")

	marshaller := jsonpb.Marshaler{
		OrigName: true,
	}

	if gin.IsDebugging() {
		marshaller.Indent = "  "
	}

	if err := marshaller.Marshal(ctx.Writer, pb); err != nil {
		ShowError(ctx, http.StatusInternalServerError, err.Error())
	}
}
