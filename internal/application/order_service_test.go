// internal/application/order_service_test.go
package application

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/witthawin0/go-hexagon/internal/domain"
	"github.com/witthawin0/go-hexagon/internal/ports/mocks"
)

func TestCreateOrder(t *testing.T) {
	mockRepo := new(mocks.MockOrderRepository)
	orderService := NewOrderService(mockRepo)

	order := &domain.Order{
		ID:          "1",
		CustomerID:  "123",
		OrderDate:   time.Now(),
		TotalAmount: 200.0,
		Products: []domain.OrderProduct{
			{ProductID: "1", Quantity: 2, Price: 100.0},
		},
	}

	mockRepo.On("Save", order).Return(nil)

	err := orderService.CreateOrder(order)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateOrder(t *testing.T) {
	mockRepo := new(mocks.MockOrderRepository)
	orderService := NewOrderService(mockRepo)

	order := &domain.Order{
		ID:          "1",
		CustomerID:  "123",
		OrderDate:   time.Now(),
		TotalAmount: 200.0,
		Products: []domain.OrderProduct{
			{ProductID: "1", Quantity: 2, Price: 100.0},
		},
	}

	mockRepo.On("Update", order).Return(nil)

	err := orderService.UpdateOrder("1", order)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteOrder(t *testing.T) {
	mockRepo := new(mocks.MockOrderRepository)
	orderService := NewOrderService(mockRepo)

	orderID := "1"

	mockRepo.On("Delete", orderID).Return(nil)

	err := orderService.DeleteOrder(orderID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetOrderByID(t *testing.T) {
	mockRepo := new(mocks.MockOrderRepository)
	orderService := NewOrderService(mockRepo)

	order := &domain.Order{
		ID:          "1",
		CustomerID:  "123",
		OrderDate:   time.Now(),
		TotalAmount: 200.0,
		Products: []domain.OrderProduct{
			{ProductID: "1", Quantity: 2, Price: 100.0},
		},
	}

	mockRepo.On("FindByID", "1").Return(order, nil)

	result, err := orderService.GetOrderByID("1")

	assert.NoError(t, err)
	assert.Equal(t, order, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAllOrders(t *testing.T) {
	mockRepo := new(mocks.MockOrderRepository)
	orderService := NewOrderService(mockRepo)

	orders := []*domain.Order{
		{
			ID:          "1",
			CustomerID:  "123",
			OrderDate:   time.Now(),
			TotalAmount: 200.0,
			Products: []domain.OrderProduct{
				{ProductID: "1", Quantity: 2, Price: 100.0},
			},
		},
		{
			ID:          "2",
			CustomerID:  "456",
			OrderDate:   time.Now(),
			TotalAmount: 400.0,
			Products: []domain.OrderProduct{
				{ProductID: "2", Quantity: 4, Price: 100.0},
			},
		},
	}

	mockRepo.On("FindAll").Return(orders, nil)

	result, err := orderService.GetAllOrders()

	assert.NoError(t, err)
	assert.Equal(t, orders, result)
	mockRepo.AssertExpectations(t)
}
