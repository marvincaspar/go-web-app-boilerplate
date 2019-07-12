package infra

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Logger represents a logger
type Logger struct {
	log *logrus.Logger
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
		log: logger,
	}
}

// SetOutput sets the standard logger output
func (l *Logger) SetOutput(out io.Writer) {
	l.log.SetOutput(out)
}

// WithFields creates an entry from the standard logger and adds multiple fields to it
func (l *Logger) WithFields(fields interface{}) *logrus.Entry {
	return l.log.WithFields(fields.(logrus.Fields))
}

// Debug logs a message at level Debug on the standard logger
func (l *Logger) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

// Info logs a message at level Info on the standard logger
func (l *Logger) Info(args ...interface{}) {
	l.log.Info(args...)
}

// Warn logs a message at level Warn on the standard logger
func (l *Logger) Warn(args ...interface{}) {
	l.log.Warn(args...)
}

// Error logs a message at level Error on the standard logger
func (l *Logger) Error(args ...interface{}) {
	l.log.Info(args...)
}

// Fatal logs a message at level Fatal on the standard logger
func (l *Logger) Fatal(args ...interface{}) {
	l.log.Fatal(args...)
}

// Panic logs a message at level Panic on the standard logger
func (l *Logger) Panic(args ...interface{}) {
	l.log.Panic(args...)
}
