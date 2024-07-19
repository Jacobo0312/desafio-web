package loader

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"app/internal/domain"
)

// NewLoaderTicketCSV creates a new ticket loader from a CSV file
func NewLoaderTicketCSV(filePath string) *LoaderTicketCSV {
	return &LoaderTicketCSV{
		filePath: filePath,
	}
}

// LoaderTicketCSV represents a ticket loader from a CSV file
type LoaderTicketCSV struct {
	filePath string
}

// Load loads the tickets from the CSV file
func (t *LoaderTicketCSV) Load() (tickets map[int]domain.TicketAttributes, err error) {

	log.Println("Loading tickets from CSV file...")

	// open the file
	f, err := os.Open(t.filePath)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return
	}
	defer f.Close()

	// read the file
	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// read the records
	tickets = make(map[int]domain.TicketAttributes)
	for _, record := range records {

		// serialize the record
		id, err3 := strconv.Atoi(record[0])

		if err3 != nil {
			fmt.Println("Error converting id to int")
			return
		}

		//Convert price to float64

		price, err4 := strconv.ParseFloat(record[5], 64)

		if err4 != nil {
			fmt.Println("Error converting price to float64")
			return
		}

		ticket := domain.TicketAttributes{
			Name:    record[1],
			Email:   record[2],
			Country: record[3],
			Hour:    record[4],
			Price:   price,
		}

		// add the ticket to the map
		tickets[id] = ticket
	}

	return
}
