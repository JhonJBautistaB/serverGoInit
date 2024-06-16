package server

import (
	"log"
	"net/http"

	"serverGoInit/cmd/routes"
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
	rts := routes.Routes{}
	srv := &http.Server{
		Addr:    ":" + s.Port,
		Handler: rts.Routes(),
	}

	log.Printf("Server %s running on port %s", s.NameServer, s.Port)
	return srv.ListenAndServe()

}
