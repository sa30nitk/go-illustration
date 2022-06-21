package v1

import (
	"fmt"
	"net/http"

	"go-illustration/httpapi/server/route"
)

func V1() []route.Route {
	return []route.Route{
		{
			http.MethodGet,
			"/gi/v1/hi",
			func(w http.ResponseWriter, r *http.Request) {
				_, err := fmt.Fprintf(w, "hi")
				if err != nil {
					return
				}
			},
		},
	}
}
