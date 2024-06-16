package server

import (
	"log"
	"net/http"
)

type Server struct {
	Port       string
	NameServer string
}

// created method constructor
func NewServer(port string, nameServer string) *Server {
	return &Server{
		Port:       port,
		NameServer: nameServer,
	}
}

// created method init server
func (s *Server) StartServer() error {
	srv := &http.Server{
		Addr:    ":" + s.Port,
		Handler: nil,
	}

	log.Printf("Server %s running on port %s", s.NameServer, s.Port)
	return srv.ListenAndServe()

}
