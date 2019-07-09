package infra

import (
	"log"
	"os"
)

// CreateLogger creates a logger instance for all components
func CreateLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}
