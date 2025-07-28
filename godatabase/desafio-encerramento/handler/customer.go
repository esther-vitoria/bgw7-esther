package handler

import (
	"net/http"

	"app/internal/customer"
	"app/internal/domain"
	"app/pkg/web/request"
	"app/pkg/web/response"
)

// NewCustomersDefault returns a new CustomersDefault
func NewCustomersDefault(sv customer.ServiceCustomer) *CustomersDefault {
	return &CustomersDefault{sv: sv}
}

// CustomersDefault is a struct that returns the customer handlers
type CustomersDefault struct {
	// sv is the customer's service
	sv customer.ServiceCustomer
}

// CustomerJSON is a struct that represents a customer in JSON format
type CustomerJSON struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Condition int    `json:"condition"`
}

// GetAll returns all customers
func (h *CustomersDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		c, err := h.sv.FindAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error getting customers")
			return
		}

		// response
		// - serialize
		csJSON := make([]CustomerJSON, len(c))
		for ix, v := range c {
			csJSON[ix] = CustomerJSON{
				Id:        v.Id,
				FirstName: v.FirstName,
				LastName:  v.LastName,
				Condition: v.Condition,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "customers found",
			"data":    csJSON,
		})
	}
}

// CustomerSpentJSON is a struct that represents a customer spent in JSON format
type CustomerSpentJSON struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Total     float64 `json:"total"`
}

// GetTopActiveCustomersDefaultByAmountSpent returns the top active customers by amount spent
func (h *CustomersDefault) GetTopActiveCustomersByAmountSpent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		c, err := h.sv.FindTopActiveCustomersByAmountSpent(5)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error getting customers")
			return
		}

		// response
		// - serialize
		csJSON := make([]CustomerSpentJSON, len(c))
		for ix, v := range c {
			csJSON[ix] = CustomerSpentJSON{
				FirstName: v.FirstName,
				LastName:  v.LastName,
				Total:     v.Total,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "customers found",
			"data":    csJSON,
		})
	}
}

// CustomerInvoicesByConditionJSON is a struct that represents a customer invoices by condition in JSON format
type CustomerInvoicesByConditionJSON struct {
	Condition string  `json:"condition"`
	Total     float64 `json:"total"`
}

// GetInvoicesByCondition returns the total invoices by customer condition
func (h *CustomersDefault) GetInvoicesByCondition() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		c, err := h.sv.FindInvoicesByCondition()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error getting customers")
			return
		}

		// response
		// - serialize
		csJSON := make([]CustomerInvoicesByConditionJSON, len(c))
		for ix, v := range c {
			csJSON[ix] = CustomerInvoicesByConditionJSON{
				Condition: v.Condition,
				Total:     v.Total,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "customers found",
			"data":    csJSON,
		})
	}
}

// RequestBodyCustomer is a struct that represents the request body for a customer
type RequestBodyCustomer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Condition int    `json:"condition"`
}

// Create creates a new customer
func (h *CustomersDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - body
		var reqBody RequestBodyCustomer
		err := request.JSON(r, &reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "error deserializing request body")
			return
		}

		// process
		// - deserialize
		c := domain.Customer{
			CustomerAttributes: domain.CustomerAttributes{
				FirstName: reqBody.FirstName,
				LastName:  reqBody.LastName,
				Condition: reqBody.Condition,
			},
		}
		// - save
		err = h.sv.Save(&c)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error saving customer")
			return
		}

		// response
		// - serialize
		cs := CustomerJSON{
			Id:        c.Id,
			FirstName: c.FirstName,
			LastName:  c.LastName,
			Condition: c.Condition,
		}
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "customer created",
			"data":    cs,
		})
	}
}
