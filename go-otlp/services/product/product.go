package product

import "context"

type ProductResponse struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductService interface {
	GetProducts(context.Context) ([]ProductResponse, error)
	GetProduct(int) (*ProductResponse, error)
	AddProduct(ProductRequest) (int, error)
}
