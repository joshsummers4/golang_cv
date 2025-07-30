package logger

import "context"

// WithRequestID adds a request ID to the context
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIDKey, requestID)
}

// WithUserID adds a user ID to the context
func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

// GetRequestID retrieves the request ID from context
func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if reqID := ctx.Value(RequestIDKey); reqID != nil {
		if id, ok := reqID.(string); ok {
			return id
		}
	}
	return ""
}

// GetUserID retrieves the user ID from context
func GetUserID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if userID := ctx.Value(UserIDKey); userID != nil {
		if id, ok := userID.(string); ok {
			return id
		}
	}
	return ""
}
