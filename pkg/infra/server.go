package infra

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)


// CreateServer create a new http server
func CreateServer(router *mux.Router, serverAddress string, options ...func(*http.Server)) *http.Server {
	server := &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      router,
	}
	for _, option := range(options) {
		option(server)
	}
	return server
}
