package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/newrelic/go-agent/v3/integrations/nrhttprouter"
	log "github.com/sirupsen/logrus"
	"go-illustration/config"
	external "go-illustration/httpapi/external/v1"
	internal "go-illustration/httpapi/internal/v1"
	"go-illustration/httpapi/route"
	"go-illustration/newrelic"
)

type Reporter interface {
	Incr(stat string, count int64) error
	PrecisionTiming(stat string, delta time.Duration) error
}

type PlaceHolder interface {
	Placeholder(ctx context.Context) *http.Response
}

type Translator interface {
	Localize(msg string, langs ...string) string
}

type Dependencies struct {
	StatsD      Reporter
	PlaceHolder PlaceHolder
	Translator  Translator
}

func StartServer(cfg config.Config, dependencies Dependencies) {
	var routes []route.Route
	routes = append(routes, external.V1(external.Deps{
		PlaceHolder: dependencies.PlaceHolder,
		Reporter:    dependencies.StatsD,
		Translator:  dependencies.Translator,
	})...)
	routes = append(routes, internal.V1()...)

	app, err := newrelic.NRApp(cfg.NR)
	if err != nil {
		return
	}

	router := nrhttprouter.New(app)
	for _, r := range routes {
		router.Handler(r.Method, r.Path, r.HandlerFunc)
	}
	router.HandlerFunc(http.MethodGet, "/ping", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "pong") })

	var srv http.Server
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT)
		<-sigint

		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Debugf("HTTP server Shutdown: %v", err)
		}

		app.Shutdown(time.Second * 10)
		close(idleConnsClosed)
	}()

	addr := fmt.Sprintf(":%d", cfg.App.Port)
	srv = http.Server{
		Addr:    addr,
		Handler: router,
	}

	err = srv.ListenAndServe()
	<-idleConnsClosed

	log.Info("server is closing")
}
