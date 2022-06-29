package middleware

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type Statsd interface {
	PrecisionTiming(stat string, delta time.Duration) error
}

func StatsdTiming(reporter Statsd) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()
			next(w, r)
			delta := time.Now().Sub(now)
			err := reporter.PrecisionTiming(fmt.Sprintf("timing_%s_%s", r.Method, r.URL.Path), delta)
			if err != nil {
				log.Errorf("Failed to publish statsd metric")
				return
			}
		}
	}
}
