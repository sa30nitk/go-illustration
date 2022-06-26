package v1

import (
	"net/http"

	"go-illustration/httpapi/middleware"
	"go-illustration/httpapi/server"
)

func V1() []server.Route {
	return []server.Route{
		{
			http.MethodGet,
			"/gi/v1/hi",
			middleware.Chain(
				middleware.Placeholder("hi"),
				middleware.RequestLog(),
			)(hiHandler()),
		},
	}
}
