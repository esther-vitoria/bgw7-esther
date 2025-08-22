package repository

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

// ProductRepositoryMock is a mock implementation of RepositoryProducts interface
type ProductRepositoryMock struct {
	mock.Mock
}

// NewProductRepositoryMock creates a new instance of ProductRepositoryMock
func NewProductRepositoryMock() *ProductRepositoryMock {
	return &ProductRepositoryMock{}
}

// SearchProducts mocks the SearchProducts method
func (m *ProductRepositoryMock) SearchProducts(query internal.ProductQuery) (map[int]internal.Product, error) {
	args := m.Called(query)
	return args.Get(0).(map[int]internal.Product), args.Error(1)
}
