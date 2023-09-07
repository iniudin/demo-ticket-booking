package worker

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/iniudin/demo-ticket-booking/internal/database"
	"github.com/iniudin/demo-ticket-booking/internal/telemetry"
	"log"
	"os"
	"time"
)

type Worker struct {
	mux     *asynq.ServeMux
	server  *asynq.Server
	handler *handler
}

func NewWorker() *Worker {
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: os.Getenv("REDIS_HOST")},
		asynq.Config{
			Concurrency: 20,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			}},
	)

	worker := &Worker{
		server:  server,
		mux:     asynq.NewServeMux(),
		handler: newHandler(),
	}
	worker.mux.Use(SpanPropagator)
	worker.mux.HandleFunc("booking:reserve", worker.handler.handleReservation)
	return worker
}

func (w *Worker) Run() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	database.SetupDatabase(ctx)

	traceShutdown, err := telemetry.NewTracer(ctx, os.Getenv("OTEL_RECEIVER_ENDPOINT"), "booking-worker")
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := traceShutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider:", err)
		}
	}()

	metricsShutdown, err := telemetry.NewMetrics(ctx, os.Getenv("OTEL_RECEIVER_ENDPOINT"), "booking-worker")
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := metricsShutdown(ctx); err != nil {
			log.Fatal("failed to shutdown MetricsProvider:", err)
		}
	}()
	if err := w.server.Run(w.mux); err != nil {
		log.Println("Worker error:", err)
	}
}
