package middleware

import (
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// Logging logs the request time
func (mw *Middleware) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {

			duration := time.Since(start)
			durationInMs := duration.Nanoseconds() / 1000000
			mw.logger.WithFields(log.Fields{
				"component":           "Middleware",
				"method":              r.Method,
				"url":                 r.URL.Path,
				"status":              200,
				"res[content-length]": r.ContentLength,
				"response-time":       strconv.Itoa(int(durationInMs)) + "ms",
			}).Info("Request logging")
		}()

		next.ServeHTTP(w, r)
	})
}
