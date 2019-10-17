package logs

import (
	"fantlab/keys"
	"fantlab/logs/logger"
	"fantlab/pb"
	"fantlab/protobuf"
	"fantlab/uuid"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"sync/atomic"
	"time"

	"github.com/golang/protobuf/proto"
)

// *******************************************************

var launchId string
var requestId uint64

func init() {
	launchId = uuid.GenerateNow()
}

func nextRequestId() uint64 {
	return atomic.AddUint64(&requestId, 1)
}

// *******************************************************

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (r *responseWriter) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

// *******************************************************

func HTTP(strFunc logger.StrFunc) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rid := fmt.Sprintf("%s-%d", launchId, nextRequestId())
			w.Header().Set(keys.HeaderRequestId, rid)

			ctx, buf := setBuffer(r.Context())

			writer := &responseWriter{
				ResponseWriter: w,
			}
			request := r.WithContext(ctx)

			t := time.Now()
			defer func() {
				d := time.Since(t)

				if err := recover(); err != nil {
					buf.append(logger.Entry{
						Date:    time.Now(),
						Message: string(debug.Stack()),
						Err:     fmt.Errorf("Panic: %v", err),
					})

					log.Print(strFunc(logger.HTTPData{
						Id:         rid,
						Request:    request,
						StatusCode: http.StatusInternalServerError,
						Time:       t,
						Duration:   d,
					}, buf.entries))

					protobuf.Handle(func(r *http.Request) (int, proto.Message) {
						return http.StatusInternalServerError, &pb.Error_Response{
							Status: pb.Error_SOMETHING_WENT_WRONG,
						}
					}).ServeHTTP(w, r)
				} else {
					log.Print(strFunc(logger.HTTPData{
						Id:         rid,
						Request:    request,
						StatusCode: writer.statusCode,
						Time:       t,
						Duration:   d,
					}, buf.entries))
				}
			}()

			next.ServeHTTP(writer, request)
		})
	}
}
