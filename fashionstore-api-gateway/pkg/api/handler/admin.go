package handler

import (
	"context"
	"net/http"

	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/client"
	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/pb"
	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


type AdminHandler struct{
	Client *client.AdminClient
}

func NewAdminHandler(c *client.AdminClient)*AdminHandler{
	return &AdminHandler{c}
}

func (a *AdminHandler) Login(c *gin.Context){
	var adminLoginDetails *pb.AdminLoginRequest

	if err := c.ShouldBindJSON(&adminLoginDetails); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields are not in the correct order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := validator.New().Struct(adminLoginDetails); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "constraint not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	adminDetails, err := a.Client.Client.Login(context.Background(),adminLoginDetails)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	sucRes := response.Responses(http.StatusOK, "successfully logged in", adminDetails, nil)
	c.JSON(http.StatusOK, sucRes)
}