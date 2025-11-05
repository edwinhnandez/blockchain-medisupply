# Build stage
FROM golang:1.21-alpine AS builder

# Instalar dependencias de compilación
RUN apk add --no-cache git gcc musl-dev ca-certificates

# Configurar Go proxy y timeouts para evitar problemas de red
ENV GOPROXY=https://proxy.golang.org,direct
ENV GOPRIVATE=
ENV GOSUMDB=sum.golang.org
ENV GOPROXY_TIMEOUT=300

# Establecer directorio de trabajo
WORKDIR /app

# Copiar go mod y sum
COPY go.mod go.sum ./

# Descargar dependencias con retry y mejor manejo de errores
RUN go mod download || (sleep 5 && go mod download) || (sleep 10 && go mod download)

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

