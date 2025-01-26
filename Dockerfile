# Imagen base con Go y herramientas necesarias
FROM golang:1.23-alpine

# Instala migrate con soporte para PostgreSQL
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Configuración del directorio de trabajo
WORKDIR /app

# Copiar los archivos de configuración de dependencias
COPY go.mod go.sum ./

# Descargar las dependencias del proyecto
RUN go mod download

# Como ya me monto un volumen con los archivos no hace falta
# Copiar todo el código fuente
#COPY . .

# Exponer el puerto de la aplicación
EXPOSE 8080

# Comando para compilar y ejecutar la aplicación
CMD ["sh", "-c", "go build -o main cmd/api/main.go && ./main"]