package main

import "fmt"

func sumar(a, b int) int {
	//mostrar la fecha y la hora
	resultado := a + b
	// Hacer algo después
	return resultado
}

func main() {
	fmt.Println(inter1(inter2(sumar))(10, 20))
// Lo mismo que la linea anterior pero haciendolo paso a paso:
	fm1 := inter2(sumar)
	fm2 := inter1(fm1)
	resultado := fm2(10, 20)
	fmt.Println(resultado)
}

func inter1(f func(int, int)int) func (int, int) int{
	return func(a, b int) int  {
		fmt.Println("inter1 ANTES")
		//Ejecutar la función recibida con la firma func(int, int)int 
		resultado := f(a, b)
		fmt.Println("inter1 DESPUES")
		return resultado
	}
	
}

func inter2(f func(int, int)int) func (int, int) int{
	return func(a, b int) int  {
		fmt.Println("inter2 ANTES")
		//Ejecutar la función recibida con la firma func(int, int)int 
		resultado := f(a, b)
		fmt.Println("inter2 DESPUES")
		return resultado
	}
	
}