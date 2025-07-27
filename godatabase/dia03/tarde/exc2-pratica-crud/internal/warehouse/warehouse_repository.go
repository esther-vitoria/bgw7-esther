package warehouse

import (
	"database/sql"
	"fmt"

	"github.com/bgw7/exc-pratica-crud/internal/domain"
	"github.com/bgw7/exc-pratica-crud/pkg/apperrors"
)

type mysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(db *sql.DB) domain.WarehouseRepository {
	return &mysqlRepository{
		db: db,
	}
}

func (database *mysqlRepository) GetAll() (warehouseList []domain.Warehouse, err error) {
	rows, err := database.db.Query(GetQuery)
	if err != nil {
		err = apperrors.ErrQueryDB
	}

	defer rows.Close()

	for rows.Next() {
		var warehouse domain.Warehouse
		err = rows.Scan(
			&warehouse.Id,
			&warehouse.WarehouseAttributes.Name,
			&warehouse.WarehouseAttributes.Address,
			&warehouse.WarehouseAttributes.Telephone,
			&warehouse.WarehouseAttributes.Capacity,
		)

		if err != nil {
			err = apperrors.ErrScanDB
			return
		}
		warehouseList = append(warehouseList, warehouse)
	}

	if len(warehouseList) == 0 {
		err = apperrors.ErrEmptyList
		return
	}

	return
}

func (database *mysqlRepository) GetById(id int) (warehouse domain.Warehouse, err error) {

	row := database.db.QueryRow(GetByIdQuery, id)

	err = row.Scan(
		&warehouse.Id,
		&warehouse.WarehouseAttributes.Name,
		&warehouse.WarehouseAttributes.Address,
		&warehouse.WarehouseAttributes.Telephone,
		&warehouse.WarehouseAttributes.Capacity,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			err = apperrors.ErrNotFound
		} else {
			err = apperrors.ErrScanDB
		}
		return
	}

	return
}

func (database *mysqlRepository) Create(warehouseAttributesNew domain.WarehouseAttributes) (warehouse domain.Warehouse, err error) {
	result, errExecDB := database.db.Exec(InsertQuery,
		(warehouseAttributesNew).Name,
		(warehouseAttributesNew).Address,
		(warehouseAttributesNew).Telephone,
		(warehouseAttributesNew).Capacity)

	if errExecDB != nil {
		err = apperrors.ErrQueryDB
		return
	}

	lastID, err := result.LastInsertId()

	if err != nil {
		err = apperrors.ErrQueryDB
		return
	}

	warehouse.Id = int(lastID)
	warehouse.WarehouseAttributes = warehouseAttributesNew
	return
}

func (database *mysqlRepository) Update(id int, warehouseAttributesNew domain.WarehouseAttributes) (warehouse domain.Warehouse, err error) {
	result, errExecDB := database.db.Exec(UpdateQuery,
		(warehouseAttributesNew).Name,
		(warehouseAttributesNew).Address,
		(warehouseAttributesNew).Telephone,
		(warehouseAttributesNew).Capacity,
		id,
	)

	if errExecDB != nil {
		fmt.Println(errExecDB)
		err = apperrors.ErrQueryDB
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		err = apperrors.ErrNotFound
		return
	}

	warehouse.Id = id
	warehouse.WarehouseAttributes = warehouseAttributesNew
	return
}

func (database *mysqlRepository) Delete(id int) (err error) {
	result, err := database.db.Exec(DeleteQuery, id)

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		err = apperrors.ErrNotFound
		return
	}
	return
}

func (database *mysqlRepository) GetReport(id int) (report domain.Report, err error) {
	row := database.db.QueryRow(GetReport, id)

	err = row.Scan(
		&report.Name,
		&report.ProductCount,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			err = apperrors.ErrNotFound
		} else {
			err = apperrors.ErrScanDB
		}
		return
	}

	return
}
