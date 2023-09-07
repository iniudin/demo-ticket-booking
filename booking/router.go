package booking

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	"os"
)

func NewBookingRouter(app *fiber.App) {
	concertStore := NewConcertStore()
	concertService := NewConcertService(concertStore)
	concertHandler := NewConcertHandler(concertService)

	orderStore := NewOrderStore()
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: os.Getenv("REDIS_HOST")})
	orderService := NewOrderService(orderStore, concertStore, client)
	orderHandler := NewOrderHandler(orderService)

	v1Router := app.Group("/api/v1")

	v1Router.Post("/concerts", concertHandler.Create)
	v1Router.Get("/concerts", concertHandler.FindAll)
	v1Router.Get("/concerts/:concertID", concertHandler.FindByID)

	v1Router.Get("/orders/:orderID", orderHandler.FindByID)
	v1Router.Post("/orders", orderHandler.BookV1)

	// BookV2 - using worker
	v2Router := app.Group("/api/v2")
	v2Router.Post("/orders", orderHandler.BookV2)
}
