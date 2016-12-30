package contexts

import "context"

// MessagesCtxKey is the contexts.Key for Messages
const MessagesCtxKey Key = "MessagesCtxKey"

// MessagesWithContext with contexts adds messages to the provided context
func MessagesWithContext(ctx context.Context, messages ...string) context.Context {
	if prevMessages, ok := ctx.Value(MessagesCtxKey).([]string); ok {
		messages = append(messages, prevMessages...)
	}

	return context.WithValue(ctx, MessagesCtxKey, messages)
}

// MessagesFromContext returns messages from the provided context
func MessagesFromContext(ctx context.Context) []string {
	if messages, ok := ctx.Value(MessagesCtxKey).([]string); ok {
		return messages
	}

	return nil
}
