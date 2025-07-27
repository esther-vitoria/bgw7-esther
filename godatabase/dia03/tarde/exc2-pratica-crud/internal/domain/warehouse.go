package domain

type WarehouseAttributes struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
	Capacity  int    `json:"capacity"`
}

type Warehouse struct {
	Id int `json:"id"`
	WarehouseAttributes
}

type Report struct {
	Name         string `json:"name"`
	ProductCount int    `json:"product_count"`
}

type WarehouseRepository interface {
	GetAll() (warehouses []Warehouse, err error)
	GetById(id int) (warehouse Warehouse, err error)
	Create(warehouseAttributesNew WarehouseAttributes) (warehouse Warehouse, err error)
	Update(id int, warehouseAttributesNew WarehouseAttributes) (warehouse Warehouse, err error)
	Delete(id int) (err error)
	GetReport(id int) (report Report, err error)
}
