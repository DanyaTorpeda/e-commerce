package shoppingcart

import (
	cartitem "e-commerce/internal/domains/cart_item"
	"time"
)

type ShoppingCart struct {
	ID         int                 `json:"id"`
	UserID     int                 `json:"user_id"`
	Items      []cartitem.CartItem `json:"items"`
	TotalPrice float64             `json:"total_price"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
}
