package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"google.com/bgw7-esther/first-server/internal/domain"
	"google.com/bgw7-esther/first-server/internal/service"
)

type ProductHandler struct {
	svc *service.ProductService
}

// Struct que pega os dados do body da requisição
type RequestBodyProduct struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// Struct que trás a resposta do POST
type ResponseBodyProduct struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   bool
}

func NewProductHandler(svc *service.ProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

func (h *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody RequestBodyProduct
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			respond(w, http.StatusBadRequest, ResponseBodyProduct{"Bad Request", nil, true})
			return
		}

		p := &domain.Product{
			Id:          h.svc.Repo.GetNewID(),
			Name:        reqBody.Name,
			Quantity:    reqBody.Quantity,
			CodeValue:   reqBody.CodeValue,
			IsPublished: reqBody.IsPublished,
			Expiration:  reqBody.Expiration,
			Price:       reqBody.Price,
		}
		if err := h.svc.CreateProduct(p); err != nil {
			respond(w, http.StatusBadRequest, ResponseBodyProduct{err.Error(), nil, true})
			return
		}
		respond(w, http.StatusCreated, ResponseBodyProduct{"Product created", p, false})
	}
}

func (h *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := h.svc.GetAllProducts()
		if err != nil {
			respond(w, http.StatusInternalServerError, ResponseBodyProduct{"Error listing products", nil, true})
			return
		}
		respond(w, http.StatusOK, ResponseBodyProduct{"Success", products, false})
	}
}

func (h *ProductHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			respond(w, http.StatusBadRequest, ResponseBodyProduct{"Invalid ID", nil, true})
			return
		}
		product, err := h.svc.GetProductByID(id)
		if err != nil {
			respond(w, http.StatusNotFound, ResponseBodyProduct{err.Error(), nil, true})
			return
		}
		respond(w, http.StatusOK, ResponseBodyProduct{"Success", product, false})
	}
}

func (h *ProductHandler) GetByMinPrice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		minPriceStr := r.URL.Query().Get("minPrice")
		minPrice, err := strconv.ParseFloat(minPriceStr, 64)
		if err != nil {
			respond(w, http.StatusBadRequest, ResponseBodyProduct{"Invalid minPrice value", nil, true})
			return
		}
		products, err := h.svc.GetProductsByPrice(minPrice)
		if err != nil {
			respond(w, http.StatusInternalServerError, ResponseBodyProduct{"Error fetching products", nil, true})
			return
		}
		respond(w, http.StatusOK, ResponseBodyProduct{"Success", products, false})
	}
}

func respond(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
