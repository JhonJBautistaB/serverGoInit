package main

import (
	"serverGoInit/cmd/server"
	"serverGoInit/config"
)

func main() {

	cfg := config.LoadConfig()
	srv := server.NewServer(cfg.PortServer, cfg.NameServer)
	if err := srv.StartServer(); err != nil {
		panic(err)
	}

}
