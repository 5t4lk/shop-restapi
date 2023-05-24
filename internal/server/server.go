package server

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(port string, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:                         "",
			Handler:                      nil,
			DisableGeneralOptionsHandler: false,
			TLSConfig:                    nil,
			ReadTimeout:                  0,
			ReadHeaderTimeout:            0,
			WriteTimeout:                 0,
			IdleTimeout:                  0,
			MaxHeaderBytes:               0,
			TLSNextProto:                 nil,
			ConnState:                    nil,
			ErrorLog:                     nil,
			BaseContext:                  nil,
			ConnContext:                  nil,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
