package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/marvincaspar/go-web-app-boilerplate/test"

	"github.com/gorilla/mux"
)

func TestLoggingMiddleware(t *testing.T) {
	logger, reader, writer := test.LoggerWithOutputCapturingMock()

	router := mux.NewRouter()
	router.HandleFunc("/", test.HandlerMock).Methods("GET")
	rr := httptest.NewRecorder()
	mw := &Middleware{
		logger: logger,
	}

	// Add the middleware again as function
	router.Use(mw.Logging)
	req, _ := http.NewRequest("GET", "/", nil)

	logResult := test.CaptureLogOutput(reader, writer, func() {
		router.ServeHTTP(rr, req)
	})

	expected := `"component":"Middleware","level":"info","method":"GET","msg":"Request logging","res[content-length]":0,"response-time":"0ms","status":200,"time":`
	if !strings.Contains(logResult, expected) {
		t.Errorf("request was not logged: got %v want %v",
			logResult, expected)
	}
}
