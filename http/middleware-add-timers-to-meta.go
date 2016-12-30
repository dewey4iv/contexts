package httpctx

import (
	"net/http"

	"github.com/dewey4iv/contexts"
)

// AddTimersToMeta adds Times present in the context to the context's Meta
func AddTimersToMeta() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)

			Meta(r, "timers", contexts.TimesFromContext(r.Context()))
		})
	}
}
