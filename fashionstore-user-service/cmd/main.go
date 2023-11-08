package main

import (
	"log"

	"github.com/aarathyaadhiv/fashionstore-user-service/pkg/config"
	"github.com/aarathyaadhiv/fashionstore-user-service/pkg/di"
)

func main() {
	config, confiErr := config.LoadConfig()
	if confiErr != nil {
		log.Fatalln("config error", confiErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatalln("server error", diErr)
	} else {
		server.Start(config)
	}
}
