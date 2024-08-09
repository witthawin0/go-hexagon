package application

import (
	"github.com/witthawin0/go-hexagon/internal/domain"
	"github.com/witthawin0/go-hexagon/internal/ports"
)

type customerService struct {
	repo ports.CustomerRepository
}

// NewCustomerService returns a new instance of CustomerService.
func NewCustomerService(repo ports.CustomerRepository) ports.CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) GetCustomerByID(id string) (*domain.Customer, error) {
	return s.repo.FindByID(id)
}

func (s *customerService) CreateCustomer(customer *domain.Customer) error {
	return s.repo.Save(customer)
}

func (s *customerService) UpdateCustomer(id string, customer *domain.Customer) error {
	return s.repo.Update(id, customer)
}

func (s *customerService) DeleteCustomer(id string) error {
	return s.repo.Delete(id)
}

func (s *customerService) GetAllCustomers() ([]*domain.Customer, error) {
	var customers []*domain.Customer
	customers, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	return customers, nil
}
