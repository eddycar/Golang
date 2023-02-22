package listar_nombres

import "fmt"

func ListarEstudiantes(lista []string) {
	for i, estudiante := range lista {
		fmt.Printf("%d\t%s\n",i+1, estudiante)
	}
}