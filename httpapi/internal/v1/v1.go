package v1

import (
	"fmt"
	"net/http"

	"go-illustration/httpapi/server"
)

func V1() []server.Route {
	return []server.Route{
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
