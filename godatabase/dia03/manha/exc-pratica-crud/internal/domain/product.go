package domain

// ProductAttributes is a struct that contains the attributes of a product
type ProductAttributes struct {
	// Name is the name of the product
	Name string `json:"name"`
	// Quantity is the quantity of the product
	Quantity int `json:"quantity"`
	// CodeValue is the code value of the product
	CodeValue string `json:"code_value"`
	// IsPublished is the published status of the product
	IsPublished bool `json:"is_published"`
	// Expiration
	Expiration string `json:"expiration"`
	// Price
	Price float64 `json:"price"`
}

// Product is a struct that contains the attributes of a product
type Product struct {
	// Id is the unique identifier of the product
	Id int `json:"id"`
	// ProductAttributes is the attributes of the product
	ProductAttributes
}

type ProductRepository interface {
	GetAll() (products []Product, err error)
	GetById(id int) (product Product, err error)
	Create(productAttributes ProductAttributes) (product Product, err error)
	Update(id int, productAttributes ProductAttributes) (product Product, err error)
	Delete(id int) (err error)
}
