package contexts

import "context"

// PayloadCtxKey is the contexts.Key for Messages
const PayloadCtxKey Key = "PayloadCtxKey"

// PayloadWithContext sets the payload on a context
func PayloadWithContext(ctx context.Context, payload interface{}) context.Context {
	return context.WithValue(ctx, PayloadCtxKey, payload)
}

// PayloadFromContext returns the payload in the context
func PayloadFromContext(ctx context.Context) interface{} {
	if payload := ctx.Value(PayloadCtxKey); payload != nil {
		return payload
	}

	return nil
}
