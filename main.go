package main

import (
	"github.com/lai0xn/docsoft/config"
	"github.com/lai0xn/docsoft/internal/server"
)

func main() {
	config.LoadENV()
	srv := server.Server{
		ADDR: config.SERVER_PORT,
	}
	srv.Run()
}
