package loader

import (
	"app/internal"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadJSON(t *testing.T) {
	t.Run("NewProductsJSON -> expected success", func(t *testing.T) {
		// Arrange
		file, err := os.CreateTemp("", "test_products_*.json")
		if err != nil {
			t.Fatalf("Falha ao criar arquivo temporário: %v", err)
		}
		defer os.Remove(file.Name())
		defer file.Close()

		// Act
		loader := NewProductsJSON(file)

		// Assert
		if loader == nil {
			t.Fatal("NewProductsJSON retornou nil")
		}

		if loader.file != file {
			t.Error("NewProductsJSON não definiu o arquivo corretamente")
		}
	})

	t.Run("Load -> expected success", func(t *testing.T) {
		// Arrange
		file, err := os.CreateTemp("", "test_products_*.json")
		if err != nil {
			t.Fatalf("Falha ao criar arquivo temporário: %v", err)
		}
		defer os.Remove(file.Name())
		defer file.Close()

		// Write test data to the file
		testData := `[
			{
				"id": 1,
				"description": "French Pastry - Mini Chocolate",
				"price": 97.01,
				"seller_id": 1
			},
			{
				"id": 2,
				"description": "French Pastry - Mini Vanilla",
				"price": 85.50,
				"seller_id": 2
			}
		]`

		_, err = file.WriteString(testData)
		if err != nil {
			t.Fatalf("Falha ao escrever dados no arquivo: %v", err)
		}

		// Reset file pointer to beginning
		_, err = file.Seek(0, 0)
		if err != nil {
			t.Fatalf("Falha ao resetar ponteiro do arquivo: %v", err)
		}

		// Act
		loader := NewProductsJSON(file)

		// Assert
		products, _ := loader.Load()

		expectedProduct1 := internal.Product{
			Id: 1,
			ProductAttributes: internal.ProductAttributes{
				Description: "French Pastry - Mini Chocolate",
				Price:       97.01,
				SellerId:    1,
			},
		}

		expectedProduct2 := internal.Product{
			Id: 2,
			ProductAttributes: internal.ProductAttributes{
				Description: "French Pastry - Mini Vanilla",
				Price:       85.50,
				SellerId:    2,
			},
		}

		assert.Equal(t, expectedProduct1, products[0])
		assert.Equal(t, expectedProduct2, products[1])
	})

	t.Run("Load -> error decoding json file", func(t *testing.T) {
		// Arrange
		file, err := os.CreateTemp("", "test_products_*.json")
		if err != nil {
			t.Fatalf("Falha ao criar arquivo temporário: %v", err)
		}
		defer os.Remove(file.Name())
		defer file.Close()

		// Write invalid JSON data to the file
		invalidJSONData := `[
			{
				"id": 1,
				"description": "French Pastry - Mini Chocolate",
				"price": 97.01,
				"seller_id": 1
			},
			{
				"id": 2,
				"description": "French Pastry - Mini Vanilla",
				"price": "invalid_price", // Invalid price type - should be number
				"seller_id": 2
			}
		` // Missing closing bracket to make JSON invalid

		_, err = file.WriteString(invalidJSONData)
		if err != nil {
			t.Fatalf("Falha ao escrever dados no arquivo: %v", err)
		}

		// Reset file pointer to beginning
		_, err = file.Seek(0, 0)
		if err != nil {
			t.Fatalf("Falha ao resetar ponteiro do arquivo: %v", err)
		}

		// Act
		loader := NewProductsJSON(file)
		products, err := loader.Load()

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "error decoding json file", err.Error())
		assert.Empty(t, products)
	})
}
