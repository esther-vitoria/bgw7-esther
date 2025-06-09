package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"google.com/bgw7-esther/first-server/internal/domain"
	"google.com/bgw7-esther/first-server/internal/handler"
	"google.com/bgw7-esther/first-server/internal/repository"
	"google.com/bgw7-esther/first-server/internal/service"
)

func main() {
	// Carrega do arquivo JSON
	products := loadProducts("/Users/enbalbino/Documents/GitHub/bgw7-esther/goweb/dia02/manha/data/products.json")
	db := make(map[int]*domain.Product)
	for _, pr := range products {
		db[pr.Id] = pr
	}

	repo := repository.NewProductRepository(db)
	svc := service.NewProductService(*repo)
	h := handler.NewProductHandler(svc)

	r := chi.NewRouter()

	r.Post("/products", h.Create())
	r.Get("/products", h.GetAll())
	r.Get("/products/{id}", h.GetByID())
	r.Get("/products/search", h.GetByMinPrice())

	log.Println("Server running at :8080")
	http.ListenAndServe(":8080", r)
}

func loadProducts(path string) []*domain.Product {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var products []*domain.Product
	_ = json.Unmarshal(data, &products)
	return products
}
