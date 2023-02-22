package main

import (
	"fmt"

	"github.com/eddycar/Golang/clase1/descuento.go"
	"github.com/eddycar/Golang/clase1/prestamo"
)

func main() {

	// ---- Ejercicio 1 ----
	fmt.Println("\nEjercicio 1")
	descuento.CalcularDescuento(3500.8, 10.5)
	// ---- Ejercicio 2 ----
	fmt.Println("\nEjercicio 2")
	prestamo.DetallePrestamo("edi", 18, 11, 90000)
}