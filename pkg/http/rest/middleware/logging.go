package middleware

import (
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"
)

// Logging logs the request time
func (mw *Middleware) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {

			duration := time.Since(start)
			durationInMs := duration.Nanoseconds() / 1000000
			mw.logger.WithFields(
				zap.String("component", "Middleware"),
				zap.String("method", r.Method),
				zap.String("url", r.URL.Path),
				zap.Int("status", http.StatusOK),
				zap.Int64("res[content-length]", r.ContentLength),
				zap.String("response-time", strconv.Itoa(int(durationInMs))+"ms"),
			).Info("Request logging")
		}()

		next.ServeHTTP(w, r)
	})
}
