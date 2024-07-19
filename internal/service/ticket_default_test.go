package service_test

import (
	"testing"

	"app/internal/domain"
	"app/internal/repository"
	"app/internal/service"

	"github.com/stretchr/testify/require"
)

// Tests for ServiceTicketDefault.GetTotalAmountTickets
func TestServiceTicketDefault_GetTotalAmountTickets(t *testing.T) {
	t.Run("success to get total tickets", func(t *testing.T) {
		// arrange
		// - repository: mock
		rp := repository.NewRepositoryTicketMock()
		// - repository: set-up
		rp.FuncGet = func() (t map[int]domain.TicketAttributes, err error) {
			t = map[int]domain.TicketAttributes{
				1: {
					Name:    "John",
					Email:   "johndoe@gmail.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
			}
			return
		}

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		total, err := sv.GetTotalAmountTickets()

		// assert
		expectedTotal := 1
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})
}

func TestServiceTicketDefault_AverageDestination(t *testing.T) {
	t.Run("success to get average destination", func(t *testing.T) {
		// arrange
		// - repository: mock
		rp := repository.NewRepositoryTicketMock()
		// - repository: set-up
		rp.FuncGet = func() (map[int]domain.TicketAttributes, error) {
			return map[int]domain.TicketAttributes{
				1: {
					Name:    "John Doe",
					Email:   "john.doe@example.com",
					Country: "Paris",
					Hour:    "10:00",
					Price:   500,
				},
				2: {
					Name:    "Jane Smith",
					Email:   "jane.smith@example.com",
					Country: "London",
					Hour:    "11:00",
					Price:   600,
				},
				3: {
					Name:    "Alice Brown",
					Email:   "alice.brown@example.com",
					Country: "Paris",
					Hour:    "12:00",
					Price:   550,
				},
				4: {
					Name:    "Bob Brown",
					Email:   "bob.brown@example.com",
					Country: "Paris",
					Hour:    "13:00",
					Price:   550,
				},
			}, nil
		}

		rp.FuncGetTicketsByDestinationCountry = func(country string) (map[int]domain.TicketAttributes, error) {
			return map[int]domain.TicketAttributes{
				1: {
					Name:    "John Doe",
					Email:   "john.doe@example.com",
					Country: "Paris",
					Hour:    "10:00",
					Price:   500,
				},
				2: {
					Name:    "Alice Brown",
					Email:   "alice.brown@example.com",
					Country: "Paris",
					Hour:    "12:00",
					Price:   550,
				},
				3: {
					Name:    "Bob Brown",
					Email:   "bob.brown@example.com",
					Country: "Paris",
					Hour:    "13:00",
					Price:   550,
				},
			}, nil

		}

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		average, err := sv.AverageDestination("Paris")

		// assert
		expectedAverage := 75.00 // (3 / 4) * 100 = 75%
		require.NoError(t, err)
		require.Equal(t, expectedAverage, average)
	})
}
