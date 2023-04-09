package main

import (
	"fmt"
	"time"
)

func sumar(a, b int) int {
	//mostrar la fecha y la hora
	resultado := a + b
	// Hacer algo después
	return resultado
}

func mostrarTiempo(f func(int, int)int) func(a, b int) int{
	return func(a, b int) int {
		fmt.Println(time.Now())
		return f(a, b)
	}
}

func mostrarPosterior(f func(int, int) int) func(int, int)int  {
	return func(a, b int) int {
		resultado := f(a, b)
		// Hacer algo después
		fmt.Println("se ejecutó la fución")
		return resultado
	}
}

// funcion main aplicando comportamiento a funcion suma
// func main()  {
// 	// Imprimiendo resultado de forma directa 
// 	fmt.Println("Resultado: ", mostrarTiempo(sumar)(10,20)) 
// 	// Imprimiendo resultado paso a paso 
// 	f := mostrarTiempo(sumar)
// 	resultado := f(10, 20)
// 	fmt.Println(resultado)
// }

// funcion main aplicando dos comportamientos a la funcion suma

func main()  {
	fmt.Println("Resultado", mostrarTiempo(mostrarPosterior(sumar))(10,20))
	// Lo mismo que la linea anterior pero haciendolo paso a paso:
	fm1 := mostrarPosterior(sumar)
	fm2 := mostrarTiempo(fm1)
	resultado := fm2(10, 20)
	fmt.Println(resultado)
}