package ports

import "github.com/witthawin0/go-hexagon/internal/domain"

type CustomerRepository interface {
	Save(customer *domain.Customer) error
	Update(id string, customer *domain.Customer) error
	Delete(id string) error
	FindByID(id string) (*domain.Customer, error)
	FindAll() ([]*domain.Customer, error)
}
