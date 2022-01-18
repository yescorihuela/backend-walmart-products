# Hola a todos => API de Promociones y Productos

Bienvenidos a la API de productos / promociones, acá vamos a tratar los resultados, las decisiones de ingeniería y lo que pudo haber quedado pendiente. Es un deleite para mi programar en Go, vuelvo a mis inicios de las _curly braces_ => "{}". Qué me gustó de esto? La implementación de MongoDB, aunque debo ser sincero: estuve largo rato de automatizar la carga de datos con mi propio docker-compose.yml en el backend sin tener que alterar mucho o prácticamente nada el archivo `import.sh`, así que empecemos...

## Tecnologías presentes

- Go 1.17.6
- Docker / Docker Compose
- MongoDB 3.6.8
- gin-gonic como router de Go

Usar **gin-gonic** tiene sus ventajas al no tener que estar creando implementaciones para el encoding/decoding de JSON en las responses, ni tampoco preocuparse de los parámetros de los requests, solamente ocuparse de validarlos (en este caso no tenemos llamadas con el verbo HTTP POST para ese fin, es un dato adicional). No utilizamos otra tecnología más allá a las utilizadas frecuentemente por la comunidad de Golang.

## Arquitectura

Al ser un gusto la programación en Go (Golang para los amigos), en base a los diversos materiales que he leído y los incontables videos que he visto sobre **Arquitectura Hexagonal**, decidí utilizarla con Golang, dado a que se da superbien, es una arquitectura también llamada **Ports and Adapters** (Puertos y Adaptadores) y es un subconjunto de la denominada **_Clean Architecture_**. La Arquitectura Hexagonal se fundamenta en la **Dependency Injection** y en un intensivo uso de interfaces, para poder intercambiar diversas implementaciones, siempre y cuando satisfagan la interfaz respetando los métodos en ella definidos, junto a eso se suma el SRP o **_Single Reponsability Principle_** de **SOLID**, tratando de separar funcionalidades, también el uso intensivo de interfaces conlleva a la buena aplicación del principio IS o **_Interfaces Segregation_**. Incluiré también un diagrama que tengo de hace tiempo y utilicé en otro proyecto (público).
Otro dato intersante: he manejado los responses a través de DTO's o **Data Transfer Object** a fin de hacer los cálculos pertinentes a este desarrollo para descuentos y especialmente los **_inputs palíndromos_**.

## Diagrama

![Arquitectura Hexagonal](Hexagonal-Architecture-Products.svg "HA Diagram")

## Ejecución del proyecto

```bash
# Ejecución en base a docker-compose por primera vez en primer plano:
$ docker compose up --build # Creará las imágenes y levantará los servicios de una vez

# Para que quede ejecutando en segundo plano desde la primera vez:
$ docker compose up --build -d # d de daemon

# Para bajar los servicios, construirlo y levantar servicios nuevamente
$ docker compose down && docker compose up --build

# Ejecución de los tests

# Primero acceder al contenedor de la app
$ docker exec -it backend-products /bin/bash

# Segundo, al estar en el prompt del root y en el orden que deseen:

go test ./... -cover # por si quieren ver el coverage
go test ./api -cover # para ver por capa
go test ./services -cover # para ver por capa
go test ./handlers -cover # para ver por capa
```

## Endpoints

| Endpoint                    | Verbo HTTP | Explicación                                                                                    | Respuestas                                |
|-----------------------------|------------|------------------------------------------------------------------------------------------------|-------------------------------------------|
| localhost:8082/v1/products                   | GET        | Trae todos los productos (sin paginar)                                                         | 200 OK \| Array of products 404 Not found |
| localhost:8082/v1/products/search?q=criteria | GET        | Trae todos los productos (sin paginar) en base al criterio de búsqueda definido por el usuario | 200 OK \| Array of products 400 Not found |

### Quedaría pendiente

- **Utilización del logger presente:** (dado a que ya está implementado), fue extendido en los niveles de logging dado a que el original (Gracias a [Alex Edwards](https://www.alexedwards.net/about)). Un logger más estructurado (en JSON) para mejorar el input de los servicios de observabilidad.

- **Paginación de los resultados:**, aunque quedara fuera del scope de este proyecto, es necesario muchas veces, localmente MongoDB vuela en devolver los resultados y Golang tiene un performance de lujo, pero en producción con datos reales y gran número de transacciones, las cosas cambian mucho.

- Por último pero no por ello, menos importante: un **_GRACEFUL SHUTDOWN_** del servicio de la API, eso está documentado en algunos lugares, ([Alex Edwards](https://www.alexedwards.net/about) lo explica también en su libro), para que tome las señales del teclado o del sistema y termine de drenar las peticiones o servir las peticiones pendientes (muy importante en producción).

**Esto es todo por ahora!**