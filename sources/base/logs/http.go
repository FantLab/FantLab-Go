package logs

import (
	"fantlab/base/logs/logger"
	"fantlab/base/uuid"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"sync/atomic"
	"time"
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

type Config struct {
	NeedsRecover bool
	ToString     logger.ToString
	PanicHandler http.HandlerFunc
}

func HTTP(config Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rid := fmt.Sprintf("%s-%d", launchId, nextRequestId())
			w.Header().Set("X-Request-Id", rid)

			ctx, buf := setBuffer(r.Context())

			writer := &responseWriter{
				ResponseWriter: w,
			}
			request := r.WithContext(ctx)

			t := time.Now()
			defer func() {
				d := time.Since(t)

				var isPanic bool

				if config.NeedsRecover {
					if err := recover(); err != nil {
						buf.Append(logger.Entry{
							Message: string(debug.Stack()),
							Err:     fmt.Errorf("Panic: %v", err),
							Time:    time.Now(),
						})

						isPanic = true
					}
				}

				s := config.ToString(logger.Request{
					Id:       rid,
					Host:     request.Host,
					Method:   request.Method,
					URI:      request.RequestURI,
					IP:       request.RemoteAddr,
					Status:   writer.statusCode,
					Entries:  buf.entries,
					Time:     t,
					Duration: d,
				})

				if "" != s {
					log.Println(s)
				}

				if isPanic {
					config.PanicHandler.ServeHTTP(w, r)
				}
			}()

			next.ServeHTTP(writer, request)
		})
	}
}