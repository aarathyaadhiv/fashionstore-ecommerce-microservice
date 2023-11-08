package api

import (
	"log"
	"net"

	"github.com/aarathyaadhiv/fashionstore-user-service/pkg/api/service"
	"github.com/aarathyaadhiv/fashionstore-user-service/pkg/config"
	"github.com/aarathyaadhiv/fashionstore-user-service/pkg/pb"
	"google.golang.org/grpc"
)


type ServerHTTP struct{
	engine *grpc.Server
}

func NewServerHTTP(userService *service.UserService)*ServerHTTP{
	engine:=grpc.NewServer()

	pb.RegisterUserServiceServer(engine,userService)
	return &ServerHTTP{engine: engine}
}

func (s *ServerHTTP) Start(c config.Config){
	lis,err:=net.Listen("tcp",c.Port)
	if err!=nil{
		log.Fatalln("failed to listen",err)
	}
	if err=s.engine.Serve(lis);err!=nil{
		log.Fatalln("failed to serve",err)
	}
}