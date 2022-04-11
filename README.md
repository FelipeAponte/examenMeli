# Examen MercadoLibre
## _Ayudando a Magneto_
Magneto solicita la creación de una API REST que permita validar si una secuencia de ADN pertenece
o no a un mutante, para esto se deben cumplir con lo siguiente:

- Crear una función para analizar la secuencia de ADN. 
- Hostear una API REST con el endpoint '/mutant/' que permita recibir soliciudes HTTP POST.
- Añadir una base de datos para obtener estadísticas por medio del endpoint '/stats'.

Para probar el programa realizado por favor descarge el código fuente de [Éste Repositorio](https://github.com/FelipeAponte/examenMeli) y siga los
pasos que se mencionan a continuación.

## Prueba de la función IsMutant
Una vez ubicado en el directorio del proyecto ingrese al directorio mutant/
```sh
cd mutant/
```
Aquí podra ejecutar el test unitario del archivo mutant_test.go y ver la
cobertura de la prueba con el siguiente comando:
```sh
go test -cover .
```
También puede realizar una prueba de rendimiendo para un arreglo de 120X120
sin coincidencias de más de 4 bases nitrogenadas contiguas de forma horizontal,
vertical o diagonal, con el siguiente comando:
```sh
go test -bench .
```
## Prueba y ejecución de la API REST
Para probar el test unitario de los diferentes endpoints creados para la API REST
ubiquese en el directorio principal del proyecto, y ejecute el siguiente comando:
```sh
go test -cover .
```
Si desea ejecutar el programa localmente es necesario contar con una vase de datos
MySql y crear la siguiente tabla:
`
CREATE TABLE dnaVerified (
id BIGINT NOT NULL AUTO_INCREMENT,
dnaThreads TEXT NOT NULL, 
isMutant TINYINT(1) NOT NULL,
PRIMARY KEY (id));
`
Además debe permitir el acceso a la base de datos en donde fue creada la tabla en el
código del programa, modificando las líneas 14, 15, y 16 del archivo database.go 
que se encuentra en el directorio database/.
```go
14	USER := "<su usuario>"
15	PASSWD := "<su contraseña>"
16	DB := "<su base de datos>"
```
Una vez echo esto proceda a crear el archivo ejecutable, ubiquese en el directorio
principal del proyecto y ejecute:
```sh
go build main.go
```
Para poner en funcionamiento la API REST ejecute el siguiente comando:
```sh
./main
```
## URL de la API REST
La URL habilidata para poder hacer peticiones HTTP GET en el endpoint '/stats' y peticiones HTTP POST
al endpoint '/mutant/' para enviar secuencias de ADN por medio del body, es la siguiente:
```sh
http://ec2-54-87-45-137.compute-1.amazonaws.com:8000/stats
```
O puede utilizar directamente la dirección IP pública del servidor:
```sh
http://54.87.45.137:8000/stats
```
### Documentación de la API REST
Para poder ver la documentación de la API REST ingrese a [Swagger Editor](https://editor.swagger.io/) y copie y pegue el siguiente código YAML:
``` yaml
openapi: 3.0.0
info:
  title: Ayudando a Magneto
  description: Servicio API-REST de análisis de secuencias de ADN para reconocer mutantes.
  version: 0.1.1
servers:
  - url: http://54.87.45.137:8000
    description: Instancia EC2_AWS de servidor Ubuntu 20.04.3 LTS
paths:
  /stats:
    get:
      summary: Devuelve las estadísticas de las verificaciones de ADN realizadas.
      description: >-
        Muestra la cantidad de mutantes, cantidad de humanos y la razón entre
        ambos.
      responses:
        '200':
          description: A JSON object of details
          content:
            application/json:
              schema:
                type: object
                properties:
                  count_mutant_dna:
                    type: integer
                  count_human_dna:
                    type: integer
                  ratio:
                    type: number
  /mutant/:
    post:
      summary: Permite enviar secuencias de ADN para que sean analizadas por la API.
      description: >-
        Si la secuencia de ADN es valida, devuelve un código de estado 200 si es
        mutante o un código de estado 403 si es humano.
      responses:
        '200':
          description: Encuentra más de 1 secuencia de bases nitrogenadas contiguas de forma horizontal, vertical o diagonal
          content:
            application/json:
              examples:
                mutante:
                  value: >-
                    {"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}
        '403':
          description: Encuentra 1 o ninguna secuencia de bases nitrogenadas contiguas de forma horizontal, vertical o diagonal
          content:
            application/json:
              examples:
                mutante:
                  value: >-
                    {"dna":["ATGCGA","CAGTGC","TTATTT","AGACGG","GCGTCA","TCACTG"]}
        '400':
          description: Reconoce que la secuencia de ADN enviada no es valida, por estar vacía, contener bases nitrogenadas erroneas o tener JSON body mal formado
          content:
            application/json:
              examples:
                mutante:
                  value: >-
                    {"adn":["AAAX","CCCC","GGGG","TTTT"]}
```