package main

import (
	"flag"
	"log"

	"github.com/igor-benko/pow-tcp-server/internal/app/client"
	"github.com/igor-benko/pow-tcp-server/internal/config"
)

var defaultEnvFileName = ".env"

func main() {
	flag.Parse()

	cfg, err := config.Init(defaultEnvFileName)
	if err != nil {
		log.Fatal(err)
	}

	client.Run(cfg)
}
