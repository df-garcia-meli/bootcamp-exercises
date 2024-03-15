package main

import (
	"desafio-go-bases/internal/tickets"
	"fmt"
)

func main() {
	// Data Creation
	sales, err := tickets.NewSales("tickets.csv")
	if err != nil {
		fmt.Println(err)
	}

	// Excercise 1
	totalTickets, err := sales.GetTotalTickets("Brazil")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Total tickets: ", totalTickets)

	// Excercise 2
	totalTicketsByPeriod, err := sales.GetCountByPeriod("night")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Total tickets by period: ", totalTicketsByPeriod)

	// Excercise 3
	averageTicketsByCountry, err := sales.AverageDestination("China")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Average tickets by country: ", averageTicketsByCountry)
}
