package server

import (
	"net/http"
	"time"
)

const PORT = ":8080"

type Server struct {
	Server *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(handler http.Handler) error {
	s.Server = &http.Server{
		Addr:           PORT,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.Server.ListenAndServe()
}
