package api

import (
	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/api/handler"
	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/api/middleware"
	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)


type ServerHTTP struct{
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler,adminHandler *handler.AdminHandler,productHandler *handler.ProductHandler,cartHandler *handler.CartHandler)*ServerHTTP{
	engine:=gin.New()

	engine.Use(gin.Logger())

	userRoutes:=engine.Group("/")
	userRoutes.POST("signup",userHandler.SignUp)
	userRoutes.POST("login",userHandler.Login)

	adminRoutes:=engine.Group("/admin")
	adminRoutes.POST("/login",adminHandler.Login)

	productRoutes:=engine.Group("/product")
	productRoutes.GET("",productHandler.ListProducts)
	productRoutes.Use(middleware.AdminAuthorizationMiddleware)
	productRoutes.POST("",productHandler.AddProduct)

	cartRoutes:=engine.Group("/cart")
	cartRoutes.Use(middleware.UserAuthorizationMiddleware)
	cartRoutes.POST("/:id",cartHandler.AddToCart)
	cartRoutes.GET("",cartHandler.ShowProductInCart)
	

	return &ServerHTTP{engine: engine}

}

func (sh *ServerHTTP)Start(c config.Config){
	sh.engine.Run(c.Port)
}