package main

import (
	"context"
	"errors"
	"fmt"
	"go-otlp/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_DB := os.Getenv("DB_DB")
	DB_PORT := os.Getenv("DB_PORT")
	DB_HOST := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_DB)

	PORT := os.Getenv("PORT")

	app := fiber.New()

	ctx := context.Background()

	otelShutdown, err := setupOTelSDK(ctx)
	if err != nil {
		return
	}

	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	routes.Register(dsn, app)

	log.Fatal(app.Listen(":" + PORT))
}
