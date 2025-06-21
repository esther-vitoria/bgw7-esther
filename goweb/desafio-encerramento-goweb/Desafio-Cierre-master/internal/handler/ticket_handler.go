package handler

import (
	"app/desafio-goweb/internal"
	"app/desafio-goweb/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHandlerTicketDefault(sv internal.ServiceTicket) *TicketDefault {

	return &TicketDefault{sv: sv}

}

type TicketDefault struct {
	sv internal.ServiceTicket
}

func (h *TicketDefault) GetAllTickets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tickets, err := h.sv.GetAllTickets()

		if err != nil {
			utils.Respond(w, http.StatusBadRequest, &utils.RespondBodyProduct{
				Message: err.Error(),
				Data:    nil,
				Error:   true,
			})
			return
		}

		utils.Respond(w, http.StatusOK, &utils.RespondBodyProduct{
			Message: "Success!",
			Data:    tickets,
		})
	}

}

func (h *TicketDefault) GetTicketsAmountByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dest := chi.URLParam(r, "dest")

		total, err := h.sv.GetTicketsAmountByDestinationCountry(dest)

		if err != nil {
			utils.Respond(w, http.StatusNotFound, &utils.RespondBodyProduct{
				Message: err.Error(),
				Data:    nil,
				Error:   true,
			})
			return
		}

		utils.Respond(w, http.StatusOK, &utils.RespondBodyProduct{
			Message: "The total tickets for this country is:",
			Data:    total,
		})
	}

}

func (h *TicketDefault) GetPercentageTicketsByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dest := chi.URLParam(r, "dest")

		avg, err := h.sv.GetPercentageTicketsByDestinationCountry(dest)

		if err != nil {
			utils.Respond(w, http.StatusNotFound, &utils.RespondBodyProduct{
				Message: err.Error(),
				Data:    nil,
				Error:   true,
			})
			return
		}

		utils.Respond(w, http.StatusOK, &utils.RespondBodyProduct{
			Message: "The tickets average for this country is:",
			Data:    avg,
		})
	}

}
