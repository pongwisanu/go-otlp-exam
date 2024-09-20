package product

import (
	"go-otlp/services/product"

	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const name = "go.otlp.api.product"

var (
	tracer = otel.Tracer(name)
	// meter  = otel.Meter(name)
	logger = otelslog.NewLogger(name)
)

type productHandler struct {
	productService product.ProductService
}

func NewProductHandler(productService product.ProductService) productHandler {
	return productHandler{productService: productService}
}

func (h productHandler) GetProducts(c *fiber.Ctx) error {
	ctx, span := tracer.Start(c.Context(), "Handler.GetProducts")
	defer span.End()

	ctx = trace.ContextWithSpan(ctx, span)

	products, err := h.productService.GetProducts(ctx)

	var msg string
	if err != nil {
		msg = err.Error()
		logger.ErrorContext(ctx, msg, "error")
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error",
			"error":   err.Error(),
		})
	}
	msg = "Successful"
	logger.InfoContext(ctx, msg, "info")

	return c.JSON(products)
}

func (h productHandler) GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": "Error",
			"error":   err.Error(),
		})
	}

	product, err := h.productService.GetProduct(id)

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error",
			"error":   err.Error(),
		})
	}

	return c.JSON(product)
}
