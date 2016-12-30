package contexts

import (
	"context"
	"time"
)

// TimesCtxKey is the contects.Key for Times
const TimesCtxKey Key = "TimesCtxKey"

// Times is what becomes embedded in the context
type Times map[string]Time

// Time holds a start and end time
type Time struct {
	Start    time.Time     `json:"start"`
	End      time.Time     `json:"end"`
	Duration time.Duration `json:"duration"`
}

// TimesWithContext merges the provided context and supplied timesdewey4
func TimesWithContext(ctx context.Context, newTimes Times) context.Context {
	prevContexts := TimesFromContext(ctx)

	if prevContexts == nil {
		prevContexts = Times(make(map[string]Time))
	}

	for k, v := range newTimes {
		prevContexts[k] = v
	}

	return context.WithValue(ctx, TimesCtxKey, prevContexts)
}

// TimesFromContext returns the Times from a context
func TimesFromContext(ctx context.Context) Times {
	if times, ok := ctx.Value(TimesCtxKey).(Times); ok {
		return times
	}

	return nil
}
