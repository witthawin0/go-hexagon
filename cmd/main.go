package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/witthawin0/go-hexagon/adapters"
	"github.com/witthawin0/go-hexagon/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	connStr := "postgresql://neon_owner:bhj9JAfC6rNx@ep-floral-hall-a15enrat.ap-southeast-1.aws.neon.tech/neon?sslmode=require"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&core.Order{})

	orderRepo := adapters.NewGormOrderReposistory(db)

	orderService := core.NewOrderService(orderRepo)

	orderhandler := adapters.NewHttpOrderHanlder(orderService)

	app.Post("/orders", orderhandler.CreateOrder)

	app.Listen(":1234")
}
