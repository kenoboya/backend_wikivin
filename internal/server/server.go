package server

import (
	"context"
	"net/http"
	"wikivin/internal/config"
	rest "wikivin/internal/transports/http"
)

type Server struct {
	server *http.Server
}

func NewServer(config *config.Config, handler *rest.Handler) *Server{
	return &Server{
		server: &http.Server{
			Addr: config.HTTP.Addr,
			Handler: handler.Init(config),
			ReadTimeout: config.HTTP.ReadTimeout,
			WriteTimeout: config.HTTP.WriteTimeout,
			MaxHeaderBytes: config.HTTP.MaxHeaderBytes,
		},
	}
}
func (s Server) Run() error{
	if err:= s.server.ListenAndServe(); err != nil{
		return err 
	}
	return nil
}

func (s Server) ShutDown(ctx context.Context) error{
	if err:= s.server.Shutdown(ctx); err != nil{
		return err 
	}
	return nil
}