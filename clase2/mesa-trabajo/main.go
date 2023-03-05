/*Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
1. Categoría C: su salario es de $1.000 por hora.
2. Categoría B: su salario es de $1.500 por hora, más un 20 % de su salario mensual.
3. Categoría A: su salario es de $3.000 por hora, más un 50 % de su salario mensual.
4.	Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
*/

package main

import "fmt"

func main() {
	salario, err := calcularSalario(120, "categoriaC")
	if err != ""{
		fmt.Println(err)
	}else {
		fmt.Printf("El salario es $%.2f" , salario)
	}
}

func calcularSalario(minutosTrabajados int, categoria string)(float32, string) {
	var horasDeTrabajo float32
	var salario float32
	horasDeTrabajo = float32(minutosTrabajados) / 60

	switch categoria {
	case "categoriaA":
		salario = horasDeTrabajo * 1000
		return salario,""
	case "categoriaB":
		salario = horasDeTrabajo*1500 + (horasDeTrabajo * 1500 * 0.2)
		return salario,""
	case "categoriaC":
		salario = horasDeTrabajo*3000 + (horasDeTrabajo * 3000 * 0.5)
		return salario,""
	default:
		return 0.0, "No se ingresó la categoria correcta"
	}
}
