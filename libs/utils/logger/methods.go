package logger

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// Error logs an error level message
func Error(ctx context.Context, message string, err error, tags []string, fields ...map[string]any) {
	var f map[string]any
	if len(fields) > 0 {
		f = fields[0]
	}
	writeToLog(ctx, LevelError, message, tags, err, f)
}

// Warn logs a warning level message
func Warn(ctx context.Context, message string, tags []string, fields ...map[string]any) {
	var f map[string]any
	if len(fields) > 0 {
		f = fields[0]
	}
	writeToLog(ctx, LevelWarn, message, tags, nil, f)
}

// Info logs an info level message
func Info(ctx context.Context, message string, tags []string, fields ...map[string]any) {
	var f map[string]any
	if len(fields) > 0 {
		f = fields[0]
	}
	writeToLog(ctx, LevelInfo, message, tags, nil, f)
}

// Debug logs a debug level message
func Debug(ctx context.Context, message string, tags []string, fields ...map[string]any) {
	var f map[string]any
	if len(fields) > 0 {
		f = fields[0]
	}
	writeToLog(ctx, LevelDebug, message, tags, nil, f)
}

// Trace logs a trace level message
func Trace(ctx context.Context, message string, tags []string, fields ...map[string]any) {
	var f map[string]any
	if len(fields) > 0 {
		f = fields[0]
	}
	writeToLog(ctx, LevelTrace, message, tags, nil, f)
}

func DebugSQLQuery(query string, results int, duration_ms int64, params ...interface{}) map[string]any {
	res := map[string]any{
		"results":     results,
		"duration_ms": duration_ms,
	}
	debugQuery := query

	for _, param := range params {
		// Convert parameter to string representation based on type
		var paramStr string
		switch v := param.(type) {
		case string:
			paramStr = fmt.Sprintf("'%s'", strings.Replace(v, "'", "''", -1))
		case []byte:
			paramStr = fmt.Sprintf("'%s'", strings.Replace(string(v), "'", "''", -1))
		case time.Time:
			paramStr = fmt.Sprintf("'%s'", v.Format("2006-01-02 15:04:05"))
		case nil:
			paramStr = "NULL"
		default:
			paramStr = fmt.Sprintf("%v", v)
		}

		// Find the first ? and replace it
		pos := strings.Index(debugQuery, "?")
		if pos != -1 {
			debugQuery = debugQuery[:pos] + paramStr + debugQuery[pos+1:]
		}
	}

	res["query"] = debugQuery

	return res
}
