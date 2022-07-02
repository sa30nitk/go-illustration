package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func RequestLog() func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Debug(r.URL.Path)
			next(w, r)
		}
	}
}
