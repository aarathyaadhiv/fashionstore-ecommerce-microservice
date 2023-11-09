package di

// import (
// 	"github.com/aarathyaadhiv/fashionstore-product-service/pkg/api"
// 	"github.com/aarathyaadhiv/fashionstore-product-service/pkg/api/service"
// 	"github.com/aarathyaadhiv/fashionstore-product-service/pkg/config"
// 	"github.com/aarathyaadhiv/fashionstore-product-service/pkg/db"
// 	"github.com/aarathyaadhiv/fashionstore-product-service/pkg/repository"
// 	"github.com/google/wire"
// )



// func InitializeAPI(c config.Config)(*api.ServerHTTP,error){
// 	wire.Build(db.ConnectDatabase,
// 	repository.NewProductRepository,
// service.NewProductService,
// api.NewServerHTTP)

// 		return &api.ServerHTTP{},nil
// }