package v1

import (
	"net/http"

	"go-illustration/httpapi/middleware"
	"go-illustration/httpapi/route"
)

func V1() []route.Route {
	return []route.Route{
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
