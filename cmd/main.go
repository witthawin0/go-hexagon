package main

import (
	"log"
	"net/http"

	"github.com/witthawin0/go-hexagon/internal/adapters/repository"
	"github.com/witthawin0/go-hexagon/internal/application"
	"github.com/witthawin0/go-hexagon/internal/infrastructure"
	"github.com/witthawin0/go-hexagon/internal/infrastructure/handlers"
)

func main() {
	// Initialize the database connection
	db := infrastructure.InitDB("your_connection_string")

	// Create the repository
	productRepo := repository.NewPostgresProductRepository(db)

	// Create the service
	productService := application.NewProductService(productRepo)

	// Create the handlers
	productHandler := handlers.NewProductHandler(productService)

	// Set up the routes
	http.HandleFunc("/products", productHandler.GetAllProducts)
	http.HandleFunc("/products/create", productHandler.CreateProduct)
	http.HandleFunc("/products/update", productHandler.UpdateProduct)
	http.HandleFunc("/products/delete", productHandler.DeleteProduct)
	http.HandleFunc("/products/get", productHandler.GetProductByID)

	// Start the HTTP server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
