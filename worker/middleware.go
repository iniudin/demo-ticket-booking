package worker

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/iniudin/demo-ticket-booking/booking"
	"github.com/iniudin/demo-ticket-booking/internal/telemetry"
	"go.opentelemetry.io/otel/trace"
)

func SpanPropagator(next asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		var tp booking.TaskPayload
		err := json.Unmarshal(t.Payload(), &tp)
		if err != nil {
			return err
		}
		spanCtx, err := telemetry.NewSpanContext(telemetry.SpanContext{
			TraceID: tp.Context.TraceID,
			SpanID:  tp.Context.SpanID,
		})
		if err != nil {
			return err
		}
		ctx = trace.ContextWithSpanContext(ctx, spanCtx)
		payload, err := json.Marshal(tp.Data)
		if err != nil {
			return err
		}
		return next.ProcessTask(ctx, asynq.NewTask(t.Type(), payload))
	})
}
