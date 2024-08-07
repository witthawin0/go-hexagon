package adapters

import (
	"github.com/witthawin0/go-hexagon/core"
	"gorm.io/gorm"
)

// Secondary adapter
type gormOrderReposistory struct {
	db *gorm.DB
}

func NewGormOrderReposistory(db *gorm.DB) core.OrderRepository {
	return &gormOrderReposistory{db: db}
}

func (r *gormOrderReposistory) Save(order core.Order) error {
	if result := r.db.Create(&order); result.Error != nil {
		return result.Error
	}

	return nil
}
