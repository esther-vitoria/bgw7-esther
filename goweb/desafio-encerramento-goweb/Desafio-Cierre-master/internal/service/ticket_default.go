package service

import (
	"app/desafio-goweb/internal"
)

func NewServiceTicketDefault(rp internal.RepositoryTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

type ServiceTicketDefault struct {
	rp internal.RepositoryTicket
}

func (s *ServiceTicketDefault) GetAllTickets() (t map[int]internal.Ticket, err error) {
	return s.rp.GetAll()
}

func (s *ServiceTicketDefault) GetTotalAmountTickets() (total int, err error) {
	tickets, err := s.rp.GetAll()
	total = len(tickets)
	return total, err
}

func (s *ServiceTicketDefault) GetTicketsAmountByDestinationCountry(country string) (total int, err error) {
	tickets, err := s.rp.GetTicketsByDestinationCountry(country)
	total = len(tickets)

	return total, err
}

func (s *ServiceTicketDefault) GetTicketsByDestinationCountry(country string) (tickets map[int]internal.Ticket, err error) {
	tickets, err = s.rp.GetTicketsByDestinationCountry(country)

	return tickets, err
}

func (s *ServiceTicketDefault) GetPercentageTicketsByDestinationCountry(country string) (avg float64, err error) {
	total, err := s.GetTotalAmountTickets()
	if err != nil {
		return
	}
	countryAmount, err := s.GetTicketsAmountByDestinationCountry(country)
	if err != nil {
		return
	}
	avg = float64(countryAmount) / float64(total)
	return
}
