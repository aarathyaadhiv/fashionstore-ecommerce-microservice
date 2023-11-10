package service

import (
	"context"
	"errors"
	
	"strconv"

	"github.com/aarathyaadhiv/fashionstore-cart-service/pkg/client"
	"github.com/aarathyaadhiv/fashionstore-cart-service/pkg/pb"
	interfaces "github.com/aarathyaadhiv/fashionstore-cart-service/pkg/repository/interface"
)


type CartService struct{
	Repo interfaces.CartRepository
	Product *client.ProductClient
	pb.UnimplementedCartServiceServer
}

func NewCartService(repo interfaces.CartRepository,product *client.ProductClient)*CartService{
	return &CartService{Repo: repo,Product: product}
}

func (c *CartService) AddToCart(ctx context.Context,req *pb.AddToCartRequest) (*pb.AddToCartResponse,error) {
	quantity, err := c.Repo.QuantityOfProductInCart(req.CartId, req.ProductId)
	if err != nil {
		return &pb.AddToCartResponse{},err
	}
	product, err := c.Product.Client.FetchProduct(context.Background(),&pb.FetchProductRequest{Id: req.ProductId})
	if err != nil {
		return &pb.AddToCartResponse{},err
	}
	productQuantity := product.Quantity - quantity
	if productQuantity > 0 {
		if quantity == 0 {
			err= c.Repo.AddToCart(req.CartId, req.ProductId, 1, product.SellingPrice)
			return &pb.AddToCartResponse{},err
		}
		amount, err := c.Repo.AmountOfProductInCart(req.CartId, req.ProductId)
		if err != nil {
			return &pb.AddToCartResponse{},err
		}
		err= c.Repo.UpdateCart(req.CartId, req.ProductId, quantity+1, amount+product.SellingPrice)
		 return &pb.AddToCartResponse{},err
	}
	return &pb.AddToCartResponse{},errors.New("product out of stock")

}

func (c *CartService) ShowCart(ctx context.Context,req *pb.ShowCartRequest) (*pb.ShowCartResponse, error) {
	page,err:=strconv.Atoi(req.Page)
	if err!=nil{
		return nil,err
	}
	count,err:=strconv.Atoi(req.Count)
	if err!=nil{
		return nil,err
	}
	
	ProductId, err := c.Repo.ProductsInCart(req.CartId,page,count)
	if err != nil {
		return nil, err
	}
	
	products,err:=c.Product.Client.ProductDetails(context.Background(),&pb.ProductDetailsRequest{Id: ProductId})
	
	if err!=nil{
		return nil,err
	}
	updatedCartProduct:=make([]*pb.CartProduct,0)
	var cartProduct *pb.CartProduct
	for _,product:=range products.Products{
		details,_:=c.Repo.FetchQuantityAndAmount(req.CartId,product.Id)
		
		cartProduct = &pb.CartProduct{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Amount:      details.Amount,
			Quantity:    details.Quantity,
		}
		if product.Quantity==0{
			cartProduct.Status="out of stock"
		}else if product.Quantity==1{
			cartProduct.Status="only 1 product remains"
		}else{
			cartProduct.Status="in stock"
		}
		
		updatedCartProduct = append(updatedCartProduct, cartProduct)
	}
	
	return &pb.ShowCartResponse{Product: updatedCartProduct}, nil
}