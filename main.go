package main

import (
	"fmt"
	"github.com/aquamantop/Parcial-GO.git/tickets"
)

func main() {
	tickets.ReadCSVFile("./tickets.csv")
	fmt.Println(tickets.GetTotalTickets("Brazil"))
	fmt.Println(tickets.GetCountByPeriod("tarde"))
}
