package logger

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// getCallerInfo retrieves caller information with relative path from project root
func getCallerInfo() (string, string, int) {
	for skip := 3; skip <= 6; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			continue
		}

		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}

		funcName := fn.Name()
		if strings.Contains(funcName, "/logger.") ||
			strings.Contains(funcName, "runtime.") ||
			strings.Contains(funcName, "testing.") {
			continue
		}

		relativePath := getRelativePath(file)

		parts := strings.Split(funcName, "/")
		if len(parts) > 0 {
			lastPart := parts[len(parts)-1]
			if idx := strings.Index(lastPart, "."); idx != -1 {
				funcName = lastPart
			}
		}

		return funcName, relativePath, line
	}

	return "", "", 0
}

func getRelativePath(fullPath string) string {
	rootIndicators := []string{
		"compose.yaml",
	}

	dir := fullPath
	for {
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			return filepath.Base(fullPath)
		}

		for _, indicator := range rootIndicators {
			indicatorPath := filepath.Join(parentDir, indicator)
			if _, err := os.Stat(indicatorPath); err == nil {
				relPath, err := filepath.Rel(parentDir, fullPath)
				if err == nil {
					return relPath
				}
			}
		}

		dir = parentDir
	}
}
