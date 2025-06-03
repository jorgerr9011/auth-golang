FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth-app cmd/api/main.go

FROM alpine:latest

RUN apk add --no-cache bash curl tar just

RUN adduser -D -g '' appuser

WORKDIR /app

RUN mkdir -p /app/logs && chmod 777 /app/logs

COPY --from=builder /app/auth-app .
COPY entrypoint.sh .

RUN chmod +x entrypoint.sh

COPY justfile . 

USER appuser

EXPOSE 8080

CMD ["./auth-app"]