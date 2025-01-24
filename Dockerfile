# Etapa 1: Construcción de la aplicación
FROM golang:1.22-alpine as builder

# Establecer el directorio de trabajo en /app
WORKDIR /app

# Copiar todo el contenido al contenedor
COPY . .

# Instalar las dependencias de Go
RUN go mod tidy

# Construir la aplicación para un entorno Linux (amd64)
RUN GOOS=linux GOARCH=amd64 go build -o app ./cmd/server/main.go

# Etapa 2: Imagen más ligera para producción
FROM alpine:latest

# Establecer el directorio de trabajo en /root
WORKDIR /root/

# Copiar el binario desde la etapa de construcción
COPY --from=builder /app/app .

# Exponer el puerto de la aplicación
EXPOSE 8080

# Ejecutar la aplicación
CMD ["./app"]
