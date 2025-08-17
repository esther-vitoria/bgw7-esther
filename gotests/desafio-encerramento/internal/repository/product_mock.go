package repository

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func NewProductServiceMock() *RepositoryMock {
	return &RepositoryMock{}
}

func (s *RepositoryMock) GetProducts(query internal.ProductQuery) (products map[int]internal.Product, err error) {
	args := s.Called(query)
	return args.Get(0).(map[int]internal.Product), args.Error(1)
}
