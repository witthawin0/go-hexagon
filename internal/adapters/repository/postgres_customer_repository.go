package repository

import (
	"database/sql"

	"github.com/witthawin0/go-hexagon/internal/domain"
	"github.com/witthawin0/go-hexagon/internal/ports"
)

type postgresCustomerRepository struct {
	db *sql.DB
}

func NewPosgresCustomerRepository(db *sql.DB) ports.CustomerRepository {
	return &postgresCustomerRepository{db: db}
}

func (r *postgresCustomerRepository) Save(customer *domain.Customer) error {
	query := `INSERT INTO customers (id, name, email, address, phone) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, customer.ID, customer.Name, customer.Email, customer.Address, customer.Phone)
	return err
}

func (r *postgresCustomerRepository) Update(id string, customer *domain.Customer) error {
	query := `UPDATE customers SET name=$2, email=$3, address=$4, phone=$5 WHERE id=$1`
	_, err := r.db.Exec(query, customer.ID, customer.Name, customer.Email, customer.Address, customer.Phone)
	return err
}

func (r *postgresCustomerRepository) Delete(id string) error {
	query := `DELETE FROM customers WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *postgresCustomerRepository) FindByID(id string) (*domain.Customer, error) {
	query := `SELECT id, name, email, address, phone FROM customers WHERE id=$1`
	row := r.db.QueryRow(query, id)

	var customer domain.Customer
	err := row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Address, &customer.Phone)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *postgresCustomerRepository) FindAll() ([]*domain.Customer, error) {
	query := `SELECT id, name, email, address, phone FROM customers`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []*domain.Customer
	for rows.Next() {
		var customer domain.Customer
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Address, &customer.Phone); err != nil {
			return nil, err
		}
		customers = append(customers, &customer)
	}

	return customers, nil
}
