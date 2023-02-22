/* Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
El banco tiene ciertas reglas para saber a qué cliente se le puede otorgar: solo le otorga préstamos
a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo.
Dentro de los préstamos que otorga, no les cobrará interés a los que su sueldo sea mayor a $100.000.
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.
Pista: Tu código tiene que poder imprimir al menos tres mensajes diferentes. */

package prestamo

import "fmt"

func esMayorDeEdad(edad int){
    if edad < 0 {
        fmt.Println("La edad ingresada no es válida")
    } else if edad < 22 {
        fmt.Println("Es menor de edad")
    } else {
        fmt.Println("Es mayor de edad")
    }
}

func cumpleAntiguedad(antiguedad int){
    if antiguedad < 0 {
        fmt.Println("La antiguedad ingresada no es válida")
    } else if antiguedad < 12 {
        fmt.Println("No cumple con antiguedad")
    } else {
        fmt.Println("Cumple con antiguedad")
    }
}

func pagaInteres(sueldo float64){
	if sueldo < 0 {
        fmt.Println("La sueldo ingresada no es válida")
    } else if sueldo < 100000 {
        fmt.Println("Paga interés")
    } else {
        fmt.Println("No paga interés")
    }
}


func DetallePrestamo(nombreCliente string, edad int, antiguedad int, sueldo float64){
	fmt.Println("Nombre del cliente:\t", nombreCliente)
	esMayorDeEdad(edad);
	cumpleAntiguedad(antiguedad)
	pagaInteres(sueldo)
}