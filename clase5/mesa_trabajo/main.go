/*Productos
Algunas tiendas de e-commerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total. La empresa tiene tres tipos de productos: Pequeño, Mediano y Grande. Pero se espera que sean muchos más.
Los costos adicionales para cada uno son:
	*Pequeño: solo tiene el costo del producto.
	*Mediano: el precio del producto + un 3% por mantenerlo en la tienda.
	*Grande: el precio del producto + un 6% por mantenerlo en la tienda y un adicional de $2500 de costo de envío.
El porcentaje de mantenerlo en la tienda se calcula sobre el precio del producto.
Se requiere una función factory que reciba el tipo de producto y el precio, y retorne una interfaz Producto que tenga el método Precio.
Se debe poder ejecutar el método Precio y que el método devuelva el precio total basándose en el costo del producto y los adicionales, en caso de que los tenga.
*/

package main

import "fmt"

const(
	pequenio = "pequeño"
	mediano = "mediano"
	grande = "grande"
)

type (
	Producto interface{
		Precio() float64
	}

	Item struct{
		costo float64
		mantenimiento float64
		adicionales float64

	}
)

func main()  {
	i1 := factory(pequenio, 1000)
	fmt.Println(i1.Precio())

	i2 := factory(mediano, 2000)
	fmt.Println(i2.Precio())

	i3 := factory(grande, 4000)
	fmt.Println(i3.Precio())
}

func (i Item)Precio()float64{
	return i.costo + (i.costo * i.mantenimiento/100) + i.adicionales
}

func factory(tipoPrducto string, precio float64) Producto{
	switch tipoPrducto {
	case pequenio:
		return Item{costo: precio}
	case mediano:
		return Item{costo: precio, mantenimiento: 3}
	case grande:
		return Item{costo: precio, mantenimiento: 6, adicionales: 2500}
	default:
		return nil
	}
}