package contexts

import "context"

// MetaCtxKey is the contexts.Key for Messages
const MetaCtxKey Key = "MetaCtxKey"

// MetaWithContext merges the meta with the provided context
func MetaWithContext(ctx context.Context, newMeta map[string]interface{}) context.Context {
	prevMeta := MetaFromContext(ctx)

	if prevMeta == nil {
		prevMeta = make(map[string]interface{})
	}

	for k, v := range newMeta {
		prevMeta[k] = v
	}

	return context.WithValue(ctx, MetaCtxKey, prevMeta)
}

// MetaFromContext pulls the Meta from the prvided context
func MetaFromContext(ctx context.Context) map[string]interface{} {
	if meta, ok := ctx.Value(MetaCtxKey).(map[string]interface{}); ok {
		return meta
	}

	return nil
}
