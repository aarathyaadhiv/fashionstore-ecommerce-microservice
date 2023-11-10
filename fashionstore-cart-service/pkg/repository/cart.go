package repository

import (
	"github.com/aarathyaadhiv/fashionstore-cart-service/pkg/models"
	interfaces "github.com/aarathyaadhiv/fashionstore-cart-service/pkg/repository/interface"
	"gorm.io/gorm"
)


type CartRepository struct{
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB)interfaces.CartRepository{
	return &CartRepository{db}
}

func (c *CartRepository) AddToCart(CartID, productId uint64, quantity uint32, amount float64) error {
	return c.DB.Exec(`INSERT INTO carts(cart_id,product_id,quantity,amount) VALUES(?,?,?,?)`, CartID, productId, quantity, amount).Error
}

func (c *CartRepository) QuantityOfProductInCart(cartId, productId uint64) (uint32, error) {
	var quantity uint32
	err := c.DB.Raw(`SELECT quantity FROM carts WHERE cart_id=? AND product_id=?`, cartId, productId).Scan(&quantity).Error
	return quantity, err
}

func (c *CartRepository) AmountOfProductInCart(cartId, productId uint64) (float64, error) {
	var amount float64
	err := c.DB.Raw(`SELECT amount FROM carts WHERE cart_id=? AND product_id=?`, cartId, productId).Scan(&amount).Error
	return amount, err
}

func (c *CartRepository) UpdateCart(CartID, ProductId uint64, quantity uint32, amount float64) error {
	return c.DB.Exec(`UPDATE carts SET quantity=?,amount=? WHERE cart_id=? AND product_id=?`, quantity, amount, CartID, ProductId).Error
}

func (c *CartRepository) ProductsInCart(cartId uint64,page,count int) ([]uint64, error) {
	var products []uint64
	offset:=(page-1)*count
	err := c.DB.Raw(`SELECT product_id FROM carts WHERE cart_id=? limit ? offset ?`, cartId,count,offset).Scan(&products).Error
	return products, err
}

func (c *CartRepository) FetchQuantityAndAmount(cartId, productId uint64) (models.CartProductDetail, error) {
	var details models.CartProductDetail
	err := c.DB.Raw(`SELECT quantity,amount FROM carts WHERE cart_id=? AND product_id=?`, cartId, productId).Scan(&details).Error
	return details, err
}
