
# Inetum matriz
Este repositorio se crea para dar solucion al ejercicio de programacion del test de aptitud de INETUM

## Rutas del proyecto.
| RUTA        | METHOD                   | DESCRIPCION|
|-----------| ----------------------------------| ----------------------------------|
| /     | GET | Ruta inicial del proyecto.|
| /rotar-matriz/ | POST| Ruta que gira una matriz 90 grados en sentido anti-horario.|

## Instrucciones
Para ejecutar el proyecto se tiene como alternativa:

### Usando docker | docker-compose
Abrir la terminal, posicionarse la carpeta del proyecto y ejecutar

    docker-compose up
Para ejecutar los test.

    docker-compose exec backend go test

### Usando un compilador de go.
Abrir la terminal, posicionarse la carpeta del proyecto y ejecutar:

    go build -o <nombre ejecutable> .
Al terminar del proceso de compilacion procedemos a ejecutar nuestro binario:

    ./<nombre ejecutable>
Para ejecutar los test.

    go test

## Test adicionales
Si necesitamos probar la aplicacion desde una linea de comandos los siguientes fragmentos de codigo CURL nos ayudaran, tener en cuenta que debemos tener la aplicacion en ejecucion.

### Para matrices de 2x2

```[CURL]
curl --location --request POST 'localhost:8080/rotate-matrix/' \
--header 'Content-Type: application/json' \
--data-raw ' {
    "matrix": [
        [1,2],
        [3,4]
    ]
}
'
```
### Para matrices de 3x3
```[CURL]
curl --location --request POST 'localhost:8080/rotate-matrix/' \
--header 'Content-Type: application/json' \
--data-raw ' {
    "matrix": [
        [1,2,3],
        [4,5,6],
        [7,8,9]
    ]
}
'
```
## Tecnologias utilizadas.

 - [Go-Lang](https://golang.org/doc/)
 - [Gin-gonic](https://github.com/gin-gonic/gin)
 - [Docker](https://docs.docker.com/)
 - [Docker-Compose](https://docs.docker.com/compose/)

### Licencia GLP V3