package httpctx

import (
	"net/http"

	"github.com/dewey4iv/contexts"
)

// Error embeds the variadic list of errors in the context
func Error(r *http.Request, errors ...error) {
	r.WithContext(contexts.ErrorsWithContext(r.Context(), errors...))
}
