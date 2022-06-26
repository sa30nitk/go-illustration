package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func RequestLog() func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Debug("executing request log")
			log.Debug(r)
			next(w, r)
		}
	}
}