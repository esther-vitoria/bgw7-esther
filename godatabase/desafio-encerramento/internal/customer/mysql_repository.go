package customer

import (
	"database/sql"
	"fmt"

	"app/internal/domain"
)

// NewCustomersMySQL creates new mysql repository for customer entity.
func NewCustomersMySQL(db *sql.DB) *CustomersMySQL {
	return &CustomersMySQL{db}
}

// CustomersMySQL is the MySQL repository implementation for customer entity.
type CustomersMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all customers from the database.
func (r *CustomersMySQL) FindAll() (c []domain.Customer, err error) {
	// execute the query
	rows, err := r.db.Query(GetQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var cs domain.Customer
		// scan the row into the customer
		err := rows.Scan(&cs.Id, &cs.FirstName, &cs.LastName, &cs.Condition)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		c = append(c, cs)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// FindTopActiveCustomersByAmountSpent returns the top active customers by amount spent.
func (r *CustomersMySQL) FindTopActiveCustomersByAmountSpent(limit int) (c []domain.CustomerSpent, err error) {
	// execute the query
	rows, err := r.db.Query(GetTop5ActiveCustomersByTotalSpent)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var cs domain.CustomerSpent
		// scan the row into the customer
		err := rows.Scan(&cs.FirstName, &cs.LastName, &cs.Total)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		c = append(c, cs)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// FindInvoicesByCondition returns the total invoices by customer condition.
func (r *CustomersMySQL) FindInvoicesByCondition() (c []domain.CustomerInvoicesByCondition, err error) {
	// execute the query
	rows, err := r.db.Query(GetTotalByCondition)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var cs domain.CustomerInvoicesByCondition
		// scan the row into the customer
		err := rows.Scan(&cs.Condition, &cs.Total)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		// append the customer to the slice
		c = append(c, cs)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the customer into the database.
func (r *CustomersMySQL) Save(c *domain.Customer) (err error) {
	// execute query
	res, err := r.db.Exec(InsertQuery,
		(*c).FirstName, (*c).LastName, (*c).Condition)
	if err != nil {
		return err
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*c).Id = int(id)

	return
}
