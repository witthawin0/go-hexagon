// internal/ports/mocks/product_repository_mock.go
package mocks

import (
	"github.com/witthawin0/go-hexagon/internal/domain"

	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Save(product *domain.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Update(id string, product *domain.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockProductRepository) FindByID(id string) (*domain.Product, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockProductRepository) FindAll() ([]*domain.Product, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*domain.Product), args.Error(1)
	}
	return nil, args.Error(1)
}
