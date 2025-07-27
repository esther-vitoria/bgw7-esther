package warehouse

import "github.com/bgw7/exc-pratica-crud/internal/domain"

type service struct {
	repository domain.WarehouseRepository
}

type Service interface {
	GetAll() (warehouseList []domain.Warehouse, err error)
	GetById(id int) (warehouse domain.Warehouse, err error)
	Create(warehouseAttributesNew domain.WarehouseAttributes) (warehouse domain.Warehouse, err error)
	Update(id int, warehouseAttributesNew domain.WarehouseAttributes) (warehouse domain.Warehouse, err error)
	Delete(id int) (err error)
	GetReport(id int) (report domain.Report, err error)
}

func NewService(repository domain.WarehouseRepository) Service {
	return &service{repository}
}

func (s *service) GetAll() (warehouseList []domain.Warehouse, err error) {
	return s.repository.GetAll()
}

func (s *service) GetById(id int) (warehouse domain.Warehouse, err error) {
	return s.repository.GetById(id)
}

func (s *service) Create(warehouseAttributesNew domain.WarehouseAttributes) (warehouse domain.Warehouse, err error) {
	return s.repository.Create(warehouseAttributesNew)
}

func (s *service) Update(id int, warehouseAttributesNew domain.WarehouseAttributes) (warehouse domain.Warehouse, err error) {
	return s.repository.Update(id, warehouseAttributesNew)
}

func (s *service) Delete(id int) (err error) {
	return s.repository.Delete(id)
}

func (s *service) GetReport(id int) (report domain.Report, err error) {
	return s.repository.GetReport(id)
}
