package repository

import (
	"context"
	"database/sql"
	"orderService/internal/domain"
)

type PsqlOrderRepository struct {
	db *sql.DB
}

func (p *PsqlOrderRepository) InsertOrder(ctx context.Context, order *domain.Order) error {
	query := `INSERT INTO orders(order_id, customer_id, order_date, status, total_amount) VALUES ($1,$2,$3,$4,$5)`
	_, err := p.db.ExecContext(ctx, query, order.OrderID, order.CustomerID, order.OrderDate, order.Status, order.TotalAmount)
	if err != nil {
		return err
	}
	return nil
}

func NewMysqlOrderRepository(db *sql.DB) OrderRepository {
	return &PsqlOrderRepository{db: db}

}
