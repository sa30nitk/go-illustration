package v1

import (
	"net/http"

	"go-illustration/httpapi/middleware"
	"go-illustration/httpapi/route"
)

func V1(reporter middleware.Statsd) []route.Route {
	return []route.Route{
		{
			http.MethodGet,
			"/gi/v1/hi",
			middleware.Chain(
				middleware.RequestLog(),
				middleware.StatsdTiming(reporter),
			)(hiHandler()),
		},
	}
}
