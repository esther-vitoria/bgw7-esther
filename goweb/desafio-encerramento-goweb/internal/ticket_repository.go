package internal

// RepositoryTicket represents the repository interface for tickets
type RepositoryTicket interface {
	// GetAll returns all the tickets
	Get() (t map[int]TicketAttributes, err error)

	// GetTicketByDestinationCountry returns the tickets filtered by destination country
	GetTicketsByDestinationCountry(country string) (t map[int]TicketAttributes, err error)
}
