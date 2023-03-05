/* Registro de estudiantes

Una universidad necesita registrar a los estudiantes y generar una funcionalidad para
imprimir el detalle de los datos de cada uno de ellos, de la siguiente manera:
● Nombre: [Nombre del alumno]
● Apellido: [Apellido del alumno]
● DNI: [DNI del alumno]
● Fecha: [Fecha ingreso alumno]
Los valores que están en corchetes deben ser reemplazados por los datos brindados por los
alumnos. Para ello es necesario generar una estructura Alumno con las variables Nombre,
Apellido, DNI, Fecha y que tenga un método detalle.

*/

package main

import "fmt"

type Alumno struct{
	Nombre string
	Apellido string
	DNI string 
	Fecha string
}

func (a *Alumno)detalle(){
	fmt.Printf("Nombre:\t\t%s\nApellido:\t%s\nDNI:\t\t%s\nFecha:\t\t%s\n",a.Nombre,a.Apellido,a.DNI,a.Fecha)
}

func main()  {
	Alumno1 := Alumno{"Eddy", "Carv", "646645354", "12-06-2023"}

	Alumno1.detalle()
}



