package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type router struct {
}

func (router *router) MapRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.StripSlashes,
		middleware.Timeout(5*time.Second),
		middleware.Heartbeat("/ping"),
	)

	r.Route("/api/v1", func(rp chi.Router) {

		rp.Route("/products", func(rp chi.Router) {
			rp.Mount("/", buildProductsRoutes())
		})

		rp.Route("/warehouses", func(rp chi.Router) {
			rp.Mount("/", buildWarehousesRoutes())
		})

	})

	return r
}

func NewRouter() *router {
	return &router{}
}
