package httpctx

import (
	"net/http"

	"github.com/dewey4iv/contexts"
)

// Meta adds the val with the coreesponding key
func Meta(r *http.Request, key string, val interface{}) {
	newMap := make(map[string]interface{})
	newMap[key] = val

	r.WithContext(contexts.MetaWithContext(r.Context(), newMap))
}
