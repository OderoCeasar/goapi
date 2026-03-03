package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/OderoCeasar/goapi/internal/config"
	"github.com/OderoCeasar/goapi/internal/handler"
	"github.com/OderoCeasar/goapi/internal/midleware"
)

// server wraps the HTTP server and its dependencies
type Server struct {
	httpServer  *http.Server
	config 		*config.Config
}

// New creates a configured server instance
func New(cfg *config.Config) *Server {
	// create the router and register routes
	mux := http.NewServeMux()

	h := handler.New(cfg)
	h.RegisterRoutes(mux)

	chain := midleware.Recovery(
		midleware.RequestID(
			midleware.Logger(mux),
		),
	)
	
	// timeouts - without the timeouts slow clients can exhaust your server connection
	httpServer := &http.Server{
		Addr: 			fmt.Sprintf(":%s", cfg.Port),
		Handler: 		chain,
		ReadTimeout: 	15 * time.Second,
		WriteTimeout: 	15 * time.Second,
		IdleTimeout: 	60 * time.Second,
	}

	return &Server{
		httpServer: httpServer,
		config: cfg,
	}
}


// start begins listening for requests
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}