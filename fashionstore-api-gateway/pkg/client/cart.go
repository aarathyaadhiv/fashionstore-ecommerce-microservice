package client

import (
	"log"

	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/config"
	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


type CartClient struct{
	Client pb.CartServiceClient
}

func NewCartClient(c config.Config)*CartClient{
	cc,err:=grpc.Dial(c.CartSvcUrl,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatalln("client connection failed at cart service",err)
	}
	return &CartClient{Client: pb.NewCartServiceClient(cc)}
}