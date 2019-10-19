package test

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"

	"github.com/marvincaspar/go-web-app-boilerplate/pkg/infra"
)

// LoggerMock creates a log mock
func LoggerMock() (*infra.Logger, *observer.ObservedLogs) {
	core, recorded := observer.New(zapcore.InfoLevel)
	logger := &infra.Logger{Log: zap.New(core)}
	return logger, recorded
}
