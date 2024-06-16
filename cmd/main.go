package main

import (
	"golangMSWorkingChi/cmd/server"
	"golangMSWorkingChi/config"
)

func main() {

	cfg := config.LoadConfig()
	srv := server.NewServer(cfg.PortServer, cfg.NameServer)
	if err := srv.StartServer(); err != nil {
		panic(err)
	}

}
