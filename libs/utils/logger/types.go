package logger

// LogLevel represents logging levels
type LogLevel int

const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
)

// String representation of log levels
func (l LogLevel) String() string {
	switch l {
	case LevelTrace:
		return "trace"
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	default:
		return "unknown"
	}
}

// Logger handles structured logging with tag-based filtering
type Logger struct {
	service     string
	version     string
	environment string
	enabledTags map[string]bool
	logLevel    LogLevel
	devMode     bool
}

// Log entry structure
type LogEntry struct {
	Timestamp   string                 `json:"@timestamp"`
	Level       string                 `json:"level"`
	Service     string                 `json:"service"`
	Version     string                 `json:"version,omitempty"`
	Environment string                 `json:"environment,omitempty"`
	Message     string                 `json:"message"`
	Tags        []string               `json:"tags"`
	RequestID   string                 `json:"request_id,omitempty"`
	UserID      string                 `json:"user_id,omitempty"`
	Function    string                 `json:"function,omitempty"`
	File        string                 `json:"file,omitempty"`
	Line        int                    `json:"line,omitempty"`
	Error       string                 `json:"error,omitempty"`
	Fields      map[string]interface{} `json:"fields,omitempty"`
}

// Context keys
type contextKey string

const (
	RequestIDKey contextKey = "request_id"
	UserIDKey    contextKey = "user_id"
)
