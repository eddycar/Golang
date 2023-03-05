/*
Una empresa de redes sociales requiere implementar una estructura usuarios con
funciones que vayan agregando información a la misma. Para optimizar y ahorrar memoria
requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del
programa y para las funciones. La estructura debe tener los campos: nombre, apellido,
edad, correo y contraseña. Y deben implementarse las funciones:
● cambiarNombre: permite cambiar el nombre y apellido.
● cambiarEdad: permite cambiar la edad.
● cambiarCorreo: permite cambiar el correo.
● cambiarContraseña: permite cambiar la contraseña.
*/

package main

import "fmt"

type usuario struct {
	nombre      string
	apellido    string
	edad        int
	correo      string
	contrasenia string
}

func (u *usuario) cambiarNombre(nombre, apellido string) {
	u.nombre = nombre
	u.apellido = apellido
}

func (u *usuario) cambiarEdad(edad int) {
	u.edad = edad
}

func (u *usuario) cambiarCorreo(correo string) {
	u.correo = correo
}

func (u *usuario) cambiarContrasenia(contrasenia string) {
	u.contrasenia = contrasenia
}

func main()  {
	usuario1 := usuario{
		nombre:      "Eddy",
		apellido:    "Carv",
		edad:        36,
		correo:      "ecarvajal@mail.com",
		contrasenia: "1234",
	}

	fmt.Println(usuario1)
	usuario1.cambiarNombre("Edilberto", "Carvajal")
	usuario1.cambiarEdad(37)
	usuario1.cambiarCorreo("edicarv@mail.com")
	usuario1.cambiarContrasenia("4321")
	fmt.Println(usuario1)
}
