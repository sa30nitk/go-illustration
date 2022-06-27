package server

import (
	"fmt"
	"net/http"

	"github.com/newrelic/go-agent/v3/integrations/nrhttprouter"
	"go-illustration/config"
	external "go-illustration/httpapi/external/v1"
	internal "go-illustration/httpapi/internal/v1"
	"go-illustration/httpapi/route"
	"go-illustration/newrelic"
)

func StartServer(cfg config.Config) error {
	var routes []route.Route
	routes = append(routes, external.V1()...)
	routes = append(routes, internal.V1()...)

	app, err := newrelic.NRApp(cfg.NR)
	if err != nil {
		return err
	}
	router := nrhttprouter.New(app)
	for _, r := range routes {
		router.Handler(r.Method, r.Path, r.HandlerFunc)
	}
	router.HandlerFunc(http.MethodGet, "/ping", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "pong") })

	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.App.Port), router)
	return err
}
