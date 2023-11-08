package main

import (
	
	"log"

	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/config"
	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/di"
)
	


func main(){
	config,configErr:=config.LoadConfig()

	if configErr!=nil{
		log.Fatal("configuration error:",configErr)
	}

	server,diErr:=di.InitializeAPI(config)
	if diErr!=nil{
		log.Fatal("cannot start serrver",diErr)
	}else{
		server.Start(config)
	}
	
}