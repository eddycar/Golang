/* Ejercicio 1 - Descuento
Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos. Para ello necesitan una aplicación que les permita calcular el descuento basándose en dos variables: su precio y el descuento en porcentaje. La tienda espera obtener como resultado el valor con el descuento aplicado y luego imprimirlo en consola.
Crear la aplicación de acuerdo a los requerimientos. */

package descuento

import "fmt"

func CalcularDescuento(precio float32, descuento float32) {
	fmt.Printf("El descuento es de $: %0.2f\n", precio * descuento /100)
}