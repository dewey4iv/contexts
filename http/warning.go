package httpctx

import (
	"net/http"

	"github.com/dewey4iv/contexts"
)

// Warnings add warnings to the provided http context
func Warnings(r *http.Request, warnings ...string) {
	r.WithContext(contexts.WarningsWithContext(r.Context(), warnings...))
}
