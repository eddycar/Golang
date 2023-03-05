/*Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
Tener un slice global de Product llamado Products, instanciado con valores.
Dos métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método. El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
Una función getById() a la cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
Ejecutar al menos una vez cada método y función definidos desde main().*/

package main

import "fmt"

type Product struct{
	ID int 
	Name string 
	Price float32 
	Description string 
	Category string
}

var Products = []Product{
	{1, "producto1", 12, "descripcion", "categoria"},
	{2, "producto2", 12, "descripcion", "categoria"},
	{3, "producto3", 12, "descripcion", "categoria"},
}

func (p *Product) Save(){
	Products = append(Products, *p)
}

func (p *Product) GetAll()[]Product{
	return Products
}

func getById(id int) Product{
	for i,producto := range Products{
		if producto.ID == id{
			return Products[i]
		}
	}
 return Product{}
}

func main(){
	
	 p1 := Product{
		ID: 4,
		Name: "Producto 4",
		Price: 14.5,
		Description: "Descripcion",
		Category: "categoria",
	 }

	 p1.Save()
	 fmt.Println("")
	 fmt.Println(p1.GetAll())
	 fmt.Println("")
	 fmt.Println(getById(3)) 
}