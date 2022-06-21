package server

import (
	"net/http"

	external "go-illustration/httpapi/external/v1"
	internal "go-illustration/httpapi/internal/v1"
	"go-illustration/httpapi/server/route"
)

func StartServer() {
	var routes []route.Route
	routes = append(routes, external.V1()...)
	routes = append(routes, internal.V1()...)

	for _, r := range routes {
		http.Handle(r.Path, r.HandlerFunc)
	}

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		return
	}
}
