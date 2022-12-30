package tickets

import (
	"testing"
)

func TestGetTotalTicketsBrazil(t *testing.T) {
	//Arrange
	data := readFile("/Users/williamgarci/Documents/bootcamp/desafio-go-bases/tickets.csv")
	ticketList := convertToJson(data)
	brazilPeople := 45

	//Act
	result, _ := GetTotalTickets("Brazil", ticketList)

	//Assert
	if result != brazilPeople {
		t.Errorf("Prueba fallo")
	}
}

func TestGetTotalTicketsMexico(t *testing.T) {
	//Arrange
	data := readFile("/Users/williamgarci/Documents/bootcamp/desafio-go-bases/tickets.csv")
	ticketList := convertToJson(data)
	mexicoPeople := 11

	//Act
	result, _ := GetTotalTickets("Mexico", ticketList)

	//Assert
	if result != mexicoPeople {
		t.Errorf("Prueba fallo")
	}
}

func TestGetTotalTicketsFailed(t *testing.T) {
	//Arrange
	data := readFile("/Users/williamgarci/Documents/bootcamp/desafio-go-bases/tickets.csv")
	ticketList := convertToJson(data)

	//Act
	_, err := GetTotalTickets("", ticketList)

	//Assert
	if err != nil {
		t.Log("Prueba Erronea exitosa")
	}
}

func TestAverageDestinationMexxico(t *testing.T) {
	// Arrange
	peopleInMexico := 11
	totalPeople := 999
	avg := 0.011011011011011011

	// Act
	result, _ := AverageDestination(peopleInMexico, totalPeople)

	// Assert
	if result != avg {
		t.Errorf("Prueba fallo")
	}
}

func TestGetPeriods(t *testing.T) {
	// Arrange
	data := readFile("/Users/williamgarci/Documents/bootcamp/desafio-go-bases/tickets.csv")
	ticketList := convertToJson(data)

	// Act
	midnight, morning, evening, night := getPeriod(ticketList)

	// Assert
	if midnight != 304 {
		t.Errorf("Midnight fallo")
	}
	if morning != 256 {
		t.Errorf("Morning fallo")
	}
	if evening != 288 {
		t.Errorf("Evening fallo")
	}
	if night != 151 {
		t.Errorf("Night fallo")
	}
}
