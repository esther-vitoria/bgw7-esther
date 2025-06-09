package repository

import (
	"errors"

	"google.com/bgw7-esther/first-server/internal/domain"
)

// Struct para armazenar os produtos
type ProductRepository struct {
	Storage map[int]*domain.Product
}

// Armazena um novo produto
func NewProductRepository(storage map[int]*domain.Product) *ProductRepository {
	return &ProductRepository{Storage: storage}
}

// Cria um Slice e trás todos os produtos
func (r *ProductRepository) GetAll() ([]*domain.Product, error) {
	var products []*domain.Product
	for _, p := range r.Storage {
		products = append(products, p)
	}
	return products, nil
}

// Verifica se o id existe, senão cria o produto no slice
func (r *ProductRepository) Create(product *domain.Product) error {
	if _, exist := r.Storage[product.Id]; exist {
		return errors.New("id already exists")
	}
	r.Storage[product.Id] = product
	return nil
}

// Verifica se o code value ja existe na base e retorna
func (r *ProductRepository) FindByCode(CodeValue string) (*domain.Product, error) {
	for _, p := range r.Storage {
		if p.CodeValue == CodeValue {
			return p, nil
		}
	}
	return nil, errors.New("not found")
}

func (r *ProductRepository) SearchByPrice(Price float64) ([]*domain.Product, error) {
	var result []*domain.Product
	for _, p := range r.Storage {
		if p.Price > Price {
			result = append(result, p)
		}
	}
	return result, nil
}

// Verifica os IDs e trás o id enviado no parametro
func (r *ProductRepository) FindByID(id int) (*domain.Product, error) {
	if p, ok := r.Storage[id]; ok {
		return p, nil
	}
	return nil, errors.New("id not found")
}

// Gera o novo ID a partir de uma checagem na base
func (r *ProductRepository) GetNewID() int {
	maxID := 0
	for id := range r.Storage {
		if id > maxID {
			maxID = id
		}
	}

	return maxID + 1

}
