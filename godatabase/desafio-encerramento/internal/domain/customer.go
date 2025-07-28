package domain

// CustomerAttributes is the struct that represents the attributes of a customer.
type CustomerAttributes struct {
	// FirstName is the first name of the customer.
	FirstName string
	// LastName is the last name of the customer.
	LastName string
	// Condition is the condition of the customer.
	Condition int
}

// Customer is the struct that represents a customer.
type Customer struct {
	// Id is the unique identifier of the customer.
	Id int
	// CustomerAttributes is the attributes of the customer.
	CustomerAttributes
}

// CustomerInvoicesByCondition is the struct that represents the total invoices by customer condition.
type CustomerInvoicesByCondition struct {
	// Condition is the condition of the customer.
	Condition string
	// Total is the total invoices by customer condition
	Total float64
}

// CustomerSpent is the struct that represents the total spent by customer.
type CustomerSpent struct {
	// FirstName is the first name of the customer.
	FirstName string
	// LastName is the last name of the customer.
	LastName string
	// Total is the total spent by customer.
	Total float64
}

// RepositoryCustomer is the interface that wraps the basic methods that a customer repository should implement.
type RepositoryCustomer interface {
	// FindAll returns all customers saved in the database.
	FindAll() (c []Customer, err error)
	// FindTopActiveCustomersByAmountSpent returns the top active customers by amount spent.
	FindTopActiveCustomersByAmountSpent(limit int) (c []CustomerSpent, err error)
	// FindInvoicesByCondition returns the total invoices by customer condition.
	FindInvoicesByCondition() (c []CustomerInvoicesByCondition, err error)
	// Save saves a customer into the database.
	Save(c *Customer) (err error)
}
