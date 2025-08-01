package domain

// InvoiceAttributes is the struct that represents the attributes of an invoice.
type InvoiceAttributes struct {
	// Datetime is the datetime of the invoice.
	Datetime string
	// Total is the total of the invoice.
	Total float64
	// CustomerId is the customer id of the invoice.
	CustomerId int
}

// Invoice is the struct that represents an invoice.
type Invoice struct {
	// Id is the id of the invoice.
	Id int
	// InvoiceAttributes is the attributes of the invoice.
	InvoiceAttributes
}

// RepositoryInvoice is the interface that wraps the basic methods that an invoice repository should implement.
type RepositoryInvoice interface {
	// FindAll returns all invoices
	FindAll() (i []Invoice, err error)
	// Save saves an invoice
	Save(i *Invoice) (err error)
	// UpdateAllTotal updates all invoices total
	UpdateAllTotal() (err error)
}
