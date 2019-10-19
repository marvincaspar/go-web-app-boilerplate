package middleware

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/marvincaspar/go-web-app-boilerplate/test"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/gorilla/mux"
)

func TestLoggingMiddleware(t *testing.T) {
	logger, recorded := test.LoggerMock()

	router := mux.NewRouter()
	router.HandleFunc("/", test.HandlerMock).Methods("GET")
	rr := httptest.NewRecorder()
	mw := &Middleware{
		logger: logger,
	}

	// Add the middleware again as function
	router.Use(mw.Logging)
	req := httptest.NewRequest("GET", "/", nil)
	router.ServeHTTP(rr, req)

	expect := []zapcore.Field{
		zap.String("component", "Middleware"),
		zap.String("method", "GET"),
		zap.String("url", "/"),
		zap.Int("status", http.StatusOK),
		zap.Int64("res[content-length]", 0),
		zap.String("response-time", "0ms"),
	}

	for _, logs := range recorded.All() {
		if !reflect.DeepEqual(expect, logs.Context) {
			t.Errorf("request was not logged: got %v want %v", logs.Context, expect)
		}
	}
}
