package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.com/bgw7-esther/first-server/handlers"
)

func main() {
	products, err := handlers.ShowProducts("data/products.json")
	if err != nil {
		log.Fatal("failed to load products:", err)
	}

	r := chi.NewRouter()
	r.Get("/products", handlers.GetProductHandler(products))
	r.Get("/products/search", handlers.SearchProductsHandler(products))
	r.Get("/products/{id}", handlers.GetProductByIdHandler(products))

	log.Println("Server running at :8080")
	http.ListenAndServe(":8080", r)
}
