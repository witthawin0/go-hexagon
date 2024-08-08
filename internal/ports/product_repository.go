package ports

import "github.com/witthawin0/go-hexagon/internal/domain"

type ProductRepository interface {
	Save(product *domain.Product) error
	Update(id string, product *domain.Product) error
	Delete(id string) error
	FindByID(id string) (*domain.Product, error)
	FindAll() ([]*domain.Product, error)
}
