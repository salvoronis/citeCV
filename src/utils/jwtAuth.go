package utils

import (
	"net/http"
)

var JwtAuth = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/auth/login", "/auth/register"}
		reqPath := r.URL.Path

		for _, val := range notAuth {
			if val == reqPath {
				next.ServeHTTP(w,r)
				return
			}
		}
	})
}
