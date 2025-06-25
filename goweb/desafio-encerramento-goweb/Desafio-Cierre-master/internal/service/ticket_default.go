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
	tickets, err := s.rp.GetAll()
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (s *ServiceTicketDefault) GetTotalAmountTickets() (total int, err error) {
	tickets, err := s.rp.GetAll()
	total = len(tickets)

	if err != nil {
		return 0, err
	}
	return total, nil
}

func (s *ServiceTicketDefault) GetTicketsAmountByDestinationCountry(country string) (total int, err error) {
	tickets, err := s.rp.GetTicketsByDestinationCountry(country)
	total = len(tickets)

	if err != nil {
		return 0, err
	}
	return total, nil
}

func (s *ServiceTicketDefault) GetTicketsByDestinationCountry(country string) (tickets map[int]internal.Ticket, err error) {
	tickets, err = s.rp.GetTicketsByDestinationCountry(country)

	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (s *ServiceTicketDefault) GetPercentageTicketsByDestinationCountry(country string) (percent float64, err error) {
	total, err := s.GetTotalAmountTickets()
	if err != nil {
		return 0, err
	}
	countryAmount, err := s.GetTicketsAmountByDestinationCountry(country)
	if err != nil {
		return 0, err
	}
	percent = (float64(countryAmount) / float64(total)) * 100
	return
}
