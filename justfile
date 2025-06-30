# Carga automáticamente .env
set dotenv-load := true

# Muestra los comandos disponibles
default:
    @echo ""
    @echo "Comandos disponibles:"
    @echo ""
    @echo "  just install               Obtiene imágenes docker y despliega el proyecto"
    @echo "  just init                  Despliega el proyecto"
    @echo "  just migrate               Ejecuta las migraciones"
    @echo "  just migrate-production    Ejecuta las migraciones en producción"

install: 
    docker compose -f docker-compose.yml up --build -d

init: 
    docker compose -f docker-compose.yml up -d

migrate:
    docker exec -it auth-golang-app migrate -path=/app/migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" drop -f
    docker exec -it auth-golang-app migrate -path /app/migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up
