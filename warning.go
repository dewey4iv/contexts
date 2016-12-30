package contexts

import "context"

// WarningsCtxKey is the contexts.Key for Messages
const WarningsCtxKey Key = "WarningsCtxKey"

// WarningsWithContext add warnings to the provided context
func WarningsWithContext(ctx context.Context, warnings ...string) context.Context {
	if prevWarnings, ok := ctx.Value(WarningsCtxKey).([]string); ok {
		warnings = append(warnings, prevWarnings...)
	}

	return context.WithValue(ctx, WarningsCtxKey, warnings)
}

// WarningsFromContext pulls warnings from the provided context
func WarningsFromContext(ctx context.Context) []string {
	if warnings, ok := ctx.Value(WarningsCtxKey).([]string); ok {
		return warnings
	}

	return nil
}
