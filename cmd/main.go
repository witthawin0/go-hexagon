package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/witthawin0/go-hexagon/internal/adapters/repository"
	"github.com/witthawin0/go-hexagon/internal/application"
	"github.com/witthawin0/go-hexagon/internal/infrastructure/handlers"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=myuser password=mypassword dbname=mydatabase sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Initialize repositories
	productRepo := repository.NewPostgresProductRepository(db)
	// customerRepo := repository.NewPostgresCustomerRepository(db)
	orderRepo := repository.NewPostgresOrderRepository(db)

	// Initialize services
	productService := application.NewProductService(productRepo)
	// customerService := application.NewCustomerService(customerRepo)
	orderService := application.NewOrderService(orderRepo)

	// Initialize handlers
	productHandler := handlers.NewProductHandler(productService)
	// customerHandler := handlers.NewCustomerHandler(customerService)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Define routes
	http.HandleFunc("GET /products", productHandler.GetAllProducts)
	http.HandleFunc("POST /products", productHandler.CreateProduct)
	http.HandleFunc("GET /products/:id", productHandler.GetProductByID)
	http.HandleFunc("PUT /products/:id", productHandler.UpdateProduct)
	http.HandleFunc("DELETE /products/:id", productHandler.DeleteProduct)

	http.HandleFunc("GET /orders", orderHandler.GetAllOrders)
	http.HandleFunc("POST /orders", orderHandler.CreateOrder)
	http.HandleFunc("GET /orders/:id", orderHandler.GetOrder)
	http.HandleFunc("PUT /orders/:id", orderHandler.UpdateOrder)
	http.HandleFunc("DELETE /orders/:id", orderHandler.DeleteOrder)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
