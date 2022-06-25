package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go-illustration/config"
	external "go-illustration/httpapi/external/v1"
	internal "go-illustration/httpapi/internal/v1"
	"go-illustration/httpapi/server/route"
)

func StartServer(cfg config.Config) error {
	var routes []route.Route
	routes = append(routes, external.V1()...)
	routes = append(routes, internal.V1()...)

	router := httprouter.New()
	for _, r := range routes {
		router.Handler(r.Method, r.Path, r.HandlerFunc)
	}
	router.HandlerFunc(http.MethodGet, "/ping", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "pong") })

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.App.Port), router)
	return err
}
