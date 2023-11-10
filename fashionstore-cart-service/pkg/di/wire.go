package di

// import (
// 	"github.com/aarathyaadhiv/fashionstore-cart-service/pkg/api"
// 	"github.com/aarathyaadhiv/fashionstore-cart-service/pkg/api/service"
// 	"github.com/aarathyaadhiv/fashionstore-cart-service/pkg/client"
// 	"github.com/aarathyaadhiv/fashionstore-cart-service/pkg/config"
// 	"github.com/aarathyaadhiv/fashionstore-cart-service/pkg/db"
// 	"github.com/aarathyaadhiv/fashionstore-cart-service/pkg/repository"
// 	"github.com/google/wire"
// )


// func InitializeAPI(c config.Config)(*api.ServerHTTP,error){
// 	wire.Build(db.ConnectDatabase,
// 	repository.NewCartRepository,
// 	client.NewProductClient,
// service.NewCartService,
// api.NewServerHTTP)
//     return &api.ServerHTTP{},nil
// }