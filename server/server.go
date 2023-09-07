package server

import (
	"context"
	"errors"
	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/iniudin/demo-ticket-booking/booking"
	"github.com/iniudin/demo-ticket-booking/internal/database"
	"github.com/iniudin/demo-ticket-booking/internal/telemetry"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	App *fiber.App
}

func NewServer() *Server {
	app := fiber.New(fiber.Config{IdleTimeout: time.Second * 10})

	app.Use(otelfiber.Middleware())
	app.Use(recover.New())

	booking.NewBookingRouter(app)
	return &Server{App: app}
}

func (s *Server) Run() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	database.SetupDatabase(ctx)
	traceShutdown, err := telemetry.NewTracer(ctx, os.Getenv("OTEL_RECEIVER_ENDPOINT"), "booking-api")
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := traceShutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider: %w", err)
		}
	}()

	metricsShutdown, err := telemetry.NewMetrics(ctx, os.Getenv("OTEL_RECEIVER_ENDPOINT"), "booking-api")
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := metricsShutdown(ctx); err != nil {
			log.Fatal("failed to shutdown MetricsProvider: %w", err)
		}
	}()
	go func() {
		if err := s.App.Listen(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	defer cancel()

	if err := s.App.ShutdownWithContext(ctx); err != nil {
		log.Fatal(err)
	}
}
