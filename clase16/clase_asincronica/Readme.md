## Problema
Un supermercado necesita un sistema para gestionar los productos frescos que tienen publicados en su página web. Para poder hacer esto, necesitan un servidor que ejecute una API que les permita manipular los productos cargados desde distintos clientes. 
Los campos que conforman un producto son:

![Tabla](/img/tabla.JPG)

### Ejercicio 1: Añadir un producto
Ahora, vamos a añadir un producto al slice cargado en memoria. Dentro de la ruta **/products** añadimos el método **POST**, al cual vamos a enviar en el cuerpo del request el nuevo producto. Este tiene ciertas restricciones, conozcámoslas:
1) No es necesario pasar el **id**, al momento de añadirlo se debe inferir del estado de la lista de productos, verificando que no se repitan, ya que debe ser un campo único.

2) Ningún dato puede estar **vacío**, exceptuando **is_published** (vacío indica un valor **false**).

3) El campo **code_value** debe ser único para cada producto.

4) Los **tipos de datos** deben coincidir con los definidos en el planteo del problema.

5) La fecha de vencimiento debe tener el formato: **XX/XX/XXXX**, además debemos verificar que **día**, **mes** y **año** sean valores válidos.

**Recordemos**: si una consulta está **mal formulada** por parte del **cliente**, el **status code** cae en los **4XX**. 

### Ejercicio 2: Traer el producto
Realizar una consulta a un método **GET** con el **id** del producto recién añadido. Tener en cuenta que la lista de productos se encuentra cargada en la **memoria**. Si terminamos la ejecución del programa, **este producto no estará en la próxima ejecución.**