package handler

import (
	"app/internal"
	"net/http"

	"github.com/bootcamp-go/web/response"
)

func NewHandlerTicketDefault(sv internal.ServiceTicket) *TicketDefault {

	return &TicketDefault{sv: sv}

}

type TicketDefault struct {
	sv internal.ServiceTicket
}

func (h *TicketDefault) GetTotalAmountTickets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		response.JSON(w, http.StatusOK, nil)
	}

}
