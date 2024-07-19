package usecase

import (
	"context"
	"orderService/internal/domain"
	"orderService/internal/repository"
	"orderService/internal/repository/rabbitmq"
)

type OrderUseCaseImpl struct {
	orderRepo repository.OrderRepository
	rabbitMq  rabbitmq.RabbitMQClient
}

func (o *OrderUseCaseImpl) CreateOrder(ctx context.Context, order *domain.Order) error {
	//Insert Into Table Order
	if err := o.orderRepo.InsertOrder(ctx, order); err != nil {
		return err
	}

	if err := o.rabbitMq.SendMessage(order); err != nil {
		return err
	}

	return nil
}

func NewOrderUseCase(orderRepo repository.OrderRepository, rabbitMq rabbitmq.RabbitMQClient) OrderUseCase {
	return &OrderUseCaseImpl{
		orderRepo: orderRepo,
		rabbitMq:  rabbitMq,
	}

}
