package infra

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger represents a logger
type Logger struct {
	Log *logrus.Logger
}

// CreateLogger creates a logger instance for all components
func CreateLogger() *Logger {
	logger := logrus.New()

	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	logger.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logger.SetLevel(logrus.InfoLevel)

	return &Logger{
		Log: logger,
	}
}
