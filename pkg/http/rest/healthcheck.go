package rest

import (
	"io"
	"net/http"
)

// NewHealthCheckHandler add route for healthcheck
func (h *Handler) NewHealthCheckHandler() {
	h.router.HandleFunc("/health", healthCheckHandler).Methods("GET")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}
