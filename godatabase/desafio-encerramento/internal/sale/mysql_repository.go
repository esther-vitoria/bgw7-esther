package sale

import (
	"database/sql"

	"app/internal/domain"
)

// NewSalesMySQL creates new mysql repository for sale entity.
func NewSalesMySQL(db *sql.DB) *SalesMySQL {
	return &SalesMySQL{db}
}

// SalesMySQL is the MySQL repository implementation for sale entity.
type SalesMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all sales from the database.
func (r *SalesMySQL) FindAll() (s []domain.Sale, err error) {
	// execute the query
	rows, err := r.db.Query(GetQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var sa domain.Sale
		// scan the row into the sale
		err := rows.Scan(&sa.Id, &sa.Quantity, &sa.ProductId, &sa.InvoiceId)
		if err != nil {
			return nil, err
		}
		// append the sale to the slice
		s = append(s, sa)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the sale into the database.
func (r *SalesMySQL) Save(s *domain.Sale) (err error) {
	// execute the query
	res, err := r.db.Exec(InsertQuery, (*s).Quantity, (*s).ProductId, (*s).InvoiceId)
	if err != nil {
		return err
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*s).Id = int(id)

	return
}
