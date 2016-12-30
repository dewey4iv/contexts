package httpctx

import "net/http"

// RenderContexts is an http middleware for json.
// It looks at all of the possible contexts and renders them to the ResponseWriter
func RenderContexts() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}
