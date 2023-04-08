/* Lista de productos
Vamos a crear una aplicación web con el framework Gin que tenga un endpoint /productos
que responda con una lista de productos.
1. Crear una estructura Productos con los valores:
● Id
● Nombre
● Precio
● Stock
● Codigo
● Publicado
● FechaDeCreacion
2. Crear un archivo productos.json con al menos seis productos (deben seguir la
misma estructura ya mencionada).
3. Leer el archivo productos.json.
4. Imprimir la lista de productos por consola.
5. Imprimir la lista de productos a través del endpoint en formato JSON. El endpoint
deberá ser de método GET.
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	ID             int     `json:"id"`
    Nombre         string  `json:"nombre"`
    Precio         float64 `json:"precio"`
    Stock          bool    `json:"stock"`
    Codigo         string  `json:"codigo"`
    Publicado      bool    `json:"publicado"`
    FechaDeCreacion string `json:"FechaDeCreacion"`
}

func main() {

	r := gin.Default()

	file, err := os.ReadFile("productos.json")
	if err != nil{
		fmt.Println(err)
		return
	}

	var producto []Producto

	err2 := json.Unmarshal(file, &producto)
	if err != nil {
		fmt.Println(err2)
		return
	}

	r.GET("/productos", func(c *gin.Context)  {
		c.JSON(http.StatusOK, producto)
	} )

	r.Run()
}
