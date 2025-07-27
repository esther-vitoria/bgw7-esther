package warehouse_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-txdb"
	"github.com/bgw7/exc-pratica-crud/internal/domain"
	"github.com/bgw7/exc-pratica-crud/internal/warehouse"
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

func TestGetAll_Warehouse(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := warehouse.NewMysqlRepository(db)

	// given
	_, err := db.Exec(`INSERT INTO warehouses(id, name, address, telephone, capacity)
					VALUES (?, ?, ?, ?, ?)`, 2, "TESTE", "TESTE", "(11) 3456-7890", 200)

	if err != nil {
		t.Fatalf("insert warehouses: %v", err)
	}

	expected := []domain.Warehouse{
		{
			Id: 1, WarehouseAttributes: domain.WarehouseAttributes{
				Name:      "Central Warehouse",
				Address:   "Rua das Flores, 123, Centro, Sao Paulo, SP",
				Telephone: "(11) 3456-7890",
				Capacity:  1000,
			},
		},
		{
			Id: 2, WarehouseAttributes: domain.WarehouseAttributes{
				Name:      "TESTE",
				Address:   "TESTE",
				Telephone: "(11) 3456-7890",
				Capacity:  200,
			},
		},
	}

	//when
	warehouse, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll error: %v", err)
	}
	//then
	assert.NoError(t, err)
	assert.Len(t, warehouse, 2)
	assert.Equal(t, expected, warehouse)
}

func TestGetById_Warehouse(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := warehouse.NewMysqlRepository(db)

	// given
	_, err := db.Exec(`INSERT INTO warehouses(id, name, address, telephone, capacity)
					VALUES (?, ?, ?, ?, ?)`, 2, "TESTE", "TESTE", "(11) 3456-7890", 200)

	if err != nil {
		t.Fatalf("insert warehouses: %v", err)
	}

	expected := domain.Warehouse{
		Id: 2, WarehouseAttributes: domain.WarehouseAttributes{
			Name:      "TESTE",
			Address:   "TESTE",
			Telephone: "(11) 3456-7890",
			Capacity:  200,
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

func TestCreate_Warehouse(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()

	repo := warehouse.NewMysqlRepository(db)

	// Dados para criação
	input := domain.WarehouseAttributes{
		Name:      "Armazém de Teste",
		Address:   "Rua dos Testes, 123",
		Telephone: "1111-2222",
		Capacity:  456,
	}

	// when: chama o método Create
	created, err := repo.Create(input)
	assert.NoError(t, err)

	// then: verifica se retornou objeto correto
	assert.NotZero(t, created.Id) // o id deve ser preenchido

	assert.Equal(t, input.Name, created.WarehouseAttributes.Name)
	assert.Equal(t, input.Address, created.WarehouseAttributes.Address)
	assert.Equal(t, input.Telephone, created.WarehouseAttributes.Telephone)
	assert.Equal(t, input.Capacity, created.WarehouseAttributes.Capacity)
}
