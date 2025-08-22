package repository

import (
	"app/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductsJSON_SearchProducts(t *testing.T) {

	t.Run("SearchProducts -> success - return all products", func(t *testing.T) {
		// Arrange
		mockData := map[int]internal.Product{
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

		repository := &ProductsJSON{
			file: nil,
			db:   mockData,
		}

		query := internal.ProductQuery{Id: 0} // Id = 0 means return all products

		// Act
		result, err := repository.SearchProducts(query)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, mockData, result)
		assert.Len(t, result, 2)
	})

	t.Run("SearchProducts -> failure - product not found", func(t *testing.T) {
		// Arrange
		repository := &ProductsJSON{
			db: map[int]internal.Product{},
		}

		query := internal.ProductQuery{Id: 99}

		// Act
		result, err := repository.SearchProducts(query)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "product not found", err.Error())
		assert.Nil(t, result)
	})
}

func TestProductsJSON_Close(t *testing.T) {

	t.Run("Close -> success - file is nil", func(t *testing.T) {
		// Arrange
		repository := &ProductsJSON{
			file: nil,
			db:   make(map[int]internal.Product),
		}

		// Act
		err := repository.Close()

		// Assert
		assert.NoError(t, err)
	})
}

func TestNewProductsJSON(t *testing.T) {

	t.Run("NewProductsJSON -> error - file not found", func(t *testing.T) {
		// Arrange
		nonExistentFile := "/path/to/non/existent/file.json"

		// Act
		repository, err := NewProductsJSON(nonExistentFile)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, repository)
		assert.Contains(t, err.Error(), "no such file or directory")
	})

	t.Run("NewProductsJSON -> success - using real products.json file", func(t *testing.T) {
		// Arrange
		// Usar o arquivo real de produtos do projeto
		realFile := "../../database/json/products.json"

		// Act
		repository, err := NewProductsJSON(realFile)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, repository)

		// Verificar se repository não é nil antes de acessar seus campos
		if repository != nil {
			assert.NotNil(t, repository.file)
			assert.NotNil(t, repository.db)
			assert.Len(t, repository.db, 100) // O arquivo tem 100 produtos

			// Verificar alguns produtos específicos
			product1, exists := repository.db[1]
			assert.True(t, exists)
			assert.Equal(t, 1, product1.Id)
			assert.Equal(t, "French Pastry - Mini Chocolate", product1.Description)
			assert.Equal(t, 97.01, product1.Price)
			assert.Equal(t, 1, product1.SellerId)

			product100, exists := repository.db[100]
			assert.True(t, exists)
			assert.Equal(t, 100, product100.Id)
			assert.Equal(t, "Vinegar - Raspberry", product100.Description)
			assert.Equal(t, 58.78, product100.Price)
			assert.Equal(t, 1, product100.SellerId)

			// Cleanup
			repository.Close()
		}
	})
}
