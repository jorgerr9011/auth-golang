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

# Despliegue

## Desarrollo
	`docker-compose up -d --build`

## Producción

	En caso de querer correr instalación y migraciones, copiar .env.production a fichero
	.env para cargar las variables de producción y poder aplicar los comandos de justfile 
	con esas variables de entorno:

	`cp .env.development .env` # desarrollo

	`cp .env.production .env` #producción

	`just install`

	`just migrate`

	`docker-compose -f docker-compose.production.yml --env-file .env.production up -d`

* Para construir la imagen en producción:

	`docker build -t jorgerr9011/proyectos:auth-golang-app_latest .`

* Para subir la imagen al repo:

	`docker push jorgerr9011/proyectos:auth-golang-app_latest`

* Ejecutar app en producción:

	`docker-compose -f docker-compose.production.yml --env-file .env.production up -d`

# Uso de golang migrate

## Para crear migraciones

* Esto crea 2 ficheros, uno up y otro down

		`docker exec -it auth-golang-app migrate create -ext sql -dir migrations -seq create_users_table`

## Aplicar migraciones

	`docker exec -it auth-golang-app migrate -path /app/migrations -database "postgres://gorm:gorm_password@db:5432/gorm?sslmode=disable" up`

# Realizar un fresh de las migraciones

	`docker exec -it auth-golang-app migrate -path=/app/migrations -database "postgres://gorm:gorm_password@db:5432/gorm?sslmode=disable" drop -f`

## Seeders

* Para ejecutar los seeders de manera manual:

		`go run `

## Autenticación:

- JWT