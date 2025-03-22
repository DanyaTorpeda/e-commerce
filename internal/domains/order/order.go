package order

import (
	orderitem "e-commerce/internal/domains/order_item"
	"time"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusCompleted OrderStatus = "completed"
)

type OrderPaymentsStatus string

const (
	PaymentsStatusPending OrderPaymentsStatus = "pending"
	PaymentsStatusPaid    OrderPaymentsStatus = "paid"
	PaymentsStatusFailed  OrderPaymentsStatus = "failed"
)

type Order struct {
	ID            int
	UserID        int
	Items         []orderitem.OrderItem
	TotalPrice    float64
	Status        OrderStatus
	PaymentStatus OrderPaymentsStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
