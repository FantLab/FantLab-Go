package pbutils

import (
	"time"

	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

func TimestampProto(t time.Time) *tspb.Timestamp {
	return &tspb.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}

func Timestamp(ts *tspb.Timestamp) time.Time {
	if ts == nil {
		return time.Unix(0, 0).UTC()
	}
	return time.Unix(ts.Seconds, int64(ts.Nanos)).UTC()
}
