# Arrancar contenedores

## Desarrollo
	`docker-compose up -d --build`

## Producción

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

