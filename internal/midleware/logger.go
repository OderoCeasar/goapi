package midleware

import (
	"log/slog"
	"net/http"
	"time"
)

// responseWriter wraps http.ResponseWriter to capture the status code
type responseWriter struct {
	http.ResponseWriter
	status int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}


func (rw *responseWriter) WriteHeader(code int) {
	if !rw.wroteHeader {
		rw.status = code
		rw.wroteHeader = true
		rw.ResponseWriter.WriteHeader(code)
	}
}

// status returns the captured status code
func (rw *responseWriter) Status() int {
	if rw.status == 0 {
		return http.StatusOK
	}

	return rw.status
}


// Logger logs every request with method, path, status, duration and requestID
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := wrapResponseWriter(w)

		// let the handle run
		next.ServeHTTP(wrapped, r)

		// after handler completes log everything we captured
		duration := time.Since(start)
		requestID := GetRequestID(r.Context())

		// slog => structured logging(key value pairs that can parse, filter)
		slog.Info("request completed",
			"method", 	r.Method,
			"path",		r.URL.Path,
			"status", 	wrapped.Status(),
			"duration_ms", duration.Milliseconds(),
			"request_id",	requestID,
			"remote_addr", r.RemoteAddr,
		)
	})
}