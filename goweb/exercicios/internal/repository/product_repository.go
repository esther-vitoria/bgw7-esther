package repository

// Repository: é responsável pelas operações diretas com o banco de dados.

import (
	"errors"

	"google.com/bgw7-esther/first-server/internal/storage"
	"google.com/bgw7-esther/first-server/internal/utils"
)

// Struct para armazenar os produtos
type ProductRepository struct {
	Storage map[int]*storage.ProductJSON
}

// Armazena um novo produto
func NewProductRepository(storage map[int]*storage.ProductJSON) *ProductRepository {
	return &ProductRepository{Storage: storage}
}

func (r *ProductRepository) Create(product *storage.ProductJSON) (*storage.ProductJSON, error) {
	// Verifica se o id existe
	if _, exist := r.Storage[product.Id]; exist {
		return nil, errors.New("id already exists")
	}

	// Verifica se o code_value existe
	for _, p := range r.Storage {
		if p.CodeValue == product.CodeValue {
			return nil, errors.New("code_value already exists")
		}
	}

	//Valida se a data foi inserida no formato correto
	if !utils.IsValidDate(product.Expiration) {
		return &storage.ProductJSON{}, errors.New("invalid date format")
	}
	//Pega um novo id
	newID := r.GetNewID()
	product.Id = newID

	// adiciona um novo produto na slice
	r.Storage[product.Id] = product

	return product, nil
}

func (r *ProductRepository) UpdateProductFull(product *storage.ProductJSON) (*storage.ProductJSON, error) {

	//Valida se a data foi inserida no formato correto
	if !utils.IsValidDate(product.Expiration) {
		return &storage.ProductJSON{}, errors.New("invalid date format")
	}
	// adiciona um novo produto na slice
	r.Storage[product.Id] = product

	return product, nil
}

func (r *ProductRepository) FindAll() (v map[int]*storage.ProductJSON, err error) {
	v = make(map[int]*storage.ProductJSON)

	// copy db
	for key, value := range r.Storage {
		v[key] = value
	}
	return
}

// Recebe um float e trás os produtos que possuem valor maior que o preço inserido
func (r *ProductRepository) SearchProducts(price float64) ([]*storage.ProductJSON, error) {
	var result []*storage.ProductJSON

	for _, p := range r.Storage {
		if p.Price > price {
			result = append(result, p)
		}

		if result == nil {
			return nil, errors.New("nothing found within this criteria")
		}
	}
	return result, nil
}

func (r *ProductRepository) PatchProduct(id int, name *string, quantity *int, codeValue *string, expiration *string, is_published *bool, price *float64) (*storage.ProductJSON, error) {

	p, exists := r.Storage[id]
	if !exists {
		return nil, errors.New("id not found")
	}

	// Só atualiza se o campo foi informado
	if name != nil {
		p.Name = *name
	}
	if price != nil {
		p.Price = *price
	}
	if quantity != nil {
		p.Quantity = *quantity
	}
	if expiration != nil {
		if !utils.IsValidDate(*expiration) {
			return nil, errors.New("invalid date format")
		}
		p.Expiration = *expiration
	}
	if codeValue != nil {
		// Verifica unicidade do code value salvo
		for _, prod := range r.Storage {
			if prod.CodeValue == *codeValue && prod.Id != id {
				return nil, errors.New("code_value already exists")
			}
		}
		p.CodeValue = *codeValue
	}

	return p, nil
}

// Deleta um produto da slice baseado no id
func (r *ProductRepository) DeleteProduct(id int) error {
	if _, exists := r.Storage[id]; !exists {
		return errors.New("id not found")
	}
	delete(r.Storage, id)
	return nil
}

// Verifica os IDs e trás o id enviado no parametro
func (r *ProductRepository) FindByID(id int) (*storage.ProductJSON, error) {
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
