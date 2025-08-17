package service

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func NewProductServiceMock() *ServiceMock {
	return &ServiceMock{}
}

func (s *ServiceMock) GetProducts(query internal.ProductQuery) (products map[int]internal.Product, err error) {
	args := s.Called(query)
	return args.Get(0).(map[int]internal.Product), args.Error(1)
}
