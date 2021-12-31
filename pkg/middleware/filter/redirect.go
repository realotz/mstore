package filter

import (
	"net/http"
)

func Redirect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		if w.Header().Get("Location") != "" {
			w.WriteHeader(302)
		}
	})
}
