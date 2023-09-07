package booking

import (
	"context"
	"github.com/iniudin/demo-ticket-booking/internal/database"
	"github.com/rs/xid"
	"time"
)

type ConcertService interface {
	Create(ctx context.Context, request CreateConcertRequest) (*Concert, error)
	FindByID(ctx context.Context, concertID string) (*Concert, error)
	FindAll(ctx context.Context) ([]Concert, error)
}

type ConcertServiceImpl struct {
	concert ConcertStore
}

func NewConcertService(concert ConcertStore) ConcertService {
	return &ConcertServiceImpl{concert: concert}
}

func (s *ConcertServiceImpl) Create(ctx context.Context, request CreateConcertRequest) (*Concert, error) {
	_, span := concertTracer.Start(ctx, "creating new Concert")
	defer span.End()

	tx := database.Begin(ctx)
	defer database.Commit(ctx, tx)
	concert, err := s.concert.Save(ctx, tx, Concert{
		ID:        xid.New(),
		Name:      request.Name,
		Location:  request.Location,
		Date:      time.Now().UTC(),
		Remaining: request.Remaining,
	})
	if err != nil {
		return nil, err
	}

	return concert, nil
}

func (s *ConcertServiceImpl) FindByID(ctx context.Context, concertId string) (*Concert, error) {
	_, span := concertTracer.Start(ctx, "find Concert by id")
	defer span.End()

	tx := database.Begin(ctx)
	defer database.Commit(ctx, tx)

	id, err := xid.FromString(concertId)
	if err != nil {
		return nil, err
	}
	concert, err := s.concert.FindByID(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	return concert, nil
}

func (s *ConcertServiceImpl) FindAll(ctx context.Context) ([]Concert, error) {
	_, span := concertTracer.Start(ctx, "get all Concert")
	defer span.End()

	tx := database.Begin(ctx)
	defer database.Commit(ctx, tx)
	concert, err := s.concert.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}
	return concert, nil
}
