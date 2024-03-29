package middlewares

import (
	"github.com/Achmadqizwini/SportKai/utils/logger"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func Logging(mux http.Handler) http.Handler {
	log := logger.NewLogger().Logger // Initialize logger

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrappedHead := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		mux.ServeHTTP(wrappedHead, r)

		// Log request details
		log.Info().
			Int("status_code", wrappedHead.statusCode).
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Dur("duration", time.Since(start)).
			Msg("HTTP request processed")
	})
}
