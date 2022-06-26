package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func DummyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("executing dummy middleware")
		log.Debug(r)
		next(w, r)
	}
}
