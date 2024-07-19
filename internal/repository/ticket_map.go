package repository

import (
	"app/internal/domain"
	"context"
)

// NewRepositoryTicketMap creates a new repository for tickets in a map
func NewRepositoryTicketMap(db map[int]domain.TicketAttributes, lastId int) RepositoryTicketMap {
	return RepositoryTicketMap{
		db:     db,
		lastId: lastId,
	}
}

// RepositoryTicketMap implements the repository interface for tickets in a map
type RepositoryTicketMap struct {
	// db represents the database in a map
	// - key: id of the ticket
	// - value: ticket
	db map[int]domain.TicketAttributes

	// lastId represents the last id of the ticket
	lastId int
}

// GetAll returns all the tickets
func (r *RepositoryTicketMap) Get(ctx context.Context) (t map[int]domain.TicketAttributes, err error) {
	// create a copy of the map
	t = make(map[int]domain.TicketAttributes, len(r.db))
	for k, v := range r.db {
		t[k] = v
	}

	return
}

// GetTicketsByDestinationCountry returns the tickets filtered by destination country
func (r *RepositoryTicketMap) GetTicketsByDestinationCountry(ctx context.Context, country string) (t map[int]domain.TicketAttributes, err error) {
	// create a copy of the map
	t = make(map[int]domain.TicketAttributes)
	for k, v := range r.db {
		if v.Country == country {
			t[k] = v
		}
	}

	return
}
