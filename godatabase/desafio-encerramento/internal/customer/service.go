package customer

import (
	"app/internal/domain"
)

// NewCustomersDefault creates new default service for customer entity.
func NewCustomersDefault(rp domain.RepositoryCustomer) *CustomersDefault {
	return &CustomersDefault{rp}
}

// CustomersDefault is the default service implementation for customer entity.
type CustomersDefault struct {
	// rp is the repository for customer entity.
	rp domain.RepositoryCustomer
}

// FindAll returns all customers.
func (s *CustomersDefault) FindAll() (c []domain.Customer, err error) {
	c, err = s.rp.FindAll()
	return
}

// FindTopActiveCustomersByAmountSpent returns the top active customers by amount spent.
func (s *CustomersDefault) FindTopActiveCustomersByAmountSpent(limit int) (c []domain.CustomerSpent, err error) {
	c, err = s.rp.FindTopActiveCustomersByAmountSpent(limit)
	return
}

// FindInvoicesByCondition returns the total invoices by customer condition.
func (s *CustomersDefault) FindInvoicesByCondition() (c []domain.CustomerInvoicesByCondition, err error) {
	c, err = s.rp.FindInvoicesByCondition()
	return
}

// Save saves the customer.
func (s *CustomersDefault) Save(c *domain.Customer) (err error) {
	err = s.rp.Save(c)
	return
}

// ServiceCustomer is the interface that wraps the basic methods that a customer service should implement.
type ServiceCustomer interface {
	// FindAll returns all customers
	FindAll() (c []domain.Customer, err error)
	// FindTopActiveCustomersByAmountSpent returns the top active customers by amount spent
	FindTopActiveCustomersByAmountSpent(limit int) (c []domain.CustomerSpent, err error)
	// FindInvoicesByCondition returns the total invoices by customer condition
	FindInvoicesByCondition() (c []domain.CustomerInvoicesByCondition, err error)
	// Save saves a customer
	Save(c *domain.Customer) (err error)
}
