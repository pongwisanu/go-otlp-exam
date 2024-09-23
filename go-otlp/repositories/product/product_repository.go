package product

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type productRepositoryDb struct {
	db *gorm.DB
}

const name = "go.otlp.api.product"

var (
	tracer = otel.Tracer(name)
	// meter  = otel.Meter(name)
	// logger = otelslog.NewLogger(name)
)

func NewProductRepositoryDb(db *gorm.DB) ProductRepository {
	return productRepositoryDb{db: db}
}

func (r productRepositoryDb) GetProducts(c context.Context) ([]Product, error) {
	_, span := tracer.Start(c, "Repository.Product")
	defer span.End()

	span.SetAttributes(
		attribute.String("http.repository", "repositories.GetProducts"),
	)

	products := []Product{}
	result := r.db.Find(&products)

	// var msg string
	if result.Error != nil {
		span.RecordError(result.Error)
		// msg = result.Error.Error()
		// logger.ErrorContext(c, msg,
		// 	"method", "repository",
		// 	"status", "error",
		// )
		return nil, result.Error
	}
	// msg = "Successful"
	// logger.InfoContext(c, msg,
	// 	"method", "repository",
	// 	"status", "success",
	// )
	span.AddEvent("Repositories.GetProducts", trace.WithAttributes(
		attribute.String("status", "success"),
	))

	return products, nil
}

func (r productRepositoryDb) GetProduct(id int) (*Product, error) {
	product := Product{}
	result := r.db.First(&product, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r productRepositoryDb) AddProduct(newProduct Product) (int, error) {

	result := r.db.Create(&newProduct)

	if result.Error != nil {
		return 0, result.Error
	}

	return newProduct.Id, nil
}
