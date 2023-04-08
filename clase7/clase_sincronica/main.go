/* Leer archivo
La empresa encargada de vender los productos de limpieza ahora necesita leer el archivo almacenado.
Para esto requiere que se imprima por pantalla mostrando los valores tbulados, con el titulo
(tabulado a la izquierda para el ID y a la derecha para el precio y Cantidad), el precio, la cantidad y, abajo el precio,
se debe visualizar el total (sumando precio por cantidad).

Ejemplo:
ID					Precio			Cantidad
111223				30012.00			1
444321				100000.00			4
434321				50.50				1

Total 				100362.50
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dataByte, err := os.ReadFile("./data.csv")
	if err != nil {
		panic(err)
	}

	data := strings.Split(string(dataByte), ";")
	fmt.Println(data)

	for i:=0; i<len(data)-1; i++{

		line := strings.Split(data[i],",")
		var newLine []string
		for j := 0; j<len(line); j++{
		 append(newLine,line[j])
		}
		fmt.Printf("%s\t\t\n", newLine)
	}
	

}

