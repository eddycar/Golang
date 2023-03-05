/* Calcular cantidad de alimento
Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hámsteres, perros y gatos, pero se espera que puedan darle refugio a muchos animales más.
Tienen los siguientes datos:
Perro: 10 kg de alimento.
Gato: 5 kg de alimento.
Hamster: 250 g de alimento.
Tarántula: 150 g de alimento.
Se solicita:
Implementar una función animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso de que no exista el animal).
Una función para cada animal que calcule la cantidad de alimento basándose en la cantidad del tipo de animal especificado.*/

package main

import "fmt"

func main()  {
	const (
		dog = "dog"
		cat = "cat"
		hamster = "hamster"
		tarantula = "tarantula"
	)

	animalDog, msg := animal(dog)
	if msg != ""{
		fmt.Println(msg)
	}

	animalCat, msg := animal(cat)
	if msg != ""{
		fmt.Println(msg)
	}

	animalHamster, msg := animal(hamster)
	if msg != ""{
		fmt.Println(msg)
	}

	animalTarantula, msg := animal(tarantula)
	if msg != ""{
		fmt.Println(msg)
	}
	
	var amount float32
	amount += animalDog(15)
	amount += animalCat(5)
	amount += animalHamster(5)
	amount += animalTarantula(5)
	fmt.Printf("Cantidad total de alimento: %.2f",amount)
}

	func dog(amount float32) float32{
		return amount * 10
	}
	
	func cat(amount float32) float32{
		return amount * 5
	}

	func hamster(amount float32) float32{
		return amount * 0.25
	}

	func tarantula(amount float32) float32{
		return amount * 0.15
	}

	func animal(animal string)(func(float32)float32, string) {
		switch animal {
		case "dog":
			return dog,""
		case "cat":
			return cat,""
		case "hamster":
			return hamster,""
		case "tarantula":
			return tarantula,""
		default:
			return nil, fmt.Sprint("No tenemos ese animal en el refugo")
		}
	}

