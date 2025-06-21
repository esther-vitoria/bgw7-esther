package service

// Service: contém a lógica comercial da API, lida com conexões externas.

import (
	"google.com/bgw7-esther/first-server/internal/repository"
	"google.com/bgw7-esther/first-server/internal/storage"
)

type ProductService struct {
	Repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) *ProductService {
	return &ProductService{Repo: r}
}

func (s *ProductService) CreateProduct(product *storage.ProductJSON) (*storage.ProductJSON, error) {
	return s.Repo.Create(product)
}

func (s *ProductService) GetProductByID(id int) (*storage.ProductJSON, error) {
	return s.Repo.FindByID(id)
}

func (s *ProductService) SearchProducts(price float64) ([]*storage.ProductJSON, error) {
	return s.Repo.SearchProducts(price)
}

func (s *ProductService) GetAll() (map[int]*storage.ProductJSON, error) {
	return s.Repo.FindAll()
}

func (s *ProductService) UpdateProductFull(product *storage.ProductJSON) (*storage.ProductJSON, error) {
	return s.Repo.UpdateProductFull(product)
}

func (s *ProductService) PatchProduct(id int, name *string, quantity *int, codeValue *string, expiration *string, is_published *bool, price *float64) (*storage.ProductJSON, error) {
	return s.Repo.PatchProduct(id, name, quantity, codeValue, expiration, is_published, price)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.Repo.DeleteProduct(id)
}

func (s *ProductService) LoadProducts(path string) []*storage.ProductJSON {
	return s.LoadProducts(path)
}
