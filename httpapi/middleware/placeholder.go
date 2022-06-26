package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Placeholder(dep string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Debug("placeholde dependency: ", dep)
			log.Debug("placeholder middleware")
			log.Debug(r)
			next(w, r)
		}
	}
}
