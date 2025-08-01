package product

import (
	"app/internal/domain"
)

// NewProductsDefault creates new default service for Product entity.
func NewProductsDefault(rp domain.RepositoryProduct) *ProducstDefault {
	return &ProducstDefault{rp}
}

// ProducstDefault is the default service implementation for Product entity.
type ProducstDefault struct {
	// rp is the repository for Product entity.
	rp domain.RepositoryProduct
}

// FindAll returns all Products.
func (s *ProducstDefault) FindAll() (p []domain.Product, err error) {
	p, err = s.rp.FindAll()
	return
}

// FindTopProductsByAmountSold returns the top Products by amount sold.
func (s *ProducstDefault) FindTopProductsByAmountSold(limit int) (p []domain.ProductAmountSold, err error) {
	p, err = s.rp.FindTopProductsByAmountSold(limit)
	return
}

// Save saves the Product.
func (s *ProducstDefault) Save(p *domain.Product) (err error) {
	err = s.rp.Save(p)
	return
}

type ServiceProduct interface {
	// FindAll returns all Products.
	FindAll() (p []domain.Product, err error)
	// FindTopProductsByAmountSold returns the top Products by amount sold.
	FindTopProductsByAmountSold(limit int) (p []domain.ProductAmountSold, err error)
	// Save saves a Product.
	Save(p *domain.Product) (err error)
}
