package utils

import (
	"fantlab/protobuf/generated/fantlab/apimodels"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func ShowError(ctx *gin.Context, code int, message string) {
	ShowProto(ctx, code, &apimodels.ResponseError{
		Code:    int32(code),
		Message: message,
	})
}

func ShowProto(ctx *gin.Context, code int, pb proto.Message) {
	accept := ctx.GetHeader("Accept")

	if accept == "application/json" {
		ctx.Header("Content-Type", "application/json; charset=utf-8")
		marshaler := jsonpb.Marshaler{
			OrigName: true,
		}
		marshaler.Marshal(ctx.Writer, pb)
	} else {
		ctx.ProtoBuf(code, pb)
	}
}
