package main

import (
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

var (
	Kitchen = "3:04"
)

// Una aerolínea pequeña tiene un sistema de reservas de pasajes a diferentes países, este retorna un archivo con la información de
// los pasajes sacados en las últimas 24 horas.
// La aerolínea necesita un programa para extraer información de las ventas del día y así, analizar las tendencias de compra.
// El archivo en cuestión es del tipo valores separados por coma (csv por su siglas en inglés), donde los campos están compuestos
// por: id, nombre, email, país de destino, hora del vuelo y precio.

func main() {
	// total, err := tickets.GetTotalTickets("Brazil")
	tickets.Init()
	// now, err := time.Parse(Kitchen, "15:45")
	// if err != nil {
	// 	println(err)
	// }
	// // the time and date in this string is intentional, as we will discuss soon
	// fmt.Println(now)
}
