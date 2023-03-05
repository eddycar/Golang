/*
Matrix
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear matrices en Go.
Para ello requieren una estructura Matrix que tenga los métodos:
	*Set: recibe una matriz de flotantes e inicializa los valores en la estructura Matrix.const
	*Print: imprime por pantalla la matriz de una forma mas visible(con saltos de linea entre filas).
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho
y si es cuadratica.
*/

// fila 0 * ancho 4 = 0
// fila 1 * ancho 4 = 4
// fila 2 * ancho 4 = 8
// fila 3 * ancho 4 = 12

// fila 0 * ancho 4 = 4
// fila 4 * ancho 4 = 8
// fila 8 * ancho 4 = 12

package main

import "fmt"

type Matrix struct {
	valores []float64
	alto    int
	ancho   int
}

func main() {
	m := Matrix{
		valores: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,13, 14, 15, 16},
		alto:    4,
		ancho:   4,
	}

	m.Set()
	m.Print()
	isCuadratic := m.Cuadratica()

	if isCuadratic {
		fmt.Println("Es cuadratica")
	} else {
		fmt.Println("No es cuadratica")
	}
}

func (m *Matrix) Set() {
	if len(m.valores) == 0 {
		fmt.Println("Los valores no fueron cargados correctamente")
		return
	}
}

func (m *Matrix) Print() {
	for fila := 0; fila < m.alto; fila++ {
		inicio := fila * m.ancho
		fin := inicio + m.ancho
		fmt.Println(m.valores[inicio:fin])
	}
}

func (m *Matrix) Cuadratica() bool {
	if m.alto == m.ancho && m.alto != 0 {
		return true
	}
	return false
}
