package service

import "app/internal"

type ProductService struct {
	rp internal.RepositoryProducts
}

func NewProductService(rp internal.RepositoryProducts) *ProductService {
	return &ProductService{rp: rp}
}

func (s *ProductService) GetProducts(query internal.ProductQuery) (products map[int]internal.Product, err error) {
	products, err = s.rp.SearchProducts(query)
	return
}

type Service interface {
	GetProducts(query internal.ProductQuery) (products map[int]internal.Product, err error)
}
