# Arrancar contenedores

`docker-compose up -d --build`

# 

`docker exec -it `

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

