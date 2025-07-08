package http

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logg.Info(
			"started",
			zap.String("method", r.Method),
			zap.String("uri", r.RequestURI),
			zap.String("IP", r.RemoteAddr),
			zap.String("user_agent", r.UserAgent()),
		)
		start := time.Now()
		next.ServeHTTP(w, r)
		s.logg.Info(
			"completed",
			zap.String("method", r.Method),
			zap.String("uri", r.RequestURI),
			zap.String("IP", r.RemoteAddr),
			zap.String("user_agent", r.UserAgent()),
			zap.Time("request_datetime", start),
			zap.Duration("duration", time.Since(start)),
		)
	})
}
