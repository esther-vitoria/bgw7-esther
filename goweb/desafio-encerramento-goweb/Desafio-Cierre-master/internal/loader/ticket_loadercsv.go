package loader

import (
	"app/desafio-goweb/internal"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
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
func (t *LoaderTicketCSV) Load() (ta map[int]internal.Ticket, err error) {
	// open the file
	f, err := os.Open(t.filePath)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return
	}
	defer f.Close()

	// read the file
	r := csv.NewReader(f)

	// read the records
	ta = make(map[int]internal.Ticket)

	for {
		record, errRead := r.Read()
		if errRead != nil {
			if errRead == io.EOF {
				break
			}

			errRead = fmt.Errorf("error reading record: %v", err)
			err = errRead
			return
		}

		// serialize the record
		id, errId := strconv.Atoi(record[0])

		if errId != nil {
			errId = fmt.Errorf("error reading record: %v", err)
			err = errId
			return
		}
		price, errP := strconv.ParseFloat(record[5], 64)

		if errP != nil {
			errP = fmt.Errorf("error reading record: %v", err)
			err = errP
			return
		}

		ta[id] = internal.Ticket{
			Id: id,
			Attributes: internal.TicketAttributes{
				Name:    record[1],
				Email:   record[2],
				Country: record[3],
				Hour:    record[4],
				Price:   price,
			},
		}

	}
	return
}
