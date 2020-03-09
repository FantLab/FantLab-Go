package anyserver

import (
	"context"
	"fantlab/base/assert"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func Test_runServers(t *testing.T) {
	var x uint32

	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
			return
		case <-time.After(2 * time.Second):
			atomic.StoreUint32(&x, 1)
			return
		}
	}))

	server := &Server{
		Start: func() error {
			ts.Start()
			return nil
		},
		Stop: func(ctx context.Context) error {
			ts.CloseClientConnections()
			ts.Close()
			return nil
		},
		ShutdownTimeout: 100 * time.Millisecond,
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	go func() {
		time.Sleep(20 * time.Millisecond)
		_, _ = http.Get(ts.URL + "/")
	}()

	runServers([]*Server{server}, ctx.Done(), func(err error) {})

	assert.True(t, atomic.LoadUint32(&x) == 0)
}
