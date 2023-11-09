package di

// import (
// 	"github.com/aarathyaadhiv/fashionstore-admin-service/pkg/api"
// 	"github.com/aarathyaadhiv/fashionstore-admin-service/pkg/api/service"
// 	"github.com/aarathyaadhiv/fashionstore-admin-service/pkg/config"
// 	"github.com/aarathyaadhiv/fashionstore-admin-service/pkg/db"
// 	"github.com/aarathyaadhiv/fashionstore-admin-service/pkg/repository"
// 	"github.com/google/wire"
// )


// func InitializeAPI(c config.Config)(*api.ServerHTTP,error){
// 	wire.Build(db.ConnectDatabase,
// 	repository.NewAdminRepository,
// service.NewAdminService,
// api.NewServerHTTP)
// 		return &api.ServerHTTP{},nil
// }