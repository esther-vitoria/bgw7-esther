package router

import (
	"net/http"

	"github.com/bgw7/exc-pratica-crud/cmd/handlers"
	"github.com/bgw7/exc-pratica-crud/database"
	"github.com/bgw7/exc-pratica-crud/internal/product"
	"github.com/go-chi/chi/v5"
)

func buildProductsRoutes() http.Handler {
	r := chi.NewRouter()
	db := database.GetDB()

	repository := product.NewMysqlRepository(db)
	service := product.NewService(repository)
	handler := handlers.NewProductHandler(service)

	r.Get("/", handler.FindAll())
	r.Get("/{productId}", handler.Show())
	r.Post("/", handler.Create())
	r.Patch("/{productId}", handler.Update())
	r.Delete("/{productId}", handler.Delete())
	return r
}
