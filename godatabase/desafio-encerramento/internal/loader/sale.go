package loader

import (
	"app/internal/domain"
	"encoding/json"
	"os"
)

func NewSalesJSON(file *os.File) *SalesJSON {
	return &SalesJSON{file: file}
}

type SalesJSON struct {
	file *os.File
}

type SaleJSON struct {
	Id        int `json:"id"`
	Quantity  int `json:"quantity"`
	ProductId int `json:"product_id"`
	InvoiceId int `json:"invoice_id"`
}

func (l *SalesJSON) Load() (s []domain.Sale, err error) {
	var ss []SaleJSON
	err = json.NewDecoder(l.file).Decode(&ss)
	if err != nil {
		return
	}

	for _, v := range ss {
		s = append(s, domain.Sale{
			Id: v.Id,
			SaleAttributes: domain.SaleAttributes{
				Quantity:  v.Quantity,
				ProductId: v.ProductId,
				InvoiceId: v.InvoiceId,
			},
		})
	}

	return
}
