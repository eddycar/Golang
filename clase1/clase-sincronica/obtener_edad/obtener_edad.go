// Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente map, debemos imprimir la edad de Benjamin.
// var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
// Por otro lado, también es necesario:
// Saber cuántos de sus empleados son mayores de 21 años.
// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
// Eliminar a Pedro del map.

package obtener_edad

import "fmt"

func ObtenerEdad(mapa map[string]int, nombre string){
	fmt.Println("La edad de ",nombre , "es", mapa[nombre])
}

func ObtenerEmpleadosMayores(empleados map[string]int){
	contador := 0
	for _, edad := range empleados {
		if edad > 21 {
			contador++
		}
	}
	fmt.Println("La cantidad de empleados mayores a 21 años son: ", contador)
}