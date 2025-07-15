package repository

import (
	"app/desafio-goweb/internal"
	"app/desafio-goweb/pkg/apperrors"
)

func NewRepositoryTicketMap(db map[int]internal.TicketAttributes, lastId int) *RepositoryTicketMap {
	// default db
	defaultDb := make(map[int]internal.TicketAttributes)
	if db != nil {
		defaultDb = db
	}
	return &RepositoryTicketMap{db: defaultDb, lastId: lastId}
}

type RepositoryTicketMap struct {
	db     map[int]internal.TicketAttributes
	lastId int
}

func (r *RepositoryTicketMap) Get() (t map[int]internal.TicketAttributes, err error) {
	t = make(map[int]internal.TicketAttributes, len(r.db))
	for k, v := range r.db {
		t[k] = v
	}

	if len(t) == 0 {
		return nil, apperrors.ErrLoadCSV
	}
	return t, nil
}

func (r *RepositoryTicketMap) GetTicketsByDestinationCountry(country string) (t map[int]internal.TicketAttributes, err error) {
	t = make(map[int]internal.TicketAttributes)
	for k, v := range r.db {
		if v.Country == country {
			t[k] = v
		}
	}

	if len(t) == 0 {
		return nil, apperrors.ErrCountryNotFound
	}

	return t, nil
}
