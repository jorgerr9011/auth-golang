migrate:
    docker exec wiki-golang_app_1 migrate -path=/app/migrations -database "postgres://gorm:gorm_password@db:5432/gorm?sslmode=disable" drop -f &
    docker exec wiki-golang_app_1 migrate -path /app/migrations -database "postgres://gorm:gorm_password@db:5432/gorm?sslmode=disable" up

init: 
    docker-compose up -d --build