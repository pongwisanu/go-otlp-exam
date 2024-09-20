package routes

import (
	product_hand "go-otlp/handlers/product"
	price_repo "go-otlp/repositories/price"
	product_repo "go-otlp/repositories/product"
	price_serv "go-otlp/services/price"
	product_service "go-otlp/services/product"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Register(dsn string, api *fiber.App) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Fail to connect database")
	}

	price_repository := price_repo.NewPriceRepositoryDb(db)
	price_service := price_serv.NewPriceService(price_repository)

	product_repository := product_repo.NewProductRepositoryDb(db)
	product_service := product_service.NewProductService(product_repository, price_service)
	product_handler := product_hand.NewProductHandler(product_service)

	api.Use(logger.New())

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Go otlp example")
	})

	api.Get("/products", product_handler.GetProducts)
	api.Get("/products/:id", product_handler.GetProduct)
}
