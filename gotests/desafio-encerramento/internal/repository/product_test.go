package repository

import (
	"app/internal"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductsJSON_SearchProducts(t *testing.T) {
	// Setup - using the existing JSON file
	repo, err := NewProductsJSON("../../database/json/products.json")
	require.NoError(t, err)
	defer repo.Close()

	t.Run("search all products", func(t *testing.T) {
		// Act
		query := internal.ProductQuery{}
		products, err := repo.SearchProducts(query)

		// Assert
		require.NoError(t, err)
		assert.Greater(t, len(products), 0, "should return products")
		assert.Equal(t, 100, len(products), "should return all 100 products")
	})

	t.Run("search product by ID", func(t *testing.T) {
		// Act
		query := internal.ProductQuery{Id: 1}
		products, err := repo.SearchProducts(query)

		// Assert
		require.NoError(t, err)
		assert.Len(t, products, 1, "should return exactly one product")

		product, exists := products[1]
		assert.True(t, exists, "product with ID 1 should exist")
		assert.Equal(t, 1, product.Id)
		assert.Equal(t, "French Pastry - Mini Chocolate", product.Description)
		assert.Equal(t, 97.01, product.Price)
		assert.Equal(t, 1, product.SellerId)
	})

	t.Run("search non-existent product", func(t *testing.T) {
		// Act
		query := internal.ProductQuery{Id: 999}
		products, err := repo.SearchProducts(query)

		// Assert
		require.NoError(t, err)
		assert.Empty(t, products, "should return empty map for non-existent product")
	})
}

func TestNewProductsJSON_FileNotFound(t *testing.T) {
	// Act
	repo, err := NewProductsJSON("non_existent_file.json")

	// Assert
	assert.Error(t, err, "should return error for non-existent file")
	assert.Nil(t, repo, "repository should be nil when file not found")
}

func TestNewProductsJSON_InvalidJSON(t *testing.T) {
	// Setup - create a temporary invalid JSON file
	tmpFile, err := os.CreateTemp("", "invalid*.json")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString("invalid json content")
	require.NoError(t, err)
	tmpFile.Close()

	// Act
	repo, err := NewProductsJSON(tmpFile.Name())

	// Assert
	assert.Error(t, err, "should return error for invalid JSON")
	assert.Nil(t, repo, "repository should be nil when JSON is invalid")
}
