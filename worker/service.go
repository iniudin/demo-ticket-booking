package worker

import (
	"context"
	"errors"
	"github.com/iniudin/demo-ticket-booking/booking"
	"github.com/iniudin/demo-ticket-booking/internal/database"
)

type orderServiceImpl struct {
	order   booking.OrderStore
	concert booking.ConcertStore
}

type orderService interface {
	Reserve(ctx context.Context, request booking.ReservationRequest) (*booking.Order, error)
}

func (s *orderServiceImpl) Reserve(ctx context.Context, request booking.ReservationRequest) (*booking.Order, error) {
	ctx, span := orderTracer.Start(ctx, "Reserve: reserve an Order")
	defer span.End()

	tx := database.Begin(ctx)
	defer database.Commit(ctx, tx)

	err := s.concert.Reserve(ctx, tx, request.ConcertID)
	if err != nil && !errors.Is(err, booking.ErrTicketIsNotAvailable) {
		span.RecordError(err)
		return nil, err
	}

	if errors.Is(err, booking.ErrTicketIsNotAvailable) {
		span.AddEvent("order rejected")
		if _, err := s.order.Save(ctx, tx, booking.Order{ID: request.OrderID, ConcertID: request.ConcertID, Status: booking.Rejected}); err != nil {
			span.RecordError(err)
			return nil, err
		}
	}
	span.AddEvent("order success")
	order, err := s.order.Save(ctx, tx, booking.Order{ID: request.OrderID, ConcertID: request.ConcertID, Status: booking.Reserved})
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	return order, nil
}
