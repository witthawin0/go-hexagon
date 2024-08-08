package ports

import "github.com/witthawin0/go-hexagon/internal/domain"

type OrderReposistory interface {
	Save(order *domain.Order) error
	Update(id string, order *domain.Order) error
	Delete(id string) error
	FindByID(id string) (domain.Order, error)
	FindAll() ([]*domain.Order, error)
}
