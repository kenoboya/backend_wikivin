package server

import (
	"context"
	"log"
	"net/http"
	"wikivin/internal/config"
	rest "wikivin/internal/transports/http"
)

type Server struct {
	server *http.Server
}

func NewServer(config *config.Config, handler rest.Handler) *Server{
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
		log.Fatal("The server didn't start")
		return err // FATAL
	}
	return nil
}

func (s Server) ShutDown(ctx context.Context) error{
	if err:= s.server.Shutdown(ctx); err != nil{
		log.Fatal("The server couldn't shut down")
		return err // FATAL
	}
	return nil
}