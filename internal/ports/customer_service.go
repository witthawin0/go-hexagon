package ports

import "github.com/witthawin0/go-hexagon/internal/domain"

type CustomerService interface {
	CreateCustomer(customer *domain.Customer) error
	UpdateCustomer(id string, customer *domain.Customer) error
	DeleteCustomer(id string) error
	GetCustomerByID(id string) (*domain.Customer, error)
	GetAllCustomers() ([]*domain.Customer, error)
}
