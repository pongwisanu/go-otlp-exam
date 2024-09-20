package product

import (
	"context"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

type productRepositoryDb struct {
	db *gorm.DB
}

const name = "go.otlp.api.product"

var (
	tracer = otel.Tracer(name)
	// meter  = otel.Meter(name)
	logger = otelslog.NewLogger(name)
)

func NewProductRepositoryDb(db *gorm.DB) ProductRepository {
	return productRepositoryDb{db: db}
}

func (r productRepositoryDb) GetProducts(ctx context.Context) ([]Product, error) {
	c, span := tracer.Start(ctx, "Repository.Product")
	defer span.End()

	products := []Product{}
	result := r.db.Find(&products)

	var msg string
	if result.Error != nil {
		msg = result.Error.Error()
		logger.ErrorContext(c, msg, "Error")
		return nil, result.Error
	}
	msg = "Successful"
	logger.InfoContext(c, msg, "Successful")

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
