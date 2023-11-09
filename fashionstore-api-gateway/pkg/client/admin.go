package client

import (
	"log"

	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/config"
	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)



type AdminClient struct{
	Client pb.AdminServiceClient
}

func NewAdminClient(c config.Config)*AdminClient{
	cc,err:=grpc.Dial(c.AdminSvcUrl,grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err!=nil{
		log.Fatalln("client connection failed for admin svc",err)
	}

	return &AdminClient{
		Client: pb.NewAdminServiceClient(cc),
	}

}