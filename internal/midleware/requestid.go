package midleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// contextkey is an unexported type for context keys in this package
type contextkey string

const RequestIDKey contextkey = "request_id"

// RequestID injects a unique ID into every request
func RequestID(next http.Handler)  http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		w.Header().Set("X-Request-ID", requestID)

		// inject into context so handlers can read it
		ctx := context.WithValue(r.Context(), RequestIDKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetRequestID retrieves the request ID from context
func GetRequestID(ctx context.Context) string {
	if id, ok := ctx.Value(RequestIDKey).(string); ok {
		return id
	}
	return ""
}
