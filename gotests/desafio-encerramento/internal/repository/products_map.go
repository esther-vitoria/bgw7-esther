package repository

import (
	"app/internal"
	"app/loader"
	"errors"
	"os"
)

// NewProductsJSON returns a new ProductsJSON repository.
func NewProductsJSON(filePath string) (*ProductsJSON, error) {
	// open the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	// create the loader
	loader := loader.NewProductsJSON(file)

	// load the products data
	products, err := loader.Load()
	if err != nil {
		file.Close()
		return nil, err
	}

	// convert slice to map for easier search operations
	db := make(map[int]internal.Product)
	for _, product := range products {
		db[product.Id] = product
	}

	return &ProductsJSON{
		file: file,
		db:   db,
	}, nil
}

// ProductsJSON is a struct that implements the RepositoryProducts interface using JSON as data source.
type ProductsJSON struct {
	// file is the JSON file reference
	file *os.File
	// db is the in-memory map of products loaded from JSON
	db map[int]internal.Product
}

// SearchProducts returns a list of products that match the query.
func (r *ProductsJSON) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {
	p = make(map[int]internal.Product)

	// search the products
	for k, v := range r.db {
		// check if each query field is set
		if query.Id > 0 && query.Id != k {
			continue
		}

		// add the product to the result
		p[k] = v

		if len(p) == 0 {
			return nil, errors.New("product not found")
		}
	}

	return
}

// Close closes the file handle
func (r *ProductsJSON) Close() error {
	if r.file != nil {
		return r.file.Close()
	}
	return nil
}
