package api

import (
	"log"
	"net"

	"github.com/aarathyaadhiv/fashionstore-cart-service/pkg/api/service"
	"github.com/aarathyaadhiv/fashionstore-cart-service/pkg/config"
	"github.com/aarathyaadhiv/fashionstore-cart-service/pkg/pb"
	"google.golang.org/grpc"
)


type ServerHTTP struct{
	engine *grpc.Server
}

func NewServerHTTP(cartService *service.CartService)*ServerHTTP{
	engine:=grpc.NewServer()

	pb.RegisterCartServiceServer(engine,cartService)
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