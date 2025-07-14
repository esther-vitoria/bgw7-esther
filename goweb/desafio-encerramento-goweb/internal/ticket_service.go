package internal

type ServiceTicket interface {
	// GetAllTickets returns all Tickets
	GetAllTickets() (tickets map[int]TicketAttributes, err error)

	// GetTicketsByDestinationCountry returns Tickets by destination country
	GetTicketsByDestinationCountry(country string) (tickets map[int]TicketAttributes, err error)

	// GetTotalAmountTickets returns the total amount of tickets
	GetTotalAmountTickets() (total int, err error)

	// GetTicketsAmountByDestinationCountry returns the amount of tickets filtered by destination country
	GetTicketsAmountByDestinationCountry(country string) (total int, err error)

	// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by destination country
	GetPercentageTicketsByDestinationCountry(country string) (avg float64, err error)
}
