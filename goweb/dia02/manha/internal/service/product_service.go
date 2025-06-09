package service

import (
	"errors"

	"google.com/bgw7-esther/first-server/internal/domain"
	"google.com/bgw7-esther/first-server/internal/repository"
)

type ProductService struct {
	Repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) *ProductService {
	return &ProductService{Repo: r}
}

func (s *ProductService) CreateProduct(product *domain.Product) error {

	if existing, _ := s.Repo.FindByCode(product.CodeValue); existing != nil {
		return errors.New("code_value must be unique")
	}

	return s.Repo.Create(product)
}

func (s *ProductService) GetAllProducts() ([]*domain.Product, error) {
	return s.Repo.GetAll()

}

func (s *ProductService) GetProductByID(id int) (*domain.Product, error) {
	return s.Repo.FindByID(id)
}

func (s *ProductService) GetProductsByPrice(minPrice float64) ([]*domain.Product, error) {
	return s.Repo.SearchByPrice(minPrice)
}
