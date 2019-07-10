package test

import (
	"bytes"
	"io"
	"os"
	"sync"

	"github.com/marvincaspar/go-web-app-boilerplate/pkg/infra"
	"github.com/sirupsen/logrus"
)

// LoggerMock creates a log mock
func LoggerMock() *infra.Logger {
	logger := logrus.New()

	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	logger.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logger.SetLevel(logrus.InfoLevel)

	return &infra.Logger{
		Log: logger,
	}
}

// LoggerWithOutputCapturingMock creates a log mock which can capture the log rsult
func LoggerWithOutputCapturingMock() (*infra.Logger, *os.File, *os.File) {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	logger := LoggerMock()
	logger.Log.SetOutput(writer)

	return logger, reader, writer
}

// CaptureLogOutput writes log output to a string
// https://medium.com/@hau12a1/golang-capturing-log-println-and-fmt-println-output-770209c791b4
func CaptureLogOutput(reader *os.File, writer *os.File, f func()) string {
	os.Stdout = writer
	os.Stderr = writer

	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}
