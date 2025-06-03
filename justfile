migrate:
    docker exec auth-golang-app migrate -path=/app/migrations -database "postgres://gorm:gorm_password@db:5432/gorm?sslmode=disable" drop -f
    docker exec auth-golang-app migrate -path /app/migrations -database "postgres://gorm:gorm_password@db:5432/gorm?sslmode=disable" up

migrate-production:
    docker exec auth-golang-app migrate -path=/app/migrations -database "postgres://gorm:gorm_password@db:5432/authdb?sslmode=disable" drop -f
    docker exec auth-golang-app migrate -path /app/migrations -database "postgres://gorm:gorm_password@db:5432/authdb?sslmode=disable" up

init: 
    docker-compose up -d --build