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

clean-containers:
    docker compose -f docker-compose.yml down
    
destroy:
    docker compose -f docker-compose.yml down -v --remove-orphans

clean-images:
    docker image prune -af --filter label=project=auth-golang

clean-stopped-containers:
    docker container prune -f

clean-unused-volumes:
    docker volume prune -f

clean-unused-networks:
    docker network prune -f

clean-all:
    docker compose -f docker-compose.yml down -v --remove-orphans
    docker image prune -af
    docker container prune -f
    docker volume prune -f
    docker network prune -f