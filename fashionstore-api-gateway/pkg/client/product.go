package client

import (
	"log"

	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/config"
	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


type ProductClient struct{
	Client pb.ProductServiceClient
}

func NewProductClient(c config.Config)*ProductClient{
	cc,err:=grpc.Dial(c.ProductSvcUrl,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatalln("client connection failiure",err)
	}
	return &ProductClient{
		Client: pb.NewProductServiceClient(cc),
	}
}