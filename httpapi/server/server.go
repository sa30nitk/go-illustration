package server

import (
	"fmt"
	"net/http"

	"go-illustration/config"
	external "go-illustration/httpapi/external/v1"
	internal "go-illustration/httpapi/internal/v1"
	"go-illustration/httpapi/server/route"
)

func StartServer(cfg config.Config) error {
	var routes []route.Route
	routes = append(routes, external.V1()...)
	routes = append(routes, internal.V1()...)

	for _, r := range routes {
		http.Handle(r.Path, r.HandlerFunc)
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "pong") })

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.App.Port), nil)
	return err
}
