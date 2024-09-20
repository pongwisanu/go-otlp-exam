package product

import (
	"context"
	"go-otlp/repositories/product"
	"go-otlp/services/price"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
)

const name = "go.otlp.api.product"

var (
	tracer = otel.Tracer(name)
	// meter  = otel.Meter(name)
	logger = otelslog.NewLogger(name)
)

type productService struct {
	productRepo product.ProductRepository
	priceServ   price.PriceService
}

func NewProductService(productRepo product.ProductRepository, priceServ price.PriceService) ProductService {
	return productService{
		productRepo: productRepo,
		priceServ:   priceServ,
	}
}

func (s productService) GetProducts(c context.Context) ([]ProductResponse, error) {
	ctx, span := tracer.Start(c, "Service.GetProducts")
	defer span.End()
	var msg string
	products, err := s.productRepo.GetProducts(c)

	if err != nil {
		msg = err.Error()
		logger.ErrorContext(ctx, msg, "error")
		return nil, err
	}

	productResponses := []ProductResponse{}

	for _, product := range products {
		price, err := s.priceServ.GetPrice(product.Id)

		if err != nil {
			msg = err.Error()
			logger.ErrorContext(ctx, msg, "error")
			return nil, err
		}

		productResponse := ProductResponse{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       price.Value,
		}
		productResponses = append(productResponses, productResponse)
	}
	msg = "Successful"
	logger.InfoContext(ctx, msg, "success")

	return productResponses, nil
}

func (s productService) GetProduct(id int) (*ProductResponse, error) {
	product, err := s.productRepo.GetProduct(id)
	if err != nil {
		return nil, err
	}

	price, err := s.priceServ.GetPrice(product.Id)
	if err != nil {
		return nil, err
	}

	productReponse := ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       price.Value,
	}

	return &productReponse, nil
}

func (s productService) AddProduct(request ProductRequest) (int, error) {
	product := product.Product{
		Name:        request.Name,
		Description: request.Description,
	}

	result, err := s.productRepo.AddProduct(product)

	if err != nil {
		return 0, err
	}

	return result, nil

}
