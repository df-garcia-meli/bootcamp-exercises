package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Ticket Struct
type Ticket struct {
	id          int
	name        string
	email       string
	destination string
	flightTime  string
	price       float64
}

// Create a new ticket
func NewTicket(id int, name string, email string, destination string, flightTime string, price float64) Ticket {
	return Ticket{
		id:          id,
		name:        name,
		email:       email,
		destination: destination,
		flightTime:  flightTime,
		price:       price,
	}
}

type CSVSalesInterface interface {
	GetTotalTickets(destination string) (int, error)
	GetCountByPeriod(period string) (int, error)
	AverageDestination(destination string) (float64, error)
}

// Sales Struct
type Sales struct {
	Tickets []Ticket
}

// Create New Sales
func NewSales(filePath string) (Sales, error) {
	sales := Sales{}

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return sales, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return sales, err
	}

	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Println(err)
		}

		price, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			fmt.Println(err)
		}

		ticket := NewTicket(id, record[1], record[2], record[3], record[4], price)
		sales.Tickets = append(sales.Tickets, ticket)
	}
	return sales, nil
}

// ejemplo 1
func (s *Sales) GetTotalTickets(destination string) (int, error) {
	var result []Ticket

	if len(s.Tickets) == 0 {
		return 0, errors.New("no tickets")
	}

	for _, ticket := range s.Tickets {
		if ticket.destination == destination {
			result = append(result, ticket)
		}
	}
	return len(result), nil
}

// ejemplo 2
func (s *Sales) GetCountByPeriod(timeString string) (int, error) {
	var result []Ticket
	if len(s.Tickets) == 0 {
		return 0, errors.New("no tickets")
	}

	var lowerTimeLimit string
	var upperTimeLimit string
	switch timeString {
	case "earlyMorning":
		lowerTimeLimit = "00:00"
		upperTimeLimit = "6:00"
	case "morning":
		lowerTimeLimit = "6:00"
		upperTimeLimit = "12:00"
	case "afternoon":
		lowerTimeLimit = "12:00"
		upperTimeLimit = "18:00"
	case "night":
		lowerTimeLimit = "18:00"
		upperTimeLimit = "24:00"
	}

	for _, ticket := range s.Tickets {
		ticketTime, _ := time.Parse("15:04", ticket.flightTime)
		lowerTime, _ := time.Parse("15:04", lowerTimeLimit)
		upperTime, _ := time.Parse("15:04", upperTimeLimit)

		if ticketTime.After(lowerTime) && ticketTime.Before(upperTime) {
			result = append(result, ticket)
		}
	}

	return len(result), nil
}

// ejemplo 3
func (s *Sales) AverageDestination(destination string) (float64, error) {
	var result []Ticket
	if len(s.Tickets) == 0 {
		return 0, errors.New("no tickets")
	}

	for _, ticket := range s.Tickets {
		if ticket.destination == destination {
			result = append(result, ticket)
		}
	}

	return float64(len(result)) / float64(len(s.Tickets)), nil
}
