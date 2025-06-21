package domain

import "google.com/bgw7-esther/first-server/internal/storage"

// Domain/Application: orquestra e coordena as diferentes funcionalidades do nosso aplicativo orientado para seu uso final.

// Interface para abstrair o repositorio
type ProductRepository interface {
	Create(product *storage.ProductJSON)
	FindByID(id int) (*storage.ProductJSON, error)
	GetAll() (map[int]*storage.ProductJSON, error)
	SearchProducts(price float64) ([]*storage.ProductJSON, error)
	UpdateProductFull(product *storage.ProductJSON) (*storage.ProductJSON, error)
	DeleteProduct(id int) error
	PatchProduct(id int, name *string, quantity *int, codeValue *string, expiration *string, is_published *bool, price *float64) (*storage.ProductJSON, error)
	LoadProducts(path string) []*storage.ProductJSON
}
