package logger

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	size        int64
	wroteHeader bool
}

func (rw *responseWriter) WriteHeader(status int) {
	if !rw.wroteHeader {
		rw.status = status
		rw.wroteHeader = true
		rw.ResponseWriter.WriteHeader(status)
	}
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	if !rw.wroteHeader {
		rw.status = http.StatusOK
		rw.wroteHeader = true
	}
	n, err := rw.ResponseWriter.Write(data)
	rw.size += int64(n)
	return n, err
}

func (rw *responseWriter) Flush() {
	if flusher, ok := rw.ResponseWriter.(http.Flusher); ok {
		flusher.Flush()
	}
}

func LoggingHandler(excludePaths []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if path should be excluded
			for _, path := range excludePaths {
				if strings.HasPrefix(r.URL.Path, path) {
					next.ServeHTTP(w, r)
					return
				}
			}

			start := time.Now()

			requestID := generateRequestID()
			ctx := WithRequestID(r.Context(), requestID)
			r = r.WithContext(ctx)
			w.Header().Set("X-Request-ID", requestID)

			userID := extractUserID(r)
			if userID != "" {
				ctx = WithUserID(ctx, userID)
				r = r.WithContext(ctx)
			}

			// Wrap response writer to capture details
			wrapped := &responseWriter{
				ResponseWriter: w,
				status:         200, // default status
			}

			// Get client IP
			clientIP := getClientIP(r)

			// Process request
			next.ServeHTTP(wrapped, r)

			// Calculate duration
			duration := time.Since(start)

			// Build log fields
			logFields := map[string]interface{}{
				"method":        r.Method,
				"path":          r.URL.Path,
				"status_code":   wrapped.status,
				"duration_ms":   duration.Milliseconds(),
				"response_size": wrapped.size,
				"remote_ip":     clientIP,
				"user_id":       userID,
				"user_agent":    r.Header.Get("User-Agent"),
			}

			// Add query params if present
			if r.URL.RawQuery != "" {
				logFields["query"] = r.URL.RawQuery
			}

			// Add referer if present
			if referer := r.Header.Get("Referer"); referer != "" {
				logFields["referer"] = referer
			}

			// Single log entry with appropriate tags based on status
			message := fmt.Sprintf("%s %s %d", r.Method, r.URL.Path, wrapped.status)
			var level LogLevel
			var tags []string

			if wrapped.status >= 500 {
				level = LevelError
				tags = []string{"http", "error"}
			} else if wrapped.status >= 400 {
				level = LevelWarn
				tags = []string{"http", "warn"}
			} else {
				level = LevelInfo
				tags = []string{"http"}
			}

			writeToLog(ctx, level, message, tags, nil, logFields)

			if duration > 1*time.Second {
				writeToLog(ctx, LevelWarn, "Slow request detected", []string{"http", "performance"}, nil, map[string]interface{}{
					"method":       r.Method,
					"path":         r.URL.Path,
					"duration_ms":  duration.Milliseconds(),
					"threshold_ms": 1000,
				})
			}
		})
	}
}

func generateRequestID() string {
	return fmt.Sprintf("req-%d", time.Now().UnixNano())
}

func getClientIP(r *http.Request) string {
	// Check X-Forwarded-For header first (load balancers, proxies)
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		// X-Forwarded-For can contain multiple IPs, take the first one
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0])
	}

	// Check X-Real-IP header (Nginx proxy)
	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		return realIP
	}

	// Fallback to RemoteAddr
	ip := r.RemoteAddr
	if idx := strings.LastIndex(ip, ":"); idx != -1 {
		ip = ip[:idx]
	}
	return ip
}

func extractUserID(r *http.Request) string {
	// TODO : Extract user ID from context
	return "no user id found" // No user ID found
}
