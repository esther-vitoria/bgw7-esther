package service

import (
	"app/internal"
	"app/internal/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductService_GetProducts(t *testing.T) {

	t.Run("GetProducts -> expected success", func(t *testing.T) {
		// Arrange
		mockRepository := repository.NewProductRepositoryMock()
		service := NewProductService(mockRepository)

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

		query := internal.ProductQuery{Id: 0}

		// Setup mock expectations
		mockRepository.On("SearchProducts", query).Return(expectedProducts, nil)

		// Act
		result, err := service.GetProducts(query)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedProducts, result)
		assert.Len(t, result, 2)

		// Verify specific product data
		product1, exists := result[1]
		assert.True(t, exists)
		assert.Equal(t, 1, product1.Id)
		assert.Equal(t, "Product 1", product1.Description)
		assert.Equal(t, 10.50, product1.Price)
		assert.Equal(t, 100, product1.SellerId)

		product2, exists := result[2]
		assert.True(t, exists)
		assert.Equal(t, 2, product2.Id)
		assert.Equal(t, "Product 2", product2.Description)
		assert.Equal(t, 25.99, product2.Price)
		assert.Equal(t, 200, product2.SellerId)

		// Verify mock expectations
		mockRepository.AssertExpectations(t)
	})

	t.Run("GetProducts -> search by specific id", func(t *testing.T) {
		// Arrange
		mockRepository := repository.NewProductRepositoryMock()
		service := NewProductService(mockRepository)

		// Mock data
		expectedProducts := map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "Specific Product",
					Price:       99.99,
					SellerId:    300,
				},
			},
		}

		query := internal.ProductQuery{Id: 1}

		// Setup mock expectations
		mockRepository.On("SearchProducts", query).Return(expectedProducts, nil)

		// Act
		result, err := service.GetProducts(query)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedProducts, result)
		assert.Len(t, result, 1)

		// Verify specific product data
		product, exists := result[1]
		assert.True(t, exists)
		assert.Equal(t, 1, product.Id)
		assert.Equal(t, "Specific Product", product.Description)
		assert.Equal(t, 99.99, product.Price)
		assert.Equal(t, 300, product.SellerId)

		// Verify mock expectations
		mockRepository.AssertExpectations(t)
	})

	t.Run("GetProducts -> empty result", func(t *testing.T) {
		// Arrange
		mockRepository := repository.NewProductRepositoryMock()
		service := NewProductService(mockRepository)

		// Mock data
		expectedProducts := map[int]internal.Product{}
		query := internal.ProductQuery{Id: 999}

		// Setup mock expectations
		mockRepository.On("SearchProducts", query).Return(expectedProducts, nil)

		// Act
		result, err := service.GetProducts(query)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedProducts, result)
		assert.Len(t, result, 0)
		assert.Empty(t, result)

		// Verify mock expectations
		mockRepository.AssertExpectations(t)
	})

	t.Run("GetProducts -> repository error", func(t *testing.T) {
		// Arrange
		mockRepository := repository.NewProductRepositoryMock()
		service := NewProductService(mockRepository)

		query := internal.ProductQuery{Id: 0}
		expectedError := errors.New("database connection error")

		// Setup mock expectations
		mockRepository.On("SearchProducts", query).Return(map[int]internal.Product{}, expectedError)

		// Act
		result, err := service.GetProducts(query)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, "database connection error", err.Error())
		assert.Empty(t, result)

		// Verify mock expectations
		mockRepository.AssertExpectations(t)
	})

	t.Run("GetProducts -> internal server error", func(t *testing.T) {
		// Arrange
		mockRepository := repository.NewProductRepositoryMock()
		service := NewProductService(mockRepository)

		query := internal.ProductQuery{Id: 0}
		expectedError := errors.New("internal server error")

		// Setup mock expectations
		mockRepository.On("SearchProducts", query).Return(map[int]internal.Product{}, expectedError)

		// Act
		result, err := service.GetProducts(query)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, "internal server error", err.Error())
		assert.Empty(t, result)

		// Verify mock expectations
		mockRepository.AssertExpectations(t)
	})

	t.Run("GetProducts -> nil repository error", func(t *testing.T) {
		// Arrange
		mockRepository := repository.NewProductRepositoryMock()
		service := NewProductService(mockRepository)

		query := internal.ProductQuery{Id: 0}

		// Setup mock expectations - return nil error but empty products
		mockRepository.On("SearchProducts", query).Return(map[int]internal.Product{}, nil)

		// Act
		result, err := service.GetProducts(query)

		// Assert
		assert.NoError(t, err)
		assert.Empty(t, result)
		assert.NotNil(t, result) // map should be initialized, not nil

		// Verify mock expectations
		mockRepository.AssertExpectations(t)
	})
}
