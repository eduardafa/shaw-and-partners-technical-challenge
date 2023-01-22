package middlewares

import "net/http"

func SetAcceptMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Accept", "application/vnd.github+json")
		next.ServeHTTP(w, r)
	})
}
