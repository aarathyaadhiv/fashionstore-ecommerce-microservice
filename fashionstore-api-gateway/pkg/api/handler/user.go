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


type UserHandler struct{
	Client *client.UserClient
}

func NewUserHandler(c *client.UserClient)*UserHandler{
	return &UserHandler{c}
}

func (ur *UserHandler) SignUp(c *gin.Context){
	var signUpDetails *pb.SignUpRequest
	if err := c.ShouldBindJSON(&signUpDetails); err != nil {
		erRes := response.Response{Statuscode: http.StatusBadRequest, Message: "fields are not provided in correct format", Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, erRes)
		return
	}
	if err := validator.New().Struct(signUpDetails); err != nil {
		erRes := response.Response{Statuscode: http.StatusBadRequest, Message: "constraints not satisfied", Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, erRes)
		return
	}

	userData, err := ur.Client.Client.SignUp(context.Background(),signUpDetails)
	if err != nil {
		erRes := response.Response{Statuscode: http.StatusInternalServerError, Message: "internal server error", Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, erRes)
		return
	}
	sucRes := response.Responses(http.StatusCreated, "successfully signedup", userData, nil)
	c.JSON(http.StatusCreated, sucRes)
}

func (ur *UserHandler) Login(c *gin.Context){
	var loginDetails *pb.LoginRequest
	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields are  provided in the bad format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := validator.New().Struct(loginDetails); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userData, err := ur.Client.Client.Login(context.Background(),loginDetails)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	sucRes := response.Responses(http.StatusOK, "successfully logged in", userData, nil)
	c.JSON(http.StatusOK, sucRes)
}