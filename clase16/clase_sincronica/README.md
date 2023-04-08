# Problema

Un supermercado necesita un sistema para gestionar los productos frescos que tienen publicados en su página web. Para poder hacer esto, necesitan un servidor que ejecute
una API que les permita manipular los productos cargados desde distintos clientes. Los campos que conforman un producto son:

![Tabla](/img/tabla.JPG)

## Ejemplo de payload

![Tabla](/img/payload.JPG)

## Ejercicio: Método PUT

Añadir el método **PUT** a nuestra API. Recordemos que este crea o reemplaza un recurso en su totalidad con el contenido en el request. Debemos validar los campos que se envían, tal como hicimos con el método POST. **Seguimos aplicando los cambios sobre la lista cargada en memoria.**
