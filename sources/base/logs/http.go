package logs

import (
	"fantlab/base/httputils"
	"fantlab/base/logs/logger"
	"fantlab/base/uuid"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var launchId string
var requestId uint64

func init() {
	launchId = uuid.GenerateNow()
}

func nextRequestId() uint64 {
	return atomic.AddUint64(&requestId, 1)
}

func HTTP(fn func(*logger.Request)) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rid := fmt.Sprintf("%s-%d", launchId, nextRequestId())
			w.Header().Set("X-Request-Id", rid)

			ctx, buf := setBuffer(r.Context())

			r = r.WithContext(ctx)

			t := time.Now()
			defer func() {
				d := time.Since(t)

				fn(&logger.Request{
					Id:       rid,
					Host:     r.Host,
					Method:   r.Method,
					URI:      r.RequestURI,
					IP:       r.RemoteAddr,
					Status:   w.(*httputils.LoggingResponseWriter).StatusCode(),
					Entries:  buf.entries,
					Time:     t,
					Duration: d,
				})
			}()

			next.ServeHTTP(w, r)
		})
	}
}
