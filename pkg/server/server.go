package server

import (
	"net/http"
	"time"
)

type Server struct {
	Server *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(handler http.Handler) error {
	s.Server = &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.Server.ListenAndServe()
}
