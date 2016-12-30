package httpctx

import (
	"encoding/json"
	"net/http"

	"github.com/dewey4iv/contexts"
)

// Response is a generic http response
type Response struct {
	Meta     map[string]interface{} `json:"__meta,omitempty"`
	Payload  interface{}            `json:"payload,omitempty"`
	Messages []string               `json:"messages,omitempty"`
	Warnings []string               `json:"warnings,omitempty"`
	Errors   []error                `json:"errors,omitempty"`
}

// RenderContexts is an http middleware for json.
// It looks at all of the possible contexts and renders them to the ResponseWriter
func RenderContexts() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)

			var response Response

			if errors := contexts.ErrorsFromContext(r.Context()); errors != nil {
				response.Errors = errors
			}

			if err := json.NewEncoder(w).Encode(response); err != nil {
				// TODO: Handle Err Here
			}
		})
	}
}
