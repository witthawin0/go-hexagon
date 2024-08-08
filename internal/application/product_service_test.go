// internal/application/product_service_test.go
package application

import (
	"testing"

	"github.com/witthawin0/go-hexagon/internal/domain"
	"github.com/witthawin0/go-hexagon/internal/ports/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	mockRepo := new(mocks.MockProductRepository)
	productService := NewProductService(mockRepo)

	product := &domain.Product{
		ID:          "1",
		Name:        "Product1",
		Description: "Description1",
		Price:       100.0,
		Stock:       10,
	}

	mockRepo.On("Save", product).Return(nil)

	err := productService.CreateProduct(product)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateProduct(t *testing.T) {
	mockRepo := new(mocks.MockProductRepository)
	productService := NewProductService(mockRepo)

	product := &domain.Product{
		ID:          "1",
		Name:        "Product1",
		Description: "Description1",
		Price:       100.0,
		Stock:       10,
	}

	mockRepo.On("Update", product).Return(nil)

	err := productService.UpdateProduct("1", product)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteProduct(t *testing.T) {
	mockRepo := new(mocks.MockProductRepository)
	productService := NewProductService(mockRepo)

	productID := "1"

	mockRepo.On("Delete", productID).Return(nil)

	err := productService.DeleteProduct(productID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetProductByID(t *testing.T) {
	mockRepo := new(mocks.MockProductRepository)
	productService := NewProductService(mockRepo)

	product := &domain.Product{
		ID:          "1",
		Name:        "Product1",
		Description: "Description1",
		Price:       100.0,
		Stock:       10,
	}

	mockRepo.On("FindByID", "1").Return(product, nil)

	result, err := productService.GetProductByID("1")

	assert.NoError(t, err)
	assert.Equal(t, product, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAllProducts(t *testing.T) {
	mockRepo := new(mocks.MockProductRepository)
	productService := NewProductService(mockRepo)

	products := []*domain.Product{
		{
			ID:          "1",
			Name:        "Product1",
			Description: "Description1",
			Price:       100.0,
			Stock:       10,
		},
		{
			ID:          "2",
			Name:        "Product2",
			Description: "Description2",
			Price:       200.0,
			Stock:       20,
		},
	}

	mockRepo.On("FindAll").Return(products, nil)

	result, err := productService.GetAllProducts()

	assert.NoError(t, err)
	assert.Equal(t, products, result)
	mockRepo.AssertExpectations(t)
}
