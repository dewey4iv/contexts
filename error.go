package contexts

import "context"

// ErrorsCtxKey is the contexts.Key for Errors
const ErrorsCtxKey Key = "ErrorsCtxKey"

// ErrorsWithContext takes a context and adds errors
func ErrorsWithContext(ctx context.Context, errors ...error) context.Context {
	if prevErrors, ok := ctx.Value(ErrorsCtxKey).([]error); ok {
		errors = append(errors, prevErrors...)
	}

	return context.WithValue(ctx, ErrorsCtxKey, errors)
}

// ErrorsFromContext pulls errors from the given context
func ErrorsFromContext(ctx context.Context) []error {
	if errors, ok := ctx.Value(ErrorsCtxKey).([]error); ok {
		return errors
	}

	return nil
}
