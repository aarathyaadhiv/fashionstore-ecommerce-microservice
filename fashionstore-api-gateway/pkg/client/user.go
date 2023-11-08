package client

import (
	"fmt"

	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/config"
	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


type UserClient struct{
	Client pb.UserServiceClient
}

func NewUserClient(c config.Config)*UserClient{
	cc,err:=grpc.Dial(c.UserSvcUrl,grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err!=nil{
		fmt.Println("could not connect at user client:",err)
		
	}

	return &UserClient{
		Client: pb.NewUserServiceClient(cc),
	}
	
}

