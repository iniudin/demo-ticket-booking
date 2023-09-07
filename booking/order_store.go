package booking

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/rs/xid"
)

type OrderStore interface {
	Save(ctx context.Context, tx pgx.Tx, order Order) (*Order, error)
	FindByID(ctx context.Context, tx pgx.Tx, id xid.ID) (*Order, error)
}

type OrderStoreImpl struct{}

func NewOrderStore() OrderStore {
	return &OrderStoreImpl{}
}

func (s *OrderStoreImpl) Save(ctx context.Context, tx pgx.Tx, order Order) (*Order, error) {
	q := `INSERT INTO orders (id, email, status, date, concert_id) VALUES ($1, $2, $3, $4, $5) ON CONFLICT(id) DO UPDATE SET status =$3`
	if _, err := tx.Exec(ctx, q, order.ID, order.Email, order.Status, order.Date, order.ConcertID); err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *OrderStoreImpl) FindByID(ctx context.Context, tx pgx.Tx, id xid.ID) (*Order, error) {
	q := `SELECT id, email, status, date, concert_id FROM orders WHERE id = $1`
	var order Order
	err := tx.QueryRow(ctx, q, id).Scan(&order.ID, &order.Email, &order.Status, &order.Date, &order.ConcertID)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
