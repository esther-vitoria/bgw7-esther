package domain

// Entidade
type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// Interface para abstrair o repositorio
type ProductRepository interface {
	GetAll() ([]*Product, error)
	Create(product *Product)
	FindByID(id int) (*Product, error)
	SearchByPrice(price float64) (*Product, error)
}
