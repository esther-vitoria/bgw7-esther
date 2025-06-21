package service

import (
	"app/internal"
)

func NewServiceTicketDefault(rp internal.RepositoryTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

type ServiceTicketDefault struct {
	rp internal.RepositoryTicket
}

func (s *ServiceTicketDefault) GetTotalAmountTickets() (total int, err error) {
	return
}
