package sale_test

import (
	"app/internal/domain"
	"app/internal/sale"
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

func TestFindAll_Sales(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := sale.NewSalesMySQL(db)

	// given
	_, err := db.Exec(`INSERT INTO sales (id, quantity, product_id, invoice_id) VALUES (?,?, ?, ?)`,
		5, 5, 5, 5)

	if err != nil {
		t.Fatalf("insert sale: %v", err)
	}

	expected := []domain.Sale{
		{
			Id: 5, SaleAttributes: domain.SaleAttributes{
				Quantity:  5,
				ProductId: 5,
				InvoiceId: 5,
			},
		},
	}

	sales, err := repo.FindAll()

	if err != nil {
		t.Fatalf("GetAll error: %v", err)
	}
	//then
	assert.Contains(t, sales, expected[0])
}

func TestCreate_Sales(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := sale.NewSalesMySQL(db)

	//given
	// given
	input := domain.Sale{
		Id: 6,
		SaleAttributes: domain.SaleAttributes{
			Quantity:  6,
			ProductId: 6,
			InvoiceId: 6,
		},
	}

	//when
	err := repo.Save(&input)

	//then
	assert.NoError(t, err)

}
