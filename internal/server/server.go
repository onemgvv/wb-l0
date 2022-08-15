package server

import (
	"context"
	"github.com/onemgvv/wb-l0/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		&http.Server{
			Addr:           ":" + cfg.HTTP.Port,
			Handler:        handler,
			MaxHeaderBytes: cfg.HTTP.MaxHeaderMegabytes,
			ReadTimeout:    cfg.HTTP.Timeout.Read,
			WriteTimeout:   cfg.HTTP.Timeout.Write,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
