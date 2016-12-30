package httpctx

import (
	"net/http"

	"github.com/dewey4iv/contexts"
)

// Payload adds a payload to an http context
func Payload(r *http.Request, payload interface{}) {
	r.WithContext(contexts.PayloadWithContext(r.Context(), payload))
}
