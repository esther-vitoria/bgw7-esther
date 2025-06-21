package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"google.com/bgw7-esther/first-server/internal/handler"
	"google.com/bgw7-esther/first-server/internal/repository"
	"google.com/bgw7-esther/first-server/internal/service"
	"google.com/bgw7-esther/first-server/internal/storage"
	"google.com/bgw7-esther/first-server/internal/utils"
)

func TestGetAll(t *testing.T) {
	t.Run("GET:/products -> expected success", func(t *testing.T) {
		pr_map := map[int]*storage.ProductJSON{
			1: {Id: 1, Name: "Banana", Quantity: 5, CodeValue: "b1234", IsPublished: true, Expiration: "20/06/2025", Price: 10.90},
			2: {Id: 2, Name: "Maça", Quantity: 6, CodeValue: "m1234", IsPublished: true, Expiration: "20/06/2025", Price: 5.90},
		}

		repo := repository.NewProductRepository(pr_map)
		svc := service.NewProductService(*repo)
		h := handler.NewProductHandler(svc)

		req := httptest.NewRequest("GET", "/products", nil)
		res := httptest.NewRecorder()
		h.GetAll()(res, req)

		expectedCode := http.StatusOK

		expectedBody := &utils.RespondBodyProduct{
			Message: "Success!",
			Data:    pr_map}

		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		expectedJSON, err := json.Marshal(expectedBody)
		require.NoError(t, err)
		require.JSONEq(t, string(expectedJSON), res.Body.String())
		require.Equal(t, expectedHeader, res.Header())

	})
}

func TestGetById(t *testing.T) {
	t.Run("GET:/products{id} -> expected success", func(t *testing.T) {
		pr_map := map[int]*storage.ProductJSON{
			1: {Id: 1, Name: "Banana", Quantity: 5, CodeValue: "b1234", IsPublished: true, Expiration: "20/06/2025", Price: 10.90},
			2: {Id: 2, Name: "Maça", Quantity: 6, CodeValue: "m1234", IsPublished: true, Expiration: "20/06/2025", Price: 5.90},
		}

		repo := repository.NewProductRepository(pr_map)
		svc := service.NewProductService(*repo)
		h := handler.NewProductHandler(svc)

		req := httptest.NewRequest("GET", "/products/1", nil)
		res := httptest.NewRecorder()
		h.GetByID()(res, req)

		expectedCode := http.StatusBadRequest

		expectedBody := &utils.RespondBodyProduct{
			Message: "Invalid ID!",
			Data:    nil,
			Error:   true,
		}

		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		expectedJSON, err := json.Marshal(expectedBody)
		require.NoError(t, err)
		require.JSONEq(t, string(expectedJSON), res.Body.String())
		require.Equal(t, expectedHeader, res.Header())

	})
}
