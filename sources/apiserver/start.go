package apiserver

import (
	"context"
	"expvar"
	"fantlab/apiserver/internal/routes"
	"fantlab/base/protobuf"
	"fantlab/core/app"
	"fantlab/core/logs"
	"fantlab/docs"
	"net/http"
	"net/http/pprof"
	"os"
	"time"

	"github.com/FantLab/go-kit/anyserver"
	"github.com/FantLab/go-kit/http/mux"

	_ "go.elastic.co/apm/module/apmsql/mysql"
)

func GenerateDocs() {
	_ = docs.Generate(os.Stdout, routes.Tree(nil, nil), "/"+routes.BasePath)
}

func Start() {
	protobuf.HandleError = func(ctx context.Context, err error) {
		logs.WithAPM(ctx).Error(err.Error())
	}

	apiServer := makeAPIServer()

	var monitoringServer *anyserver.Server

	if apiServer.SetupError == nil {
		monitoringServer = makeMonitoringServer()
	}

	anyserver.RunWithGracefulShutdown(func(err error) {
		logs.Logger().Error(err.Error())
	}, apiServer, monitoringServer)

	time.Sleep(1 * time.Second)
}

func makeAPIServer() (server *anyserver.Server) {
	server = new(anyserver.Server)

	var services *app.Services

	services, server.SetupError, server.DisposeBag = app.MakeServices()

	if server.SetupError != nil {
		return
	}

	httpHandler, markAsUnavailable := routes.MakeHandler(services)

	httpServer := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: httpHandler,
	}

	server.Start = func() error {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	}
	server.Stop = func(ctx context.Context) error {
		markAsUnavailable()
		httpServer.SetKeepAlivesEnabled(false)
		return httpServer.Shutdown(ctx)
	}
	server.ShutdownTimeout = 5 * time.Second

	return
}

func makeMonitoringServer() (server *anyserver.Server) {
	port := os.Getenv("MONITORING_PORT")
	if port == "" {
		return nil
	}

	server = new(anyserver.Server)

	var httpHandler http.Handler

	{
		rootGroup := new(mux.Group)
		{
			rootGroup.Endpoint(http.MethodGet, "/pprof/:index", http.HandlerFunc(pprof.Index))
			rootGroup.Endpoint(http.MethodGet, "/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
			rootGroup.Endpoint(http.MethodGet, "/pprof/profile", http.HandlerFunc(pprof.Profile))
			rootGroup.Endpoint(http.MethodGet, "/pprof/symbol", http.HandlerFunc(pprof.Symbol))
			rootGroup.Endpoint(http.MethodGet, "/pprof/trace", http.HandlerFunc(pprof.Trace))
		}
		{
			rootGroup.Endpoint(http.MethodGet, "/expvar", expvar.Handler())
		}

		routerConfig := &mux.Config{
			RootGroup: rootGroup,
			NotFoundHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			}),
			CommonPrefix: "debug",
		}

		httpHandler, _ = mux.NewRouter(routerConfig)
	}

	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: httpHandler,
	}

	server.Start = func() error {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	}
	server.Stop = func(ctx context.Context) error {
		httpServer.SetKeepAlivesEnabled(false)

		return httpServer.Shutdown(ctx)
	}
	server.ShutdownTimeout = 5 * time.Second

	return
}
