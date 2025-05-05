package main

import (
	"flag"
	"log"

	"github.com/igor-benko/pow-tcp-server/internal/app/server"
	"github.com/igor-benko/pow-tcp-server/internal/config"
)

var defaultEnvFileName = ".env"

func main() {
	flag.Parse()

	cfg, err := config.Init(defaultEnvFileName)
	if err != nil {
		log.Fatal(err)
	}

	server.Run(cfg)
}
