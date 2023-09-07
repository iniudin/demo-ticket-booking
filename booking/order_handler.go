package booking

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type OrderHandler interface {
	BookV1(ctx *fiber.Ctx) error
	BookV2(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
}

type OrderHandlerImpl struct {
	order OrderService
}

func NewOrderHandler(order OrderService) OrderHandler {
	return &OrderHandlerImpl{order: order}
}

func (h *OrderHandlerImpl) BookV1(ctx *fiber.Ctx) error {
	var request CreateOrderRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	response, err := h.order.Book(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"data": response})

}
func (h *OrderHandlerImpl) BookV2(ctx *fiber.Ctx) error {
	var request CreateOrderRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	response, err := h.order.BookV2(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"data": response})

}
func (h *OrderHandlerImpl) FindByID(ctx *fiber.Ctx) error {
	orderID := ctx.Params("orderID")
	response, err := h.order.FindByID(ctx.UserContext(), orderID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"data": response})
}
