package handler

import (
	"app/internal"
	"app/internal/service"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductHandler_Get(t *testing.T) {

	t.Run("GET:/products -> expected success", func(t *testing.T) {
		// Arrange
		mockService := service.NewProductServiceMock()
		handler := NewProductHandler(mockService)

		// Mock data
		expectedProducts := map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "Product 1",
					Price:       10.50,
					SellerId:    100,
				},
			},
			2: {
				Id: 2,
				ProductAttributes: internal.ProductAttributes{
					Description: "Product 2",
					Price:       25.99,
					SellerId:    200,
				},
			},
		}

		// Setup mock expectations
		mockService.On("GetProducts", internal.ProductQuery{Id: 0}).Return(expectedProducts, nil)

		// Create request
		req, err := http.NewRequest("GET", "/products", nil)
		assert.NoError(t, err)

		// Create response recorder
		rr := httptest.NewRecorder()

		// Act
		handler.Get()(rr, req)

		// Assert
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

		// Parse response body
		var response map[string]interface{}
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		assert.NoError(t, err)

		// Verify response structure
		assert.Equal(t, "success", response["message"])
		assert.Contains(t, response, "data")

		// Verify mock expectations
		mockService.AssertExpectations(t)
	})

	t.Run("GET:/products -> invalid id", func(t *testing.T) {
		// Arrange
		mockService := service.NewProductServiceMock()
		handler := NewProductHandler(mockService)

		// Create request
		req, err := http.NewRequest("GET", "/products?id=invalid", nil)
		assert.NoError(t, err)

		// Create response recorder
		rr := httptest.NewRecorder()

		// Act
		handler.Get()(rr, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

		// Parse response body
		var response map[string]interface{}
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		assert.NoError(t, err)

		// Verify error response structure
		assert.Equal(t, "Bad Request", response["status"])
		assert.Equal(t, "invalid id", response["message"])

		// Verify mock expectations
		mockService.AssertExpectations(t)
	})

	t.Run("GET:/products -> internal error", func(t *testing.T) {
		// Arrange
		mockService := service.NewProductServiceMock()
		handler := NewProductHandler(mockService)

		// Setup mock expectations
		mockService.On("GetProducts", internal.ProductQuery{Id: 0}).Return(map[int]internal.Product{}, errors.New("internal error"))

		// Create request
		req, err := http.NewRequest("GET", "/products", nil)
		assert.NoError(t, err)

		// Create response recorder
		rr := httptest.NewRecorder()

		// Act
		handler.Get()(rr, req)

		// Assert
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

		// Parse response body
		var response map[string]interface{}
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "internal error", response["message"])

		// Verify error response structure
		assert.Equal(t, "Internal Server Error", response["status"])

		// Verify mock expectations
		mockService.AssertExpectations(t)
	})
}
