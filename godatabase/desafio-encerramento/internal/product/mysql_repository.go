package product

import (
	"database/sql"

	"app/internal/domain"
)

// NewProductsMySQL creates new mysql repository for product entity.
func NewProductsMySQL(db *sql.DB) *ProductsMySQL {
	return &ProductsMySQL{db}
}

// ProductsMySQL is the MySQL repository implementation for product entity.
type ProductsMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all products from the database.
func (r *ProductsMySQL) FindAll() (p []domain.Product, err error) {
	// execute the query
	rows, err := r.db.Query(GetQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var pr domain.Product
		// scan the row into the product
		err := rows.Scan(&pr.Id, &pr.Description, &pr.Price)
		if err != nil {
			return nil, err
		}
		// append the product to the slice
		p = append(p, pr)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// FindTopProductsByAmountSold returns the top products by amount sold.
func (r *ProductsMySQL) FindTopProductsByAmountSold(limit int) (p []domain.ProductAmountSold, err error) {
	// execute the query
	rows, err := r.db.Query(GetTop5SoldProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var pr domain.ProductAmountSold
		// scan the row into the product
		err := rows.Scan(&pr.Description, &pr.Total)
		if err != nil {
			return nil, err
		}
		// append the product to the slice
		p = append(p, pr)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the product into the database.
func (r *ProductsMySQL) Save(p *domain.Product) (err error) {
	// execute the query
	res, err := r.db.Exec(InsertQuery, (*p).Description, (*p).Price)
	if err != nil {
		return err
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*p).Id = int(id)

	return
}
