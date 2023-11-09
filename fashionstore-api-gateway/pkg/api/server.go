package api

import (
	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/api/handler"
	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)


type ServerHTTP struct{
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler,adminHandler *handler.AdminHandler)*ServerHTTP{
	engine:=gin.New()

	engine.Use(gin.Logger())

	userRoutes:=engine.Group("/")
	userRoutes.POST("signup",userHandler.SignUp)
	userRoutes.POST("login",userHandler.Login)

	adminRoutes:=engine.Group("/admin")
	adminRoutes.POST("/login",adminHandler.Login)

	return &ServerHTTP{engine: engine}

}

func (sh *ServerHTTP)Start(c config.Config){
	sh.engine.Run(c.Port)
}