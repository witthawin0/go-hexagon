package adapters

import (
	"database/sql"

	"github.com/witthawin0/go-hexagon/core"
)

type postgresOrderReposistory struct {
	db *sql.DB
}

func NewPostgresOrderReposistory(db *sql.DB) core.OrderRepository {
	return &postgresOrderReposistory{db: db}
}

func (r *postgresOrderReposistory) Save(order core.Order) error {
	return nil
}
