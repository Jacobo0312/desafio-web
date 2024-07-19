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
func NewServiceTicketDefault(rp repository.RepositoryTicketMap) ServiceTicketDefault {
	return ServiceTicketDefault{
		rp: rp,
	}
}

// GetTotalTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalTickets() (total int, err error) {

	tickets, err := s.rp.Get()

	if err != nil {
		return 0, err
	}
	return len(tickets), nil
}

// GetTotalAmountTickets returns the total amount of tickets
func (s *ServiceTicketDefault) GetTotalAmountTickets() (total int, err error) {
	tickets, err := s.rp.Get()
	if err != nil {
		return
	}

	total = len(tickets)
	return
}

func (s *ServiceTicketDefault) AverageDestination(destination string) (float64, error) {
	tickets, err := s.rp.Get()
	if err != nil {
		return 0, err
	}
	total, err := s.rp.GetTicketsByDestinationCountry(destination)

	if err != nil {
		return 0, err
	}

	average := float64(len(total)) / float64(len(tickets)) * 100
	return average, nil
}

func (s *ServiceTicketDefault) GetTotalAmountTicketsByCountry(country string) (total int, err error) {
	tickets, err := s.rp.GetTicketsByDestinationCountry(country)
	if err != nil {
		return
	}

	total = len(tickets)
	return
}
