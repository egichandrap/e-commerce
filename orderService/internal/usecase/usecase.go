package usecase

import (
	"context"
	"orderService/internal/domain"
)

type OrderUseCase interface {
	CreateOrder(ctx context.Context, order *domain.Order) error
}
