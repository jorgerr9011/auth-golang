# Arrancar contenedores

`docker-compose up -d --build`

# 

`docker exec -it `

# Uso de golang migrate

## Para crear migraciones

* Esto crea 2 ficheros, uno up y otro down

	`docker exec -it wiki-golang_app_1 migrate create -ext sql -dir migrations -seq create_documents_table`

## Aplicar migraciones

`docker exec -it wiki-golang_app_1 migrate -path /app/migrations -database "postgres://gorm:gorm_password@db:5432/gorm?sslmode=disable" up`
