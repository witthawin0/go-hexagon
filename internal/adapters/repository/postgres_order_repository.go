package repository

import (
	"database/sql"

	"github.com/witthawin0/go-hexagon/internal/domain"
	"github.com/witthawin0/go-hexagon/internal/ports"
)

type PostgresOrderRepository struct {
	db *sql.DB
}

func NewPostgresOrderRepository(db *sql.DB) ports.OrderRepository {
	return &PostgresOrderRepository{db: db}
}

func (r *PostgresOrderRepository) Save(order *domain.Order) error {
	query := `INSERT INTO orders (id, customer_id, total_amount, status) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, order.ID, order.CustomerID, order.TotalAmount, order.Status)
	return err
}

func (r *PostgresOrderRepository) Update(id string, order *domain.Order) error {
	query := `UPDATE orders SET customer_id=$2, total_amount=$3, status=$4 WHERE id=$1`
	_, err := r.db.Exec(query, order.ID, order.CustomerID, order.TotalAmount, order.Status)
	return err
}

func (r *PostgresOrderRepository) Delete(id string) error {
	query := `DELETE FROM orders WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PostgresOrderRepository) FindByID(id string) (*domain.Order, error) {
	query := `SELECT id, customer_id, total_amount, status FROM orders WHERE id=$1`
	row := r.db.QueryRow(query, id)

	var order domain.Order
	err := row.Scan(&order.ID, &order.CustomerID, &order.TotalAmount, &order.Status)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *PostgresOrderRepository) FindAll() ([]*domain.Order, error) {
	query := `SELECT id, customer_id, total_amount, status FROM orders`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*domain.Order
	for rows.Next() {
		var order domain.Order
		if err := rows.Scan(&order.ID, &order.CustomerID, &order.TotalAmount, &order.Status); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	return orders, nil
}
