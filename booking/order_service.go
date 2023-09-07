package booking

import (
	"context"
	"github.com/iniudin/demo-ticket-booking/internal/database"
	"time"

	"github.com/hibiken/asynq"
	"github.com/rs/xid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type OrderService interface {
	Book(ctx context.Context, request CreateOrderRequest) (*Order, error)
	BookV2(ctx context.Context, request CreateOrderRequest) (*Order, error)
	FindByID(ctx context.Context, orderID string) (*Order, error)
}

type OrderServiceImpl struct {
	Order      OrderStore
	Concert    ConcertStore
	dispatcher *asynq.Client
}

func (s *OrderServiceImpl) Book(ctx context.Context, request CreateOrderRequest) (*Order, error) {
	_, span := orderTracer.Start(ctx, "Book: reserve Order")
	defer span.End()

	tx := database.Begin(ctx)
	defer database.Commit(ctx, tx)

	concertID, err := xid.FromString(request.ConcertID)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	err = s.Concert.Reserve(ctx, tx, concertID)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	order, err := s.Order.Save(ctx, tx, Order{
		ID:        xid.New(),
		Email:     request.Email,
		ConcertID: concertID,
		Status:    Reserved,
		Date:      time.Now().UTC(),
	})
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	return order, nil
}

func (s *OrderServiceImpl) BookV2(ctx context.Context, request CreateOrderRequest) (*Order, error) {
	_, span := orderTracer.Start(ctx, "BookV2: reserve Order using worker")
	defer span.End()

	tx := database.Begin(ctx)
	defer database.Commit(ctx, tx)

	concertID, err := xid.FromString(request.ConcertID)
	if err != nil {
		return nil, err
	}
	order, err := s.Order.Save(ctx, tx, Order{
		ID:        xid.New(),
		Email:     request.Email,
		ConcertID: concertID,
		Status:    Created,
		Date:      time.Now().UTC(),
	})
	if err != nil {
		return nil, err
	}

	task, err := NewReservationTask(ctx, order)
	if err != nil {
		return nil, err
	}

	taskInfo, err := s.dispatcher.EnqueueContext(ctx, task, asynq.Queue("critical"))
	if err != nil {
		return nil, err
	}
	span.AddEvent("task is created", trace.WithAttributes(attribute.String("task_info_id", taskInfo.ID)))

	return order, nil
}

func (s *OrderServiceImpl) FindByID(ctx context.Context, orderID string) (*Order, error) {
	_, span := orderTracer.Start(ctx, "find Order by id")
	defer span.End()

	tx := database.Begin(ctx)
	defer database.Commit(ctx, tx)

	id, err := xid.FromString(orderID)
	if err != nil {
		return nil, err
	}
	order, err := s.Order.FindByID(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func NewOrderService(order OrderStore, concert ConcertStore, client *asynq.Client) OrderService {
	return &OrderServiceImpl{Order: order, Concert: concert, dispatcher: client}
}
