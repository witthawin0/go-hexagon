package ports

import "github.com/witthawin0/go-hexagon/internal/domain"

type OrderService interface {
	CreateOrder(order *domain.Order) error
	UpdateOrder(id string, order *domain.Order) error
	DeleteOrder(id string) error
	GetOrderByID(id string) (*domain.Order, error)
	GetAllOrders() ([]*domain.Order, error)
}
