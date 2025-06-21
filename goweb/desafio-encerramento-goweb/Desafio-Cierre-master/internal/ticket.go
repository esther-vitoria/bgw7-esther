package internal

import "context"

// TicketAttributes is an struct that represents a ticket
type TicketAttributes struct {
	// Name represents the name of the owner of the ticket
	Name string `json:"name"`
	// Email represents the email of the owner of the ticket
	Email string `json:"email"`
	// Country represents the destination country of the ticket
	Country string `json:"country"`
	// Hour represents the hour of the ticket
	Hour string `json:"hour"`
	// Price represents the price of the ticket
	Price float64 `json:"price"`
}

type TicketJson struct {
	Id int `json:"id"`
	// Name represents the name of the owner of the ticket
	Name string `json:"name"`
	// Email represents the email of the owner of the ticket
	Email string `json:"email"`
	// Country represents the destination country of the ticket
	Country string `json:"country"`
	// Hour represents the hour of the ticket
	Hour string `json:"hour"`
	// Price represents the price of the ticket
	Price float64 `json:"price"`
}

// Ticket represents a ticket
type Ticket struct {
	// Id represents the id of the ticket
	Id int `json:"id"`
	// Attributes represents the attributes of the ticket
	Attributes TicketAttributes `json:"attributes"`
}

type RepositoryTicket interface {
	Get(ctx context.Context) (t map[int]TicketAttributes, err error)
	GetTicketByDestinationCountry(ctx context.Context, country string) (t map[int]TicketAttributes, err error)
}

type ServiceTicket interface {
	GetTotalAmountTickets() (total int, err error)
}
