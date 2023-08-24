package tickets

import (
	"encoding/csv"
	"os"
)

type Ticket struct {
	ID      string
	Nombre  string
	Email   string
	Destino string
	Hora    string
	Precio  string
}

var Tickets = []Ticket{}

// Funcion para leer el archivo .csv y pasarlo a slice
func ReadCSVFile(filename string) []Ticket {
	// Abrir el archivo CSV
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Crear un lector CSV
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// Chequear que haya suficientes campos en la fila
		if len(record) < 6 {
			continue
		}

		ticket := Ticket{
			ID:      record[0],
			Nombre:  record[1],
			Email:   record[2],
			Destino: record[3],
			Hora:    record[4],
			Precio:  record[5],
		}

		Tickets = append(Tickets, ticket)
	}
	return Tickets
}

// Requerimiento 1 - Obtener cuántas personas viajan a un país determinado.
func GetTotalTickets(destination string) int {
	totalTickets := 0

	for _, ticket := range Tickets {
		if ticket.Destino == destination {
			totalTickets++
		}
	}

	return totalTickets
}

// Requerimiento 2 - Calcular cuántas personas viajan segun etapa del dia:
func GetCountByPeriod(period string) int {
	count := 0

	startTime, endTime := getPeriodTimeRange(period)
	if startTime == nil || endTime == nil {
		panic("Período de tiempo no válido")
	}
	for _, ticket := range Tickets {
		ticketHour, err := time.Parse("15:04", ticket.Hora)
		if err != nil {
			fmt.Errorf("Error al formatear la hora: %s", ticket.Hora)
		}

		if ticketHour.After(*startTime) && ticketHour.Before(*endTime) {
			count++
		}
	}
	return count
}
