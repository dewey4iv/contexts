package httpctx

import (
	"net/http"
	"time"

	"github.com/dewey4iv/contexts"
)

// Time adds a contects.Time to an http context
func Time(r *http.Request, timer string, start time.Time, end time.Time) {
	newTimes := contexts.Times(make(map[string]contexts.Time))
	newTimes[timer] = contexts.Time{
		Start:    start,
		End:      end,
		Duration: time.Duration(end.UnixNano() - start.UnixNano()),
	}

	(*r) = *r.WithContext(contexts.TimesWithContext(r.Context(), newTimes))
}

// Times adds contects.Times to an http context
func Times(r *http.Request, newTimer contexts.Times) {
	prevContexts := contexts.TimesFromContext(r.Context())

	for k, v := range newTimer {
		prevContexts[k] = v
	}

	(*r) = *r.WithContext(contexts.TimesWithContext(r.Context(), prevContexts))
}
