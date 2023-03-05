package main

import (
	"fmt"

	"github.com/eddycar/Golang/clase2/clase-asincronica/calcular_promedio"
	"github.com/eddycar/Golang/clase2/clase-asincronica/impuesto_salario"
)

func main() {
	// ---- Ejercicio 1 ----
	fmt.Println("Ejercicio 1")
	fmt.Println("El impuesto del salario es de: ",impuesto_salario.ImpuestoSalario(100000))

	// ---- Ejercicio 2 ----
	fmt.Println("Ejercicio 1")
	notas := []float32{5, 4.3, -4, 3.3, 3.5, 4}
	promedio, err := calcular_promedio.PromedioNotas(notas...)
	if err != "" {
		fmt.Println(err)
	}else {
		fmt.Println("El promedio es: ", promedio)
	}
}