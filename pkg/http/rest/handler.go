package rest

import (
	"log"

	"github.com/gorilla/mux"
)

// Handler handles http rest requests
type Handler struct {
	logger *log.Logger
	router *mux.Router
}

// HTTPError data model for http error
type HTTPError struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

// CreateHandler create a new http rest handler
func CreateHandler(l *log.Logger) *Handler {
	h := &Handler{
		logger: l,
		router: mux.NewRouter(),
	}
	// mw := middleware.CreateMiddleware(h.logger)

	// h.router.Use(mw.JSONResponse)
	// h.router.Use(mw.Logging)
	// h.router.Use(middleware.CORS)

	return h
}

// GetRouter returns the router
func (h *Handler) GetRouter() *mux.Router {
	return h.router
}
