/*-----employee
Una empresa necesta realizar una buena gestión de sus empleados, para esto realizamos un pequeño 
programa nos ayudará a gestionar correctamente dichos empleados, Los objetivos son:
	*Crear una estructura Person con los campos ID, Name, DateBirth.
	*Crear una estructura Employee con los campos: ID, Position y una composición con laa estructura Person.
	*Realizar un Método a la estructura Employee que se llame PrintEmployee(), lo que hará es realizar la 
	impresión  de los campos de empleado.
	*Instanciar en la función main() tanto una Person como un Employee cargando sus respectivos campos y,
	por último, ejecutar el método PrintEmployee().
Si logramos realizar este pequeño programa. pudimos ayudar a la empresa a solucionar la gestión de los 
empleados.
*/

package main

import "fmt"

type Person struct {
	ID int
	Name string
	DateOfBirth string
}

type Employee struct{
	ID int
	Position string
	Person Person
}

func main(){
	employee1 := Employee{
		ID: 1,
		Position: "Developer",
		Person: Person{
			ID: 1,
			Name:"Eddy",
			DateOfBirth: "17-10-1997",
		},
	}
	employee1.PrintEmployee()
	employee1.getter()
	employee1.setter(4)
	employee1.PrintEmployee()
}


func (e Employee)getter()Employee {
	return e
}

func (e *Employee)setter(id int){
	e.ID = id
}

func (e Employee) PrintEmployee(){
	fmt.Println(e)
}

