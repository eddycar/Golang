/*Calcular estadisticas 
Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los estudiantes
de un curso. Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.

Para esto, se solicita generar una función que indique qué tipo de cálculo se requiere realizar (mínimo, máximo promedio)
y que evuelva otra función y un mesaje (en caso de que el cálculo no esté definido), que se le pueda pasar una cantidad N
de enteros y devuelva el cálculo que se indicó en la función anterior
*/

package main

import (
	"fmt"
	"errors"
)

func main() {

	const (
		minimo   = "minimo"
		promedio = "promedio"
		maximo   = "maximo"
	)

	minFunc, err := operacion(minimo)
	promFunc, err := operacion(promedio)
	maxFunc, err := operacion(maximo)

	if err != nil {
		fmt.Println(err)
	}

	valorMinimo := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("La calificación con menor puntuación es: ", valorMinimo)
	fmt.Println("El promedio de las calificaciones es: ", valorPromedio)
	fmt.Println("La calificación con mayor puntuación es: ", valorMaximo)

}

func calcularMinimo(num ...int) int {
	minimo := num[0]
	for _, valor := range num {
		if valor < minimo {
			minimo = valor
		}
	}
	return minimo
}

func calcularMaximo(num ...int) int {
	var maximo int
	for _, valor := range num {
		if valor > maximo {
			maximo = valor
		}
	}
	return maximo
}

func calcularPromedio(num ...int) int {
	var suma int
	for _, valor := range num {
		suma += valor
	}
	return suma / len(num)
}

func operacion(calculo string) (func(...int)int, error){
	switch calculo {
	case "minimo":
		min := calcularMinimo
		return min, nil
	case "promedio":
		prom := calcularPromedio
		return prom, nil
	case "maximo":
		max := calcularMaximo
		return max, nil
	default:
		 return nil, errors.New("No se reconoce el calculo solicitado")
	}

}