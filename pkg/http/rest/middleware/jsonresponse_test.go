package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/marvincaspar/go-web-app-boilerplate/test"
	"github.com/gorilla/mux"
)

func TestJSONResponseMiddleware(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/", test.HandlerMock).Methods("GET")
	rr := httptest.NewRecorder()
	mw := &Middleware{}

	// Add the middleware again as function
	router.Use(mw.JSONResponse)
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(rr, req)

	expected := "application/json"
	if contentType := rr.Header().Get("Content-type"); contentType != expected {
		t.Errorf("request header contains wrong content type: got %v want %v",
			contentType, expected)
	}

}
