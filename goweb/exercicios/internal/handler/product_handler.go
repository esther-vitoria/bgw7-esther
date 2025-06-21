package handler

// Handler: recebe solicitações do cliente, valida as entradas e retorna as respostas.

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"google.com/bgw7-esther/first-server/internal/service"
	"google.com/bgw7-esther/first-server/internal/storage"
	"google.com/bgw7-esther/first-server/internal/utils"
)

type ProductHandler struct {
	svc *service.ProductService
}

func NewProductHandler(svc *service.ProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

func (h *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var reqBody *storage.ProductJSON
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			utils.Respond(w, http.StatusBadRequest, &utils.RespondBodyProduct{
				Message: "Bad Request!",
				Data:    nil,
				Error:   true,
			})
			return
		}

		product, err := h.svc.CreateProduct(reqBody)
		if err != nil {
			utils.Respond(w, http.StatusBadRequest, &utils.RespondBodyProduct{
				Message: err.Error(),
				Data:    nil,
				Error:   true,
			})
			return
		}
		utils.Respond(w, http.StatusCreated, &utils.RespondBodyProduct{
			Message: "Product Created!",
			Data:    product,
		})
	}
}

func (h *ProductHandler) UpdateProductFull() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var reqBody *storage.ProductJSON
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			utils.Respond(w, http.StatusBadRequest, &utils.RespondBodyProduct{
				Message: "Bad Request!",
				Data:    nil,
				Error:   true,
			})
			return
		}

		product, err := h.svc.UpdateProductFull(reqBody)
		if err != nil {
			utils.Respond(w, http.StatusBadRequest, &utils.RespondBodyProduct{
				Message: err.Error(),
				Data:    nil,
				Error:   true,
			})
			return
		}
		utils.Respond(w, http.StatusOK, &utils.RespondBodyProduct{
			Message: "Product updated!",
			Data:    product,
		})
	}
}

// Responde o produto baseado no id
func (h *ProductHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			utils.Respond(w, http.StatusBadRequest, &utils.RespondBodyProduct{
				Message: "Invalid ID!",
				Data:    nil,
				Error:   true,
			})
			return
		}
		product, err := h.svc.GetProductByID(id)
		if err != nil {
			utils.Respond(w, http.StatusNotFound, &utils.RespondBodyProduct{
				Message: err.Error(),
				Data:    nil,
				Error:   true,
			})
			return
		}
		utils.Respond(w, http.StatusOK, &utils.RespondBodyProduct{
			Message: "Success!",
			Data:    product,
		})
	}
}

// Responde todos os produtos do arquivo
func (h *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := h.svc.GetAll()

		if err != nil {
			utils.Respond(w, http.StatusBadRequest, &utils.RespondBodyProduct{
				Message: err.Error(),
				Data:    nil,
				Error:   true,
			})
			return
		}

		utils.Respond(w, http.StatusOK, &utils.RespondBodyProduct{
			Message: "Success!",
			Data:    products,
		})
	}
}

// utils.Respond the produts based on the price put on the url
func (h *ProductHandler) SearchProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		priceGtStr := r.URL.Query().Get("priceGt")

		if priceGtStr == "" {
			utils.Respond(w, http.StatusBadRequest, &utils.RespondBodyProduct{
				Message: "priceGt param is required!",
				Data:    nil,
				Error:   true,
			})
			return
		}

		priceGt, err := strconv.ParseFloat(priceGtStr, 64)
		if err != nil {
			utils.Respond(w, http.StatusBadRequest, &utils.RespondBodyProduct{
				Message: "priceGt must be a number!",
				Data:    nil,
				Error:   true,
			})
			return
		}

		products, err := h.svc.SearchProducts(priceGt)
		if err != nil {
			//utils.Respond(w, http.StatusBadRequest, ResponseBodyProduct{err.Error(), nil, true})
			utils.Respond(w, http.StatusBadRequest, &utils.RespondBodyProduct{
				Message: err.Error(),
				Data:    nil,
				Error:   true,
			})
			return
		}
		utils.Respond(w, http.StatusOK, &utils.RespondBodyProduct{
			Message: "Success!",
			Data:    products,
		})
	}
}

// Faz update parcial dos campos
func (h *ProductHandler) Patch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req *storage.ProductJSON
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.Respond(w, http.StatusBadRequest, &utils.RespondBodyProduct{
				Message: "Invalid JSON body!",
				Data:    nil,
				Error:   true,
			})
			return
		}
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			utils.Respond(w, http.StatusBadRequest, &utils.RespondBodyProduct{
				Message: "Invalid ID!",
				Data:    nil,
				Error:   true,
			})
			return
		}

		updated, err := h.svc.PatchProduct(id, &req.Name, &req.Quantity, &req.CodeValue, &req.Expiration, &req.IsPublished, &req.Price)
		if err != nil {
			utils.Respond(w, http.StatusBadRequest, &utils.RespondBodyProduct{
				Message: err.Error(),
				Data:    nil,
				Error:   true,
			})
			return
		}
		utils.Respond(w, http.StatusOK, &utils.RespondBodyProduct{
			Message: "Product patched!",
			Data:    updated,
		})
	}
}

// Deleta um produto baseado no id inserido
func (h *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			utils.Respond(w, http.StatusBadRequest, &utils.RespondBodyProduct{
				Message: "Invalid ID!",
				Data:    nil,
				Error:   true,
			})
			return
		}

		err = h.svc.DeleteProduct(id)
		if err != nil {
			utils.Respond(w, http.StatusNotFound, &utils.RespondBodyProduct{
				Message: err.Error(),
				Data:    nil,
				Error:   true,
			})
			return
		}
		utils.Respond(w, http.StatusNoContent, nil)
	}
}
