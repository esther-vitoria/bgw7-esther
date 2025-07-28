package domain

// ProductAttributes is the struct that represents the attributes of a product.
type ProductAttributes struct {
	// Description is the description of the product.
	Description string
	// Price is the price of the product.
	Price float64
}

// Product is the struct that represents a product.
type Product struct {
	// Id is the unique identifier of the product.
	Id int
	// ProductAttributes is the attributes of the product.
	ProductAttributes
}

// ProductAmountSold is the struct that represents the total amount sold by product.
type ProductAmountSold struct {
	// Description is the description of the product.
	Description string
	// Total is the total amount sold by product.
	Total float64
}

// RepositoryProduct is the interface that wraps the basic methods that a product repository must have.
type RepositoryProduct interface {
	// FindAll returns all products saved in the database.
	FindAll() (p []Product, err error)
	// FindTopProductsByAmountSold returns the top sold products.
	FindTopProductsByAmountSold(limit int) (p []ProductAmountSold, err error)
	// Save saves a product into the database.
	Save(p *Product) (err error)
}
