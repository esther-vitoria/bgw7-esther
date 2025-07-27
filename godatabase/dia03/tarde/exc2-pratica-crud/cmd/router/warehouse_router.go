package router

import (
	"net/http"

	"github.com/bgw7/exc-pratica-crud/cmd/handlers"
	"github.com/bgw7/exc-pratica-crud/database"
	"github.com/bgw7/exc-pratica-crud/internal/warehouse"
	"github.com/go-chi/chi/v5"
)

func buildWarehousesRoutes() http.Handler {
	r := chi.NewRouter()
	db := database.GetDB()

	repository := warehouse.NewMysqlRepository(db)
	service := warehouse.NewService(repository)
	handler := handlers.NewWarehouseHandler(service)

	r.Get("/", handler.FindAll())
	r.Get("/{warehouseId}", handler.Show())
	r.Get("/reportProducts/{warehouseId}", handler.GetReport())
	r.Post("/", handler.Create())
	r.Patch("/{warehouseId}", handler.Update())
	r.Delete("/{warehouseId}", handler.Delete())
	return r
}
