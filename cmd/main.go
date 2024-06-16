package main

import (
	"golangMSWorkingChi/config"
	"golangMSWorkingChi/server"
)

func main() {

	cfg := config.LoadConfig()
	srv := server.NewServer(cfg.PortServer, cfg.NameServer)
	if err := srv.StartServer(); err != nil {
		panic(err)
	}

}
