package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/client"
	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/pb"
	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type CartHandler struct{
	client *client.CartClient
}

func NewCartHandler(client *client.CartClient)*CartHandler{
	return &CartHandler{client: client}
}

func (cr *CartHandler) AddToCart(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.Responses(http.StatusBadRequest, "userid not retrieved", nil, errors.New("userid retrieval error").Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	productId:=c.Param("id")
	product,err:=strconv.Atoi(productId)
	if err!=nil{
		errRes:=response.Responses(http.StatusBadRequest,"query param recovery error",nil,errors.New("error in fetching path params").Error())
		c.JSON(http.StatusBadRequest,errRes)
		return
	}
	
	
	_,err=cr.client.Client.AddToCart(context.Background(),&pb.AddToCartRequest{CartId: id.(uint64),ProductId: uint64(product)})
	if err!=nil{
		errRes:=response.Responses(http.StatusInternalServerError,"internal server error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.Responses(http.StatusOK,"successfully added to cart",nil,nil)
	c.JSON(http.StatusOK,succRes)
	
}

func (cr *CartHandler) ShowProductInCart(c *gin.Context){
	id,ok:=c.Get("userId")
	if !ok {
		errRes := response.Responses(http.StatusBadRequest, "userid not retrieved", nil, errors.New("userid retrieval error").Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	page:=c.DefaultQuery("page","1")
	count:=c.DefaultQuery("count","3")
	products,err:=cr.client.Client.ShowCart(context.Background(),&pb.ShowCartRequest{CartId: id.(uint64),Page: page,Count: count})
	if err!=nil{
		errRes:=response.Responses(http.StatusInternalServerError,"internal server error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.Responses(http.StatusOK,"successfully showing products in cart",products,nil)
	c.JSON(http.StatusOK,succRes)
}
