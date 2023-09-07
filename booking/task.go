package booking

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"go.opentelemetry.io/otel/trace"
)

type TaskContext struct {
	SpanID  string
	TraceID string
}

type TaskPayload struct {
	Context TaskContext
	Data    interface{}
}

const (
	TypeReservation = "booking:reserve"
)

func NewReservationTask(ctx context.Context, order *Order) (*asynq.Task, error) {
	span := trace.SpanFromContext(ctx)
	defer span.End()

	data := TaskPayload{
		Context: TaskContext{
			SpanID:  span.SpanContext().SpanID().String(),
			TraceID: span.SpanContext().TraceID().String(),
		},
		Data: order,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeReservation, payload), nil
}
