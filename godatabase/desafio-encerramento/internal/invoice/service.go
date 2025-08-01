package invoice

import (
	"app/internal/domain"
)

// NewInvoicesDefault creates new default service for invoice entity.
func NewInvoicesDefault(rp domain.RepositoryInvoice) *InvoicesDefault {
	return &InvoicesDefault{rp}
}

// InvoicesDefault is the default service implementation for invoice entity.
type InvoicesDefault struct {
	// rp is the repository for invoice entity.
	rp domain.RepositoryInvoice
}

// FindAll returns all invoices.
func (s *InvoicesDefault) FindAll() (i []domain.Invoice, err error) {
	i, err = s.rp.FindAll()
	return
}

// Save saves the invoice.
func (s *InvoicesDefault) Save(i *domain.Invoice) (err error) {
	err = s.rp.Save(i)
	return
}

// UpdateAllTotal updates all invoices total.
func (s *InvoicesDefault) UpdateAllTotal() (err error) {
	err = s.rp.UpdateAllTotal()
	return
}

// ServiceInvoice is the interface that wraps the basic methods that an invoice service should implement.
type ServiceInvoice interface {
	// FindAll returns all invoices
	FindAll() (i []domain.Invoice, err error)
	// Save saves an invoice
	Save(i *domain.Invoice) (err error)
	// UpdateAllTotal updates all invoices total
	UpdateAllTotal() (err error)
}
