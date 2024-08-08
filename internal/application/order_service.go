package application

import (
	"github.com/witthawin0/go-hexagon/internal/domain"
	"github.com/witthawin0/go-hexagon/internal/ports"
	"github.com/witthawin0/go-hexagon/internal/utils/logger"
)

type orderServiceImpl struct {
	repo ports.OrderRepository
}

func NewOrderService(repo ports.OrderRepository) ports.OrderService {
	return &orderServiceImpl{repo: repo}
}

func (s *orderServiceImpl) CreateOrder(order *domain.Order) error {
	logger.Info("Creating order", order)

	err := s.repo.Save(order)
	if err != nil {
		logger.Error("Error creating order:", err)
		return err
	}

	logger.Info("Order created successfully", order.ID)

	return nil
}

func (s *orderServiceImpl) UpdateOrder(id string, order *domain.Order) error {
	logger.Info("Updating order:", order)

	err := s.repo.Update(id, order)
	if err != nil {
		logger.Error("Error updating order:", err)
		return err
	}

	logger.Info("Order updated successfully:", order.ID)

	return nil
}

func (s *orderServiceImpl) DeleteOrder(id string) error {
	logger.Info("Deleting order:", id)

	err := s.repo.Delete(id)
	if err != nil {
		logger.Error("Error deleting order:", err)
		return err
	}

	logger.Info("Order deleted successfully:", id)

	return nil
}

func (s *orderServiceImpl) GetOrderByID(id string) (*domain.Order, error) {
	logger.Info("Retrieving order by ID:", id)

	order, err := s.repo.FindByID(id)
	if err != nil {
		logger.Error("Error retrieving order:", err)
		return nil, err
	}

	return order, nil
}

func (s *orderServiceImpl) GetAllOrders() ([]*domain.Order, error) {
	logger.Info("Retrieving all orders")

	orders, err := s.repo.FindAll()
	if err != nil {
		logger.Error("Error retrieving orders:", err)
		return nil, err
	}

	return orders, nil
}
