// internal/ports/mocks/order_repository_mock.go
package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/witthawin0/go-hexagon/internal/domain"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Save(order *domain.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepository) Update(id string, order *domain.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockOrderRepository) FindByID(id string) (*domain.Order, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Order), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockOrderRepository) FindAll() ([]*domain.Order, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*domain.Order), args.Error(1)
	}
	return nil, args.Error(1)
}
