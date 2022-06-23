package server

import (
	"fmt"
	"net/http"

	"go-illustration/config"
	external "go-illustration/httpapi/external/v1"
	internal "go-illustration/httpapi/internal/v1"
	"go-illustration/httpapi/server/route"
)

func StartServer(cfg config.Config) {
	var routes []route.Route
	routes = append(routes, external.V1()...)
	routes = append(routes, internal.V1()...)

	for _, r := range routes {
		http.Handle(r.Path, r.HandlerFunc)
	}

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Printf("Failed to launch server with error: %s\n", err)
		panic("Failed to start")
	}
}
