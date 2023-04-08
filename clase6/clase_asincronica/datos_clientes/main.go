/*Datos de clientes
Un estudio contable necesita acceder a los datos de sus empleados para poder realizar
distintas liquidaciones. Para ello, cuentan con todo el detalle necesario en un archivo TXT.
1. Desarrollar el código necesario para leer los datos de un archivo llamado
“customers.txt”. Sin embargo, debemos tener en cuenta que la empresa no nos ha
pasado el archivo a leer por el programa. Dado que no contamos con el archivo
necesario, se obtendrá un error. En tal caso, el programa deberá arrojar un panic al
intentar leer un archivo que no existe, mostrando el mensaje “El archivo indicado no
fue encontrado o está dañado”.
2. Más allá de eso, deberá siempre imprimirse por consola “Ejecución finalizada”.
*/

package main

import (
	"fmt"
	"os"
)


func main()  {
	fmt.Println("Iniciando ...")
	defer func ()  {
		err := recover()
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("\nEjecución finalizada")
	}()

	_, err := os.Open("customers.txt")
	if err != nil {
		panic(err)
	}
}