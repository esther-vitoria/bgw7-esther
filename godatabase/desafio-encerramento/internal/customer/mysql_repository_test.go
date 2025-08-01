package customer_test

import (
	"app/internal/customer"
	"app/internal/domain"
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

func TestFindAll_Customer(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := customer.NewCustomersMySQL(db)

	// given
	_, err := db.Exec("INSERT INTO customers (id, first_name, last_name, `condition`) VALUES (?, ?, ?, ?)",
		10, "Esther", "Vitoria", 1)

	if err != nil {
		t.Fatalf("insert customer: %v", err)
	}

	expected := []domain.Customer{
		{
			Id: 10, CustomerAttributes: domain.CustomerAttributes{
				FirstName: "Esther",
				LastName:  "Vitoria",
				Condition: 1,
			},
		},
	}

	customer, err := repo.FindAll()

	if err != nil {
		t.Fatalf("GetAll error: %v", err)
	}
	//then
	assert.Contains(t, customer, expected[0])
}

func TestCreate_Invoice(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := customer.NewCustomersMySQL(db)

	// given
	input := domain.Customer{
		Id: 10, CustomerAttributes: domain.CustomerAttributes{
			FirstName: "Esther",
			LastName:  "Vitoria",
			Condition: 1,
		},
	}

	//when
	err := repo.Save(&input)

	//then
	assert.NoError(t, err)

}

func TestFindTopActiveCustomersByAmountSpent(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := customer.NewCustomersMySQL(db)

	// when
	_, err := repo.FindTopActiveCustomersByAmountSpent(5)

	// then
	assert.NoError(t, err)
}

func TestFindInvoicesByCondition(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := customer.NewCustomersMySQL(db)

	// when
	_, err := repo.FindInvoicesByCondition()

	// then
	assert.NoError(t, err)
}
