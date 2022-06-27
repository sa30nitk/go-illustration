package v1

import (
	"fmt"
	"net/http"

	"go-illustration/httpapi/route"
)

func V1() []route.Route {
	return []route.Route{
		{
			http.MethodGet,
			"/gi/internal/v1/bye",
			func(w http.ResponseWriter, r *http.Request) {
				_, err := fmt.Fprintf(w, "bye")
				if err != nil {
					return
				}
			},
		},
	}
}
