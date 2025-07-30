package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var l *Logger

func newLogger(service, version, environment string, tags []string, logLevel LogLevel, devMode bool) *Logger {
	enabledTags := make(map[string]bool, len(tags))
	for _, tag := range tags {
		enabledTags[tag] = true
	}

	return &Logger{
		service:     service,
		version:     version,
		environment: environment,
		enabledTags: enabledTags,
		logLevel:    logLevel,
		devMode:     devMode,
	}
}

func (logger *Logger) shouldLog(level LogLevel, tags []string) bool {
	if level < logger.logLevel {
		return false
	}

	if level == LevelError {
		return true
	}

	if len(logger.enabledTags) == 0 {
		return true
	}

	for _, tag := range tags {
		if logger.enabledTags[tag] {
			return true
		}
	}

	return false
}

func writeToLog(ctx context.Context, level LogLevel, message string, tags []string, err error, fields map[string]interface{}) {
	// if !l.shouldLog(level, tags) {
	// 	return
	// }

	entry := LogEntry{
		Timestamp:   time.Now().UTC().Format(time.RFC3339Nano),
		Level:       level.String(),
		Service:     l.service,
		Version:     l.version,
		Environment: l.environment,
		Message:     message,
		Tags:        tags,
		Fields:      fields,
	}

	if ctx != nil {
		if reqID := ctx.Value(RequestIDKey); reqID != nil {
			if id, ok := reqID.(string); ok {
				entry.RequestID = id
			}
		}
		if userID := ctx.Value(UserIDKey); userID != nil {
			if id, ok := userID.(string); ok {
				entry.UserID = id
			}
		}
	}

	entry.Function, entry.File, entry.Line = getCallerInfo()

	if err != nil {
		entry.Error = err.Error()
	}

	jsonData, _ := json.Marshal(entry)
	fmt.Fprintln(os.Stdout, string(jsonData))

}

func Init(service, version, environment string, tags []string, logLevel LogLevel) {
	devMode := environment == "dev"
	l = newLogger(service, version, environment, tags, logLevel, devMode)
}
