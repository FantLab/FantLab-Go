package utils

import (
	"time"

	tspb "github.com/golang/protobuf/ptypes/timestamp"
)

func ProtoTS(t time.Time) *tspb.Timestamp {
	return &tspb.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}
