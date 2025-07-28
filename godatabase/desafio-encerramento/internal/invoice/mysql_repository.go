package invoice

import (
	"database/sql"
	"fmt"

	"app/internal/domain"
)

// NewInvoicesMySQL creates new mysql repository for invoice entity.
func NewInvoicesMySQL(db *sql.DB) *InvoicesMySQL {
	return &InvoicesMySQL{db}
}

// InvoicesMySQL is the MySQL repository implementation for invoice entity.
type InvoicesMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all invoices from the database.
func (r *InvoicesMySQL) FindAll() (i []domain.Invoice, err error) {
	rows, err := r.db.Query(GetQuery)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var inv domain.Invoice
		var total sql.NullFloat64

		err := rows.Scan(&inv.Id, &inv.Datetime, &total, &inv.CustomerId)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		if total.Valid {
			inv.Total = total.Float64
		} else {
			inv.Total = 0
		}
		i = append(i, inv)
	}
	err = rows.Err()
	return
}

// Save saves the invoice into the database.
func (r *InvoicesMySQL) Save(i *domain.Invoice) (err error) {
	// execute the query
	res, err := r.db.Exec(InsertQuery, (*i).Datetime, (*i).Total, (*i).CustomerId)
	if err != nil {
		return err
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*i).Id = int(id)

	return
}

// UpdateAllTotal updates all invoices total
func (r *InvoicesMySQL) UpdateAllTotal() (err error) {
	// execute the query
	_, err = r.db.Exec(UpdateInvoiceTotals)
	return
}
