// internal/ports/product_service.go
package ports

import "github.com/witthawin0/go-hexagon/internal/domain"

type ProductService interface {
	CreateProduct(product *domain.Product) error
	UpdateProduct(product *domain.Product) error
	DeleteProduct(id string) error
	GetProductByID(id string) (*domain.Product, error)
	GetAllProducts() ([]*domain.Product, error)
}
