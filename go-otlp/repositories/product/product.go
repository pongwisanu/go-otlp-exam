package product

import "context"

type Product struct {
	Id          int `gorm:"primary"`
	Name        string
	Description string
}

func (c *Product) TableName() string {
	return "catalog.product"
}

type ProductRepository interface {
	GetProducts(context.Context) ([]Product, error)
	GetProduct(int) (*Product, error)
	AddProduct(Product) (int, error)
}
