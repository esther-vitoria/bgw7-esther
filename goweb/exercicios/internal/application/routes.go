package application

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"google.com/bgw7-esther/first-server/internal/handler"
	"google.com/bgw7-esther/first-server/internal/middlewares"
	"google.com/bgw7-esther/first-server/internal/repository"
	"google.com/bgw7-esther/first-server/internal/service"
	"google.com/bgw7-esther/first-server/internal/storage"
)

type router struct{}

func NewRouter() *router {
	return &router{}
}

func (rt router) MapRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middlewares.TokenAuth)

	r.Route("/products", func(r chi.Router) {
		r.Mount("/", ProductsRoutes())
	})

	return r
}

func ProductsRoutes() http.Handler {
	// Carrega o arquivo JSON
	products := storage.LoadProducts("/Users/enbalbino/Documents/GitHub/bgw7-esther/goweb/exercicios/data/products.json")
	db := make(map[int]*storage.ProductJSON)
	for _, pr := range products {
		db[pr.Id] = pr
	}

	repo := repository.NewProductRepository(db)
	svc := service.NewProductService(*repo)
	h := handler.NewProductHandler(svc)
	r := chi.NewRouter()

	r.Post("/", h.Create())
	r.Patch("/{id}", h.Patch())
	r.Delete("/{id}", h.Delete())
	r.Put("/", h.UpdateProductFull())
	r.Get("/{id}", h.GetByID())
	r.Get("/", h.GetAll())
	r.Get("/search", h.SearchProducts())

	log.Println("Server running at :8080")

	return r
}
