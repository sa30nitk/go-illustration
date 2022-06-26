package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(ms ...Middleware) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		for i := len(ms) - 1; i >= 0; i-- {
			if ms[i] == nil {
				continue
			}
			next = ms[i](next)
		}
		return next
	}
}
