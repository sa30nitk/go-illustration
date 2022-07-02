package v1

import (
	"net/http"

	"go-illustration/httpapi/middleware"
	"go-illustration/httpapi/route"
)

type Deps struct {
	PlaceHolder PlaceHolder
	Reporter    middleware.Statsd
}

func V1(deps Deps) []route.Route {
	return []route.Route{
		{
			http.MethodGet,
			"/gi/v1/hi",
			middleware.Chain(
				middleware.RequestLog(),
				middleware.StatsdTiming(deps.Reporter),
			)(hiHandler(deps.PlaceHolder)),
		},
	}
}
