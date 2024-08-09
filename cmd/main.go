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

	router := http.NewServeMux()

	// Initialize repositories
	productRepo := repository.NewPostgresProductRepository(db)
	customerRepo := repository.NewPosgresCustomerRepository(db)
	orderRepo := repository.NewPostgresOrderRepository(db)

	// Initialize services
	productService := application.NewProductService(productRepo)
	customerService := application.NewCustomerService(customerRepo)
	orderService := application.NewOrderService(orderRepo)

	// Initialize handlers
	productHandler := handlers.NewProductHandler(productService)
	customerHandler := handlers.NewCustomerHandler(customerService)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Define routes

	// router.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("create a todo")
	// 	w.WriteHeader(http.StatusCreated)
	// 	w.Write([]byte("Todo created"))
	// }).Methods(http.MethodPost)

	router.HandleFunc("GET /products", productHandler.GetAllProducts)
	router.HandleFunc("POST /products", productHandler.CreateProduct)
	router.HandleFunc("GET /products/{id}", productHandler.GetProductByID)
	router.HandleFunc("PUT /products/{id}", productHandler.UpdateProduct)
	router.HandleFunc("DELETE /products/{id}", productHandler.DeleteProduct)

	router.HandleFunc("GET /orders", orderHandler.GetAllOrders)
	router.HandleFunc("POST /orders", orderHandler.CreateOrder)
	router.HandleFunc("GET /orders/{id}", orderHandler.GetOrder)
	router.HandleFunc("PUT /orders/{id}", orderHandler.UpdateOrder)
	router.HandleFunc("DELETE /orders/{id}", orderHandler.DeleteOrder)

	router.HandleFunc("GET /customers", customerHandler.GetAllCustomers)
	router.HandleFunc("POST /customers", customerHandler.CreateCustomer)
	router.HandleFunc("GET /customers/{id}", customerHandler.GetCustomerByID)
	router.HandleFunc("PUT /customers/{id}", customerHandler.UpdateCustomer)
	router.HandleFunc("DELETE /customers/{id}", customerHandler.DeleteCustomer)

	log.Fatal(http.ListenAndServe(":8080", router))
}
