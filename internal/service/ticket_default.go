package service

import (
	"app/internal/repository"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp repository.RepositoryTicketMap
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp repository.RepositoryTicketMap) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetTotalTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalTickets() (total int, err error) {
	return
}