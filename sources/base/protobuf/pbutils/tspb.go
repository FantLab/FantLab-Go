package pbutils

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	tspb "github.com/golang/protobuf/ptypes/timestamp"
)

func TimestampProto(t time.Time) (ts *tspb.Timestamp) {
	ts, _ = ptypes.TimestampProto(t)
	return
}

func Timestamp(ts *tspb.Timestamp) time.Time {
	t, _ := ptypes.Timestamp(ts)
	return t
}
