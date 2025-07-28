package invoice_test

import (
	"app/internal/domain"
	"app/internal/invoice"
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

func TestFindAll_Invoice(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := invoice.NewInvoicesMySQL(db)

	// given
	_, err := db.Exec(`INSERT INTO invoices (id, datetime, total, customer_id) VALUES (?, ?, ?, ?)`,
		10, "2023-01-01 10:00:00", 10.50, 1)

	if err != nil {
		t.Fatalf("insert invoice: %v", err)
	}

	expected := []domain.Invoice{
		{
			Id: 10, InvoiceAttributes: domain.InvoiceAttributes{
				Datetime:   "2023-01-01 10:00:00",
				Total:      10.50,
				CustomerId: 1,
			},
		},
	}

	invoice, err := repo.FindAll()

	if err != nil {
		t.Fatalf("GetAll error: %v", err)
	}
	//then
	assert.Contains(t, invoice, expected[0])
}

func TestCreate_Invoice(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := invoice.NewInvoicesMySQL(db)

	// given
	input := domain.Invoice{
		Id: 10, InvoiceAttributes: domain.InvoiceAttributes{
			Datetime:   "2023-01-01 10:00:00",
			Total:      10.50,
			CustomerId: 1,
		},
	}

	//when
	err := repo.Save(&input)

	//then
	assert.NoError(t, err)

}

func TestUpdateInvoiceTotals(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := invoice.NewInvoicesMySQL(db)

	// when
	err := repo.UpdateAllTotal()

	// then
	assert.NoError(t, err)
}
