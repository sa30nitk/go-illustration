package v1

import (
	"net/http"

	"go-illustration/httpapi/middleware"
	"go-illustration/httpapi/server/route"
)

func V1() []route.Route {
	return []route.Route{
		{
			http.MethodGet,
			"/gi/v1/hi",
			middleware.DummyMiddleware(middleware.RequestLog(hiHandler())),
			//middlewareChain(middleware.RequestLog(), middleware.DummyMiddleware()),
		},
	}
}
