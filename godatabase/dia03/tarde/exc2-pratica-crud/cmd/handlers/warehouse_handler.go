package handlers

import (
	"net/http"
	"strconv"

	"github.com/bgw7/exc-pratica-crud/internal/domain"
	"github.com/bgw7/exc-pratica-crud/internal/warehouse"
	"github.com/bgw7/exc-pratica-crud/pkg/web/request"
	"github.com/bgw7/exc-pratica-crud/pkg/web/response"
	"github.com/go-chi/chi/v5"
)

type WarehouseHandler struct {
	warehouseService warehouse.Service
}

func NewWarehouseHandler(warehouseService warehouse.Service) *WarehouseHandler {
	return &WarehouseHandler{
		warehouseService,
	}
}

func (h *WarehouseHandler) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		warehouses, err := h.warehouseService.GetAll()

		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{"data": warehouses})
	}
}

func (h *WarehouseHandler) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		warehouseId, errParseID := strconv.Atoi(chi.URLParam(r, "warehouseId"))

		if errParseID != nil {
			response.Error(w, http.StatusBadRequest, "could not parse Warehouse ID")
			return
		}

		warehouse, errGetId := h.warehouseService.GetById(warehouseId)
		if errGetId != nil {
			response.Error(w, http.StatusBadRequest, errGetId.Error())
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{"data": warehouse})
	}
}

func (h *WarehouseHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody := domain.WarehouseAttributes{}

		if err := request.JSON(r, &reqBody); err != nil {
			response.Error(w, http.StatusBadRequest, "could not parse body")
			return
		}
		warehouse, err := h.warehouseService.Create(reqBody)

		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{"data": warehouse})
	}
}

func (h *WarehouseHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		warehouseId, errParseID := strconv.Atoi(chi.URLParam(r, "warehouseId"))

		if errParseID != nil {
			response.Error(w, http.StatusBadRequest, "could not parse Warehouse ID")
			return
		}

		warehouse, errGetId := h.warehouseService.GetById(warehouseId)
		if errGetId != nil {
			response.Error(w, http.StatusBadRequest, errGetId.Error())
			return
		}

		reqBody := warehouse.WarehouseAttributes

		errParseBody := request.JSON(r, &reqBody)
		if errParseBody != nil {
			response.Error(w, http.StatusBadRequest, "could not parse body")
			return
		}

		warehouse, errUpdateBody := h.warehouseService.Update(warehouseId, reqBody)
		if errUpdateBody != nil {
			response.Error(w, http.StatusBadRequest, errUpdateBody.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{"data": warehouse})
	}
}

func (h *WarehouseHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		warehouseId, errParseID := strconv.Atoi(chi.URLParam(r, "warehouseId"))

		if errParseID != nil {
			response.Error(w, http.StatusBadRequest, "could not parse Warehouse ID")
			return
		}

		_, errGetId := h.warehouseService.GetById(warehouseId)
		if errGetId != nil {
			response.Error(w, http.StatusBadRequest, errGetId.Error())
			return
		}

		errDelete := h.warehouseService.Delete(warehouseId)
		if errDelete != nil {
			response.Error(w, http.StatusBadRequest, errDelete.Error())
			return
		}

		response.Text(w, http.StatusNoContent, "")
	}
}

func (h *WarehouseHandler) GetReport() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		warehouseId, errParseID := strconv.Atoi(chi.URLParam(r, "warehouseId"))

		if errParseID != nil {
			response.Error(w, http.StatusBadRequest, "could not parse Warehouse ID")
			return
		}

		warehouse, errGetId := h.warehouseService.GetReport(warehouseId)
		if errGetId != nil {
			response.Error(w, http.StatusBadRequest, errGetId.Error())
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{"data": warehouse})
	}
}
