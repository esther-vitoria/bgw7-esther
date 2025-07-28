package sale

import (
	"app/internal/domain"
)

// NewSalesDefault creates new default service for sale entity.
func NewSalesDefault(rp domain.RepositorySale) *SalesDefault {
	return &SalesDefault{rp}
}

// SalesDefault is the default service implementation for sale entity.
type SalesDefault struct {
	// rp is the repository for sale entity.
	rp domain.RepositorySale
}

// FindAll returns all sales.
func (sv *SalesDefault) FindAll() (s []domain.Sale, err error) {
	s, err = sv.rp.FindAll()
	return
}

// Save saves the sale.
func (sv *SalesDefault) Save(s *domain.Sale) (err error) {
	err = sv.rp.Save(s)
	return
}

// ServiceSale is the interface that wraps the basic ServiceSale methods.
type ServiceSale interface {
	// FindAll returns all sales.
	FindAll() (s []domain.Sale, err error)
	// Save saves a sale.
	Save(s *domain.Sale) (err error)
}
