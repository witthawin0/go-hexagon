// internal/application/product_service.go
package application

import (
	"github.com/witthawin0/go-hexagon/internal/domain"
	"github.com/witthawin0/go-hexagon/internal/ports"
	"github.com/witthawin0/go-hexagon/internal/utils/logger"
)

type ProductServiceImpl struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) ports.ProductService {

	return &ProductServiceImpl{repo: repo}
}

func (s *ProductServiceImpl) CreateProduct(product *domain.Product) error {
	logger.Info("Creating product", product)

	err := s.repo.Save(product)
	if err != nil {
		logger.Error("Error creating prduct:", err)
		return err
	}

	logger.Info("Product created successfully", product.ID)

	return nil
}

func (s *ProductServiceImpl) UpdateProduct(id string, product *domain.Product) error {
	logger.Info("Updating product:", product)

	err := s.repo.Update(id, product)
	if err != nil {
		logger.Error("Error updating product:", err)
		return err
	}

	logger.Info("Product updated successfully:", product.ID)

	return nil
}

func (s *ProductServiceImpl) DeleteProduct(id string) error {
	logger.Info("Retrieving product by ID:", id)

	err := s.repo.Delete(id)
	if err != nil {
		logger.Error("Error deleting product:", err)
		return err
	}

	logger.Info("Product deleted successfully:", id)

	return nil
}

func (s *ProductServiceImpl) GetProductByID(id string) (*domain.Product, error) {
	logger.Info("Retrieving product by ID:", id)

	product, err := s.repo.FindByID(id)
	if err != nil {
		logger.Error("Error retrieving product:", err)

		return nil, err
	}

	return product, nil
}

func (s *ProductServiceImpl) GetAllProducts() ([]*domain.Product, error) {
	logger.Info("Retrieving all products")

	products, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}
