package booking

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ConcertHandler interface {
	Create(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
}

type ConcertHandlerImpl struct {
	concert ConcertService
}

func (h *ConcertHandlerImpl) Create(ctx *fiber.Ctx) error {
	var request CreateConcertRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	response, err := h.concert.Create(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"data": response})
}

func (h *ConcertHandlerImpl) FindByID(ctx *fiber.Ctx) error {
	concertID := ctx.Params("concertID")
	response, err := h.concert.FindByID(ctx.UserContext(), concertID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"data": response})
}

func (h *ConcertHandlerImpl) FindAll(ctx *fiber.Ctx) error {
	response, err := h.concert.FindAll(ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"data": response})
}

func NewConcertHandler(concert ConcertService) ConcertHandler {
	return &ConcertHandlerImpl{concert: concert}
}
