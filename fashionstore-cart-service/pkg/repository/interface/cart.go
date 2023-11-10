package interfaces

import "github.com/aarathyaadhiv/fashionstore-cart-service/pkg/models"


type CartRepository interface{
	AddToCart(CartID, productId uint64, quantity uint32, amount float64) error 
	QuantityOfProductInCart(cartId, productId uint64) (uint32, error)
	AmountOfProductInCart(cartId, productId uint64) (float64, error)
	UpdateCart(CartID, ProductId uint64, quantity uint32, amount float64) error
	ProductsInCart(cartId uint64,page,count int) ([]uint64, error)
	FetchQuantityAndAmount(cartId, productId uint64) (models.CartProductDetail, error)
}