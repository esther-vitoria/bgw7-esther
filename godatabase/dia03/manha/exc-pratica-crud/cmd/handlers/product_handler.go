package handlers

import (
	"net/http"
	"strconv"

	"github.com/bgw7/exc-pratica-crud/internal/domain"
	"github.com/bgw7/exc-pratica-crud/internal/product"
	"github.com/bgw7/exc-pratica-crud/pkg/web/request"
	"github.com/bgw7/exc-pratica-crud/pkg/web/response"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	service product.Service
}

func NewProductHandler(service product.Service) *ProductHandler {
	return &ProductHandler{
		service,
	}
}

func (h *ProductHandler) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := h.service.GetAll()

		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{"data": products})
	}
}

func (h *ProductHandler) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		productId, errParseID := strconv.Atoi(chi.URLParam(r, "productId"))

		if errParseID != nil {
			response.Error(w, http.StatusBadRequest, "could not parse Product ID")
			return
		}

		product, errGetId := h.service.GetById(productId)
		if errGetId != nil {
			response.Error(w, http.StatusBadRequest, errGetId.Error())
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{"data": product})
	}
}

func (h *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody := domain.ProductAttributes{}

		if err := request.JSON(r, &reqBody); err != nil {
			response.Error(w, http.StatusBadRequest, "could not parse body")
			return
		}
		product, err := h.service.Create(reqBody)

		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{"data": product})
	}
}

func (h *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		productId, errParseID := strconv.Atoi(chi.URLParam(r, "productId"))

		if errParseID != nil {
			response.Error(w, http.StatusBadRequest, "could not parse Product ID")
			return
		}

		product, errGetId := h.service.GetById(productId)
		if errGetId != nil {
			response.Error(w, http.StatusBadRequest, errGetId.Error())
			return
		}

		reqBody := product.ProductAttributes

		errParseBody := request.JSON(r, &reqBody)
		if errParseBody != nil {
			response.Error(w, http.StatusBadRequest, "could not parse body")
			return
		}

		product, errUpdateBody := h.service.Update(productId, reqBody)
		if errUpdateBody != nil {
			response.Error(w, http.StatusBadRequest, errUpdateBody.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{"data": product})
	}
}

func (h *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		productId, errParseID := strconv.Atoi(chi.URLParam(r, "productId"))

		if errParseID != nil {
			response.Error(w, http.StatusBadRequest, "could not parse Product ID")
			return
		}

		_, errGetId := h.service.GetById(productId)
		if errGetId != nil {
			response.Error(w, http.StatusBadRequest, errGetId.Error())
			return
		}

		errDelete := h.service.Delete(productId)
		if errDelete != nil {
			response.Error(w, http.StatusBadRequest, errDelete.Error())
			return
		}

		response.Text(w, http.StatusNoContent, "")
	}
}
