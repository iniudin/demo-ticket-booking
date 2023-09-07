package worker

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/hibiken/asynq"
	"github.com/iniudin/demo-ticket-booking/booking"
)

type handler struct {
	service orderService
}

func newHandler() *handler {
	service := &orderServiceImpl{order: booking.NewOrderStore(), concert: booking.NewConcertStore()}
	return &handler{service: service}
}
func (h handler) handleReservation(ctx context.Context, t *asynq.Task) error {
	ctx, span := orderTracer.Start(ctx, "handlerRReservation: reserve an Order")
	defer span.End()

	var order booking.Order
	err := json.Unmarshal(t.Payload(), &order)
	if err != nil {
		return err
	}
	_, err = h.service.Reserve(ctx, booking.ReservationRequest{
		ConcertID: order.ConcertID,
		OrderID:   order.ID,
	})
	if err != nil && !errors.Is(err, booking.ErrTicketIsNotAvailable) {
		return err
	}
	return nil
}
