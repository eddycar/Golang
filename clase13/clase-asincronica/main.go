/* Ejercicio: Hola mundo
Vamos a crear una aplicación web con el framework Gin que tenga un endpoint /hola-mundo que responda con un mensaje.
Tener en cuenta que:
El endpoint deberá ser de método GET.
La respuesta deberá ser recibida en formato JSON. Ejemplo:
{
    "mensaje": "¡Hola Mundo!"
}

*/

package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {

	router := gin.Default()

	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensaje": "Hola Mundo!",
		})
	})
 
	router.Run()
}





