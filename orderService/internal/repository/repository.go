package repository

import (
	"context"
	"orderService/internal/domain"
)

type OrderRepository interface {
	InsertOrder(ctx context.Context, order *domain.Order) error
}
