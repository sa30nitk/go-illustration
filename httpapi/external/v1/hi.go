package v1

import (
	"fmt"
	"net/http"
)

func hiHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "hi")
		if err != nil {
			return
		}
	}
}
