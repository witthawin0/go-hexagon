package adapters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/witthawin0/go-hexagon/core"
)

type httpOrderHandler struct {
	service core.OrderService
}

func NewHttpOrderHanlder(service core.OrderService) *httpOrderHandler {
	return &httpOrderHandler{service: service}
}

func (h *httpOrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order core.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}

	if err := h.service.CreateOrder(order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}
