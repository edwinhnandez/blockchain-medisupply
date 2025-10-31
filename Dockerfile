# Build stage
FROM golang:1.21-alpine AS builder

# Instalar dependencias de compilación
RUN apk add --no-cache git gcc musl-dev

# Establecer directorio de trabajo
WORKDIR /app

# Copiar go mod y sum
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar código fuente
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

# Runtime stage
FROM alpine:latest

# Instalar ca-certificates para HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar el binario desde el build stage
COPY --from=builder /app/main .

# Exponer puerto
EXPOSE 8080

# Comando para ejecutar
CMD ["./main"]

