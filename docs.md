## INTRODUCCIÓN

cmd/server/main.go: 

- Este archivo es el punto de entrada de la aplicación. Aquí configurarás el servidor y las rutas.

go.mod / go.sum: 

- Archivos que gestionan las dependencias de tu proyecto.


Teóricamente la estructura de Golang y Gin para creación de REST APIs es la siguiente:

- pkg/handler:

    * Controladores de las rutas

- pkg/service:

    * Lógica de negocio

- pkg/model:

    * Definición de los modelos (estructura de datos)

- pkg/repository:

    * Interacción con la base de datos

- /internal:

    * Lógica de negocio de la aplicación (código interno)

- /scripts:

    * Scripts útiles para el proyecto (si los hay)

- /migrations:

    * Scripts de migraciones para la base de datos

- /web:

    * Archivos estáticos o de frontend (si los hay)


### Base de datos

Voy a utilizar para la base de datos el ORM Gorm, y como base de datos PostgreSQL.