package repository

import (
	"app/desafio-goweb/internal"
	"app/desafio-goweb/pkg/apperrors"
)

func NewRepositoryTicketMap(db map[int]internal.Ticket, lastId int) *RepositoryTicketMap {
	// default db
	defaultDb := make(map[int]internal.Ticket)
	if db != nil {
		defaultDb = db
	}
	return &RepositoryTicketMap{db: defaultDb, lastId: lastId}
}

type RepositoryTicketMap struct {
	db     map[int]internal.Ticket
	lastId int
}

func (r *RepositoryTicketMap) GetAll() (t map[int]internal.Ticket, err error) {
	t = make(map[int]internal.Ticket, len(r.db))
	for k, v := range r.db {
		t[k] = v
	}

	return
}

func (r *RepositoryTicketMap) GetTicketsByDestinationCountry(country string) (t map[int]internal.Ticket, err error) {
	t = make(map[int]internal.Ticket)
	for k, v := range r.db {
		if v.Attributes.Country == country {
			t[k] = v
		}
	}

	if len(t) == 0 {
		return nil, apperrors.ErrCountryNotFound
	}

	return t, nil
}
