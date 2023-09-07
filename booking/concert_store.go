package booking

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/rs/xid"
)

type ConcertStore interface {
	Save(ctx context.Context, tx pgx.Tx, concert Concert) (*Concert, error)
	FindByID(ctx context.Context, tx pgx.Tx, id xid.ID) (*Concert, error)
	FindAll(ctx context.Context, tx pgx.Tx) ([]Concert, error)
	Reserve(ctx context.Context, tx pgx.Tx, id xid.ID) error
}

type ConcertStoreImpl struct{}

func NewConcertStore() ConcertStore {
	return &ConcertStoreImpl{}
}

func (s *ConcertStoreImpl) Save(ctx context.Context, tx pgx.Tx, concert Concert) (*Concert, error) {
	q := `INSERT INTO concerts(id, name, location, date, remaining) VALUES ($1, $2, $3, $4, $5) ON CONFLICT(id) DO UPDATE SET remaining =$5`
	if _, err := tx.Exec(ctx, q, concert.ID, concert.Name, concert.Location, concert.Date, concert.Remaining); err != nil {
		return nil, err
	}
	return &concert, nil
}

func (s *ConcertStoreImpl) Reserve(ctx context.Context, tx pgx.Tx, id xid.ID) error {
	q := `SELECT id, name, location, date, remaining FROM concerts WHERE id = $1 FOR NO KEY UPDATE`

	var concert Concert
	if err := tx.QueryRow(ctx, q, id).Scan(&concert.ID, &concert.Name, &concert.Location, &concert.Date, &concert.Remaining); err != nil {
		return err
	}

	if concert.Remaining-1 < 0 {
		return ErrTicketIsNotAvailable
	}
	concert.Remaining -= 1
	if _, err := s.Save(ctx, tx, concert); err != nil {
		return err
	}
	return nil
}

func (s *ConcertStoreImpl) FindByID(ctx context.Context, tx pgx.Tx, id xid.ID) (*Concert, error) {
	q := `SELECT id, name, location, date, remaining FROM concerts WHERE id = $1 LIMIT 1`
	var concert Concert
	if err := tx.QueryRow(ctx, q, id).Scan(&concert.ID, &concert.Name, &concert.Location, &concert.Date, &concert.Remaining); err != nil {
		return nil, err
	}

	return &concert, nil
}

func (s *ConcertStoreImpl) FindAll(ctx context.Context, tx pgx.Tx) ([]Concert, error) {
	q := `SELECT id, name, location, date, remaining FROM concerts`
	var concerts []Concert
	rows, err := tx.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var concert Concert
		if err := rows.Scan(&concert.ID, &concert.Name, &concert.Location, &concert.Date, &concert.Remaining); err != nil {
			return nil, err
		}
		concerts = append(concerts, concert)
	}
	return concerts, nil
}
