package product

import "github.com/bgw7/exc-pratica-crud/internal/domain"

type service struct {
	repository domain.ProductRepository
}

type Service interface {
	GetAll() (products []domain.Product, err error)
	Create(productAttributes domain.ProductAttributes) (product domain.Product, err error)
	Update(id int, productAttributes domain.ProductAttributes) (product domain.Product, err error)
	GetById(id int) (product domain.Product, err error)
	Delete(id int) (err error)
}

func NewService(repository domain.ProductRepository) Service {
	return &service{repository}
}

func (s *service) GetAll() (products []domain.Product, err error) {
	return s.repository.GetAll()
}

func (s *service) GetById(id int) (product domain.Product, err error) {
	return s.repository.GetById(id)
}

func (s *service) Create(productAttributes domain.ProductAttributes) (product domain.Product, err error) {
	return s.repository.Create(productAttributes)
}

func (s *service) Update(id int, productAttributes domain.ProductAttributes) (product domain.Product, err error) {
	return s.repository.Update(id, productAttributes)
}

func (s *service) Delete(id int) (err error) {
	return s.repository.Delete(id)
}
