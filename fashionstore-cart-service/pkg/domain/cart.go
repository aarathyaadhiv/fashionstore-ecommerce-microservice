package domain

type Cart struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	CartID uint `json:"cart_id"`
	ProductID uint `json:"product_id"`
	Quantity uint    `json:"quantity" `
	Amount   float64 `json:"amount"`
}
