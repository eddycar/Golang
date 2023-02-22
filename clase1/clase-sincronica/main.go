package main

import (
	"fmt"

	"github.com/eddycar/Golang/clase1/clase-sincronica/listar_nombres"
	"github.com/eddycar/Golang/clase1/clase-sincronica/obtener_edad"
)

func main() {

	// ---- Ejercicio 1 ----
	fmt.Println("\nEjercicio 1")
	var estudiantes = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}

	estudiantes = append(estudiantes, "Gabriela")

	listar_nombres.ListarEstudiantes(estudiantes)


	// ---- Ejerrcicio 2
	fmt.Println("\nEjercicio 2")

	var empleados = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	obtener_edad.ObtenerEdad(empleados, "Benjamin")
	obtener_edad.ObtenerEmpleadosMayores(empleados)

	empleados["Federico"] = 25
	fmt.Println(empleados)
	delete(empleados, "Pedro")
	fmt.Println(empleados)
}