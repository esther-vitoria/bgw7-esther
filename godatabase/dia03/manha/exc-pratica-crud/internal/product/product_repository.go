package product

import (
	"database/sql"
	"fmt"

	"github.com/bgw7/exc-pratica-crud/internal/domain"
	"github.com/bgw7/exc-pratica-crud/pkg/apperrors"
)

type mysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(db *sql.DB) domain.ProductRepository {
	return &mysqlRepository{
		db: db,
	}
}

func (database *mysqlRepository) GetAll() (productsList []domain.Product, err error) {
	rows, err := database.db.Query(GetQuery)
	if err != nil {
		err = apperrors.ErrQueryDB
	}

	defer rows.Close()

	for rows.Next() {
		var product domain.Product
		err = rows.Scan(
			&product.Id,
			&product.ProductAttributes.Name,
			&product.ProductAttributes.Quantity,
			&product.ProductAttributes.CodeValue,
			&product.ProductAttributes.IsPublished,
			&product.ProductAttributes.Expiration,
			&product.ProductAttributes.Price,
		)

		if err != nil {
			fmt.Println(err)
			err = apperrors.ErrScanDB
			return
		}
		productsList = append(productsList, product)
	}

	if len(productsList) == 0 {
		err = apperrors.ErrEmptyList
		return
	}

	return
}

func (database *mysqlRepository) GetById(id int) (product domain.Product, err error) {

	row := database.db.QueryRow(GetByIdQuery, id)

	err = row.Scan(
		&product.Id,
		&product.ProductAttributes.Name,
		&product.ProductAttributes.Quantity,
		&product.ProductAttributes.CodeValue,
		&product.ProductAttributes.IsPublished,
		&product.ProductAttributes.Expiration,
		&product.ProductAttributes.Price,
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

func (database *mysqlRepository) Create(productAttributes domain.ProductAttributes) (product domain.Product, err error) {
	result, errExecDB := database.db.Exec(InsertQuery,
		(productAttributes).Name,
		(productAttributes).Quantity,
		(productAttributes).CodeValue,
		(productAttributes).IsPublished,
		(productAttributes).Expiration,
		(productAttributes).Price)

	if errExecDB != nil {
		err = apperrors.ErrQueryDB
		return
	}

	lastID, err := result.LastInsertId()

	if err != nil {
		err = apperrors.ErrQueryDB
		return
	}

	product.Id = int(lastID)
	product.ProductAttributes = productAttributes
	return
}

func (database *mysqlRepository) Update(id int, productAttributes domain.ProductAttributes) (product domain.Product, err error) {
	result, errExecDB := database.db.Exec(UpdateQuery,
		(productAttributes).Name,
		(productAttributes).Quantity,
		(productAttributes).CodeValue,
		(productAttributes).IsPublished,
		(productAttributes).Expiration,
		(productAttributes).Price,
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

	product.Id = id
	product.ProductAttributes = productAttributes
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
