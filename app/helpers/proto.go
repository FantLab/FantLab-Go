package helpers

import (
	"time"

	tspb "github.com/golang/protobuf/ptypes/timestamp"
)

func TimestampProto(t time.Time) *tspb.Timestamp {
	return &tspb.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}
