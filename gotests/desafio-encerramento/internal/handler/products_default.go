package handler

import (
	"app/internal"
	"app/internal/service"
	"app/plataform/web/response"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	service service.Service
}

func NewProductHandler(service service.Service) *ProductHandler {
	return &ProductHandler{
		service,
	}
}

// ProductJSON is an struct that represents a product in JSON format.
type ProductJSON struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	SellerId    int     `json:"seller_id"`
}

// Get returns a list of products that match the query.
func (h *ProductHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - query
		var query internal.ProductQuery
		if r.URL.Query().Has("id") {
			var err error
			query.Id, err = strconv.Atoi(r.URL.Query().Get("id"))
			if err != nil {
				response.Error(w, http.StatusBadRequest, "invalid id")
				return
			}
		}

		// process
		// - search products
		p, err := h.service.GetProducts(query)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "internal error")
			return
		}

		// response
		// - serialize products
		data := make(map[int]ProductJSON)
		for k, v := range p {
			data[k] = ProductJSON{
				Id:          v.Id,
				Description: v.Description,
				Price:       v.Price,
				SellerId:    v.SellerId,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}
