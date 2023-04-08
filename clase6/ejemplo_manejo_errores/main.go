/* Impuestos salario

Definir una variable llamada salary y asignarle un valor de tipo int.
Crear un error con errors.New() con el mensaje "Salario muy bajo"
Crear una función que valide que salary sea menor o igual a 10.000 o retorne un mensaje que diga "Salario muy alto".
Comparar los errores son la función Is() dentro del main.

*/

package main

import (
	"errors"
	"fmt"
)

const lowSalary = 12000
var salary = 11000
var errSalary = errors.New("Salario muy bajo")

func validateSalary(salary int)(string, error){
	if salary <= lowSalary{
		return "", errSalary
	}
	return "Salario muy alto", nil
}

func main()  {
	msg, err := validateSalary(salary)
	if errors.Is(err, errSalary){
		fmt.Println(err)
		return
	}
	fmt.Println(msg)
}