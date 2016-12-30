package httpctx

import (
	"net/http"

	"github.com/dewey4iv/contexts"
)

// Messages adds messages to the provided http context
func Messages(r *http.Request, messages ...string) {
	r.WithContext(contexts.MessagesWithContext(r.Context(), messages...))
}
