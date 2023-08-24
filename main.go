package main

import (
	"fmt"
	"github.com/aquamantop/Parcial-GO.git/tickets"
)

func main() {
	tickets.ReadCSVFile("./tickets.csv")

	destino := "Russia"
	horario := "tarde"

	totalChan := make(chan int)
	perPeriodChan := make(chan int)
	averageTimeChan := make(chan float64)

	go func() {
		total := tickets.GetTotalTickets(destino)
		totalChan <- total
	}()

	go func() {
		perPeriod := tickets.GetCountByPeriod(horario)
		perPeriodChan <- perPeriod
	}()

	go func() {
		averageTime := tickets.PercentageDestination(destino)
		averageTimeChan <- averageTime
	}()

	// Recibe y muestra los resultados de los canales
	total := <-totalChan
	perPeriod := <-perPeriodChan
	averageTime := <-averageTimeChan

	fmt.Printf("Total de personas viajando a %s: %d\n", destino, total)
	fmt.Printf("Total de personas viajando a la %s: %d\n", horario, perPeriod)
	fmt.Printf("Porcentaje de personas que viajan a %s en un día: %.2f%%\n", destino, averageTime)

}
