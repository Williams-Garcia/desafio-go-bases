package tickets

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// -Una función que calcule cuántas personas viajan a un país determinado:
// -Una o varias funciones que calcule cuántas personas viajan en madrugada (0 → 6), mañana (7 → 12), tarde (13 → 19) y noche (20 → 23)
// -Calcular el promedio de personas que viajan a un país determinado en un dia:

var (
	errFailedDestination = &customError{"Error: el destino no puede estar en blanco"}
	errFailedRead        = &customError{"Error: el archivo indicado no fue encontrado o está dañado"}
	errFailedConvert     = &customError{"Error: el archivo no se pudo convertir"}
)

type customError struct {
	msg string
}

func (c *customError) Error() string {
	return c.msg
}

type Ticket struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Country string `json:"country"`
	Time    string `json:"time"`
	Other   string `json:"other"`
}

func Init() {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	country := "Brazil"
	data := readFile("tickets.csv")
	ticketList := convertToJson(data)
	totalPeople, err := GetTotalTickets(country, ticketList)
	if err != nil {
		panic(err.Error())
	}
	avgDestination, err := AverageDestination(totalPeople, len(ticketList))
	if err != nil {
		panic(err.Error())
	}
	midnight, morning, evening, night := getPeriod(ticketList)
	fmt.Printf("%d personas viajaron a %s\n", totalPeople, country)
	fmt.Printf("Promedio de personas que viajaron a %s : %v \n", country, avgDestination)
	fmt.Printf("Midnight: %d\n", midnight)
	fmt.Printf("Morning: %d\n", morning)
	fmt.Printf("evening: %d\n", evening)
	fmt.Printf("night: %d\n", night)
}

func readFile(fileName string) [][]string {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Errorf("%w", errFailedRead))
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	return data
}

func convertToJson(data [][]string) []Ticket {
	var ticketList []Ticket
	for i, line := range data {
		if i > 0 { // omit header line
			var ticket Ticket
			for j, field := range line {
				if j == 0 {
					var err error
					ticket.Id, err = strconv.Atoi(field)
					if err != nil {
						continue
					}
				} else if j == 1 {
					ticket.Name = field
				} else if j == 2 {
					ticket.Email = field
				} else if j == 3 {
					ticket.Country = field
				} else if j == 4 {
					ticket.Time = field
				} else if j == 5 {
					ticket.Other = field
				}

			}
			ticketList = append(ticketList, ticket)
		}
	}
	return ticketList
}

// ejemplo 1
func GetTotalTickets(destination string, ticketList []Ticket) (destinationCount int, err error) {
	var count = 0
	if destination == "" {
		return 0, errFailedDestination
	}
	for _, ticket := range ticketList {
		if ticket.Country == destination {
			count++
		}
	}
	return count, nil
}

// ejemplo 2
func getPeriod(ticketList []Ticket) (midnight int, morning int, evening int, night int) {

	for _, ticket := range ticketList {
		hr := strings.Split(ticket.Time, ":")
		hrInt, err := strconv.Atoi(hr[0])
		if err != nil {
			panic(err)
		}
		if hrInt == 0 || hrInt <= 6 {
			midnight++
		} else if hrInt == 7 || hrInt <= 12 {
			morning++
		} else if hrInt == 13 || hrInt <= 19 {
			evening++
		} else {
			night++
		}
	}

	return midnight, morning, evening, night
}

// ejemplo 3
func AverageDestination(peopleInDestination int, totalPeople int) (avg float64, err error) {
	return float64(peopleInDestination) / float64(totalPeople), nil
}
