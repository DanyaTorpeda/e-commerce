package payment

import "time"

type PaymentStatus string

const (
	StatusPending   PaymentStatus = "pending"
	StatusCompleted PaymentStatus = "completed"
	StatusFailed    PaymentStatus = "failed"
)

type PaymentMethod string

const (
	MethodPayPal PaymentMethod = "paypal"
	MethodStripe PaymentMethod = "stripe"
)

type Payment struct {
	ID            int
	OrderID       int
	Amount        float64
	Status        PaymentStatus
	PaymentMethod PaymentMethod
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
