package product_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-txdb"
	"github.com/bgw7/exc-pratica-crud/internal/domain"
	"github.com/bgw7/exc-pratica-crud/internal/product"
)

func init() {
	txdb.Register("txdb", "mysql", "root:root@tcp(localhost:3306)/my_db_test")
}

func openTestDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("txdb", fmt.Sprintf("db-%d", time.Now().UnixNano()))
	if err != nil {
		t.Fatalf("failed to open txdb: %v", err)
	}
	return db
}

func TestGetAll_Product(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := product.NewMysqlRepository(db)

	// given
	_, err := db.Exec(`INSERT INTO products 
        (id, name, quantity, code_value, is_published, expiration, price, warehouse_id)
        VALUES (?,?,?,?,?,?,?,?)`, 2, "TESTE", 229, "68788-9100", true, "2022-03-06", 30.75, 1)

	if err != nil {
		t.Fatalf("insert product: %v", err)
	}

	expected := []domain.Product{
		{
			Id: 1, ProductAttributes: domain.ProductAttributes{
				Name:        "Corn Shoots",
				Quantity:    244,
				CodeValue:   "0009-1111",
				IsPublished: false,
				Expiration:  "2022-01-08",
				Price:       23.27,
				WarehouseId: 1,
			},
		},
		{
			Id: 2, ProductAttributes: domain.ProductAttributes{
				Name:        "TESTE",
				Quantity:    229,
				CodeValue:   "68788-9100",
				IsPublished: true,
				Expiration:  "2022-03-06",
				Price:       30.75,
				WarehouseId: 1,
			},
		},
	}

	//when
	products, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll error: %v", err)
	}
	//then
	assert.NoError(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, expected, products)
}

func TestGetById_Product(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := product.NewMysqlRepository(db)

	// given
	_, err := db.Exec(`INSERT INTO products 
        (id, name, quantity, code_value, is_published, expiration, price, warehouse_id)
        VALUES (?,?,?,?,?,?,?,?)`, 2, "TESTE", 229, "68788-9100", true, "2022-03-06", 30.75, 1)

	if err != nil {
		t.Fatalf("insert product: %v", err)
	}

	expected := domain.Product{
		Id: 2, ProductAttributes: domain.ProductAttributes{
			Name:        "TESTE",
			Quantity:    229,
			CodeValue:   "68788-9100",
			IsPublished: true,
			Expiration:  "2022-03-06",
			Price:       30.75,
			WarehouseId: 1,
		},
	}

	//when
	product, err := repo.GetById(2)
	if err != nil {
		t.Fatalf("GetById error: %v", err)
	}

	//then
	assert.NoError(t, err)
	assert.Equal(t, expected, product)
}
