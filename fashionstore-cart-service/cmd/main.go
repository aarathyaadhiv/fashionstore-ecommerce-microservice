package main

import (
	"log"

	"github.com/aarathyaadhiv/fashionstore-cart-service/pkg/config"
	"github.com/aarathyaadhiv/fashionstore-cart-service/pkg/di"
)



func main(){
	config,configErr:=config.LoadConfig()
	if configErr!=nil{
		log.Fatalln("failed to load config ",configErr)
	}
	server,diErr:=di.InitializeAPI(config)
	if diErr!=nil{
		log.Fatalln("failed in initialization",diErr)
	}else{
		server.Start(config)
	}

}