package product_test

import (
	"app/internal/domain"
	"app/internal/product"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-txdb"
)

func init() {
	txdb.Register("txdb", "mysql", "root:root@tcp(localhost:3306)/fantasy_products_test")
}

func openTestDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("txdb", fmt.Sprintf("db-%d", time.Now().UnixNano()))
	if err != nil {
		t.Fatalf("failed to open txdb: %v", err)
	}
	return db
}

func TestFindAll_Product(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := product.NewProductsMySQL(db)

	// given
	_, err := db.Exec(`INSERT INTO products (id, description, price) VALUES (?, ?, ?)`,
		10, "TESTE", 10.50)

	if err != nil {
		t.Fatalf("insert product: %v", err)
	}

	expected := []domain.Product{
		{
			Id: 10, ProductAttributes: domain.ProductAttributes{
				Description: "TESTE",
				Price:       10.50,
			},
		},
	}

	product, err := repo.FindAll()

	if err != nil {
		t.Fatalf("GetAll error: %v", err)
	}
	//then
	assert.Contains(t, product, expected[0])
}

func TestCreate_Product(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := product.NewProductsMySQL(db)

	//given
	// given
	input := domain.Product{
		Id: 11,
		ProductAttributes: domain.ProductAttributes{
			Description: "TESTE",
			Price:       11.50,
		},
	}

	//when
	err := repo.Save(&input)

	//then
	assert.NoError(t, err)

}

func TestFindTopProductsByAmountSold(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := product.NewProductsMySQL(db)

	//given
	expectedProducts := []domain.ProductAmountSold{
		{Description: "Product 4", Total: 200},
		{Description: "Product 3", Total: 113},
		{Description: "Product 6", Total: 70},
		{Description: "Product 1", Total: 50},
		{Description: "Product 2", Total: 50},
	}

	// when
	products, err := repo.FindTopProductsByAmountSold(5)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, products)
	assert.NotEmpty(t, products)
	assert.Equal(t, expectedProducts, products)
}
