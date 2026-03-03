package midleware

import (
	"log/slog"
	"net/http"

	"github.com/OderoCeasar/goapi/pkg/response"
)

// Recovery catches panics and returns a 500 instead of crashing the server
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func () {
			if err := recover(); err != nil {
				requestID := GetRequestID(r.Context())

				slog.Error("panic recovered",
					"error",   err,
					"request_id", requestID,
					"path", 	r.URL.Path,
					"method", r.Method,
				)

				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
		} ()

		next.ServeHTTP(w, r)
	})
}