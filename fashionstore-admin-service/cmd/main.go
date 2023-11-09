package main

import (
	"log"

	"github.com/aarathyaadhiv/fashionstore-admin-service/pkg/config"
	"github.com/aarathyaadhiv/fashionstore-admin-service/pkg/di"
)




func main(){
	config,confiErr:=config.LoadConfig()

	if confiErr!=nil{
		log.Fatalln("failiure in loading config",confiErr)
	}

	server,diErr:=di.InitializeAPI(config)

	if diErr!=nil{
		log.Fatalln("initialization error",diErr)
	}else{
		server.Start(config)
	}

}