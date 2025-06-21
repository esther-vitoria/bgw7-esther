package repository

import (
	"app/internal"
	"context"
)

func NewRepositoryTicketMap(dbFile map[int]internal.Ticket, lastId int) *RepositoryTicketMap {
	defaultDb := make(map[int]internal.Ticket)

	if dbFile != nil {
		defaultDb = dbFile
	}

	return &RepositoryTicketMap{
		db:     defaultDb,
		lastId: lastId,
	}
}

type RepositoryTicketMap struct {
	db     map[int]internal.Ticket
	lastId int
}

func (r *RepositoryTicketMap) Get(ctx context.Context) (t map[int]internal.TicketAttributes, err error) {
	t = make(map[int]internal.TicketAttributes, len(r.db))
	for k, v := range r.db {
		t[k] = v.Attributes
	}

	return
}

func (r *RepositoryTicketMap) GetTicketByDestinationCountry(ctx context.Context, country string) (t map[int]internal.TicketAttributes, err error) {
	t = make(map[int]internal.TicketAttributes)
	for k, v := range r.db {
		if v.Attributes.Country == country {
			t[k] = v.Attributes
		}
	}

	return
}
