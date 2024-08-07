// internal/adapters/repository/postgres_product_repository.go
package repository

import (
	"database/sql"

	"github.com/witthawin0/go-hexagon/internal/domain"
	"github.com/witthawin0/go-hexagon/internal/ports"
)

type PostgresProductRepository struct {
	db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) ports.ProductRepository {
	return &PostgresProductRepository{db: db}
}

func (r *PostgresProductRepository) Save(product *domain.Product) error {
	_, err := r.db.Exec("INSERT INTO products (id, name, description, price, stock) VALUES ($1, $2, $3, $4, $5)",
		product.ID, product.Name, product.Description, product.Price, product.Stock)
	return err
}

func (r *PostgresProductRepository) Update(product *domain.Product) error {
	_, err := r.db.Exec("UPDATE products SET name=$1, description=$2, price=$3, stock=$4 WHERE id=$5",
		product.Name, product.Description, product.Price, product.Stock, product.ID)
	return err
}

func (r *PostgresProductRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id=$1", id)
	return err
}

func (r *PostgresProductRepository) FindByID(id string) (*domain.Product, error) {
	row := r.db.QueryRow("SELECT id, name, description, price, stock FROM products WHERE id=$1", id)
	product := &domain.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *PostgresProductRepository) FindAll() ([]*domain.Product, error) {
	rows, err := r.db.Query("SELECT id, name, description, price, stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*domain.Product{}
	for rows.Next() {
		product := &domain.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
