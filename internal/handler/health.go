package handler

import (
	"net/http"

	"github.com/OderoCeasar/goapi/internal/config"
	"github.com/OderoCeasar/goapi/pkg/response"
)

type Handler struct {
	config *config.Config
}


// new handler
func New(cfg *config.Config) *Handler {
	return &Handler{config: cfg}
}

// RegisterRoutes wires up all the HTTP rooutes to handlers
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", h.HealthCheck)
	mux.HandleFunc("GET /api/v1/ping", h.Ping)
	mux.HandleFunc("GET /api/v1/status", h.Status)
	mux.HandleFunc("GET /api/v1/info", h.Info)
}


// HealthCheck - used by load balancers and orchestrators(EC2)
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{
		"status": "healthy",
		"service": h.config.AppName,
	})
}


// Ping is a simple endpoint to test routing and response formatting
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{
		"message": "pong",
	})
}

// Status - new endpoint to return the status of the project
func (h *Handler) Status(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{
		"environment" : h.config.Env,
		"version": h.config.Version,
	})
}

// Info simple endpoint for testing routes(getting app info)
func (h *Handler) Info(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{
		
		"app_name": h.config.AppName,
		"app_environment" : h.config.Env,
	})
}