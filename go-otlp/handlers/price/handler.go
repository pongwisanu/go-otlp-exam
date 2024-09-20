package price

import (
	"go-otlp/services/price"

	"github.com/gofiber/fiber/v2"
)

type priceHandler struct {
	priceService price.PriceService
}

func NewPriceHandler(priceService price.PriceService) priceHandler {
	return priceHandler{priceService: priceService}
}

func (h priceHandler) GetPrices(c *fiber.Ctx) error {
	prices, err := h.priceService.GetPrices()

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error",
			"error":   err.Error(),
		})
	}

	return c.JSON(prices)
}

func (h priceHandler) GetPrice(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": "Error",
			"error":   err.Error(),
		})
	}

	price, err := h.priceService.GetPrice(id)

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error",
			"error":   err.Error(),
		})
	}

	return c.JSON(price)
}