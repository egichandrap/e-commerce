package domain

import (
	"fmt"
	"orderService/internal/delivery/http/request"
	"time"
)

type Order struct {
	OrderID     string    `json:"order_id"`
	CustomerID  string    `json:"customer_id"`
	OrderDate   time.Time `json:"order_date"`
	Status      string    `json:"status"`
	TotalAmount float64   `json:"total_amount"`
}

func New(order *request.Request) *Order {
	timestamp := time.Now().Format("20060102150405")
	customerIDLen := len(order.CustomerID)
	lastFiveDigits := order.CustomerID
	if customerIDLen > 5 {
		lastFiveDigits = order.CustomerID[customerIDLen-5:]
	}
	orderID := fmt.Sprintf("OD%s%s", timestamp, lastFiveDigits)

	return &Order{
		OrderID:     orderID,
		CustomerID:  order.CustomerID,
		OrderDate:   time.Now(),
		Status:      order.Status,
		TotalAmount: order.Amount,
	}
}
