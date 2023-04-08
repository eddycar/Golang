/*Impuestos de salario
En la función main, definir una variable salary y asignarle un valor de tipo “int”. Luego, crear
un error personalizado con un struct que implemente Error() con el mensaje “Error: el salario
ingresado no alcanza el mínimo imponible” y lanzarlo en caso de que salary sea menor a
150.000. De lo contrario, imprimir por consola el mensaje “Debe pagar impuesto”.
*/

package main

import (
	"errors"
	"fmt"
)

func main()  {
	var salary int = 150000

	if salary < 150000{
		fmt.Println(errors.New("Error: el salario ingresado no alcanza el mínimo imponible"))
	}else{
		fmt.Println("Debe pagar impuesto")
	}

}