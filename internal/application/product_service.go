// internal/application/product_service.go
package application

import (
	"github.com/witthawin0/go-hexagon/internal/domain"
	"github.com/witthawin0/go-hexagon/internal/ports"
)

type ProductServiceImpl struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) ports.ProductService {
	return &ProductServiceImpl{repo: repo}
}

func (s *ProductServiceImpl) CreateProduct(product *domain.Product) error {
	return s.repo.Save(product)
}

func (s *ProductServiceImpl) UpdateProduct(product *domain.Product) error {
	return s.repo.Update(product)
}

func (s *ProductServiceImpl) DeleteProduct(id string) error {
	return s.repo.Delete(id)
}

func (s *ProductServiceImpl) GetProductByID(id string) (*domain.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductServiceImpl) GetAllProducts() ([]*domain.Product, error) {
	return s.repo.FindAll()
}
