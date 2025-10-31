.PHONY: help build run test clean docker-build docker-up docker-down docker-logs install-deps

help: ## Muestra esta ayuda
	@echo "Comandos disponibles:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

install-deps: ## Instala dependencias de Go
	@echo "Instalando dependencias..."
	go mod download
	go mod tidy

build: ## Compila la aplicación
	@echo "Compilando aplicación..."
	go build -o bin/medisupply cmd/api/main.go

run: ## Ejecuta la aplicación localmente
	@echo "Iniciando aplicación..."
	go run cmd/api/main.go

test: ## Ejecuta tests unitarios
	@echo "Ejecutando tests..."
	go test -v ./tests/...

test-coverage: ## Ejecuta tests con coverage
	@echo "Ejecutando tests con coverage..."
	go test -coverprofile=coverage.out ./tests/...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generado en coverage.html"

test-short: ## Ejecuta solo tests unitarios (skip integración)
	@echo "Ejecutando tests unitarios..."
	go test -short -v ./tests/...

lint: ## Ejecuta linter
	@echo "Ejecutando linter..."
	golangci-lint run

docker-build: ## Construye imagen Docker
	@echo "Construyendo imagen Docker..."
	docker-compose build

docker-up: ## Inicia servicios con Docker Compose
	@echo "Iniciando servicios..."
	docker-compose up -d
	@echo "Servicios iniciados"
	@echo "API: http://localhost:8080"
	@echo "IPFS Gateway: http://localhost:8081"

docker-up-local: ## Inicia servicios incluyendo DynamoDB local
	@echo "Iniciando servicios (con DynamoDB local)..."
	docker-compose --profile local up -d

docker-down: ## Detiene servicios Docker
	@echo "Deteniendo servicios..."
	docker-compose down

docker-logs: ## Muestra logs de Docker Compose
	docker-compose logs -f

docker-clean: ## Limpia contenedores, imágenes y volúmenes
	@echo "Limpiando Docker..."
	docker-compose down -v
	docker system prune -f

setup-dynamodb: ## Crea tabla en DynamoDB
	@echo "Creando tabla en DynamoDB..."
	aws dynamodb create-table \
		--table-name transacciones-blockchain \
		--attribute-definitions AttributeName=idTransaction,AttributeType=S \
		--key-schema AttributeName=idTransaction,KeyType=HASH \
		--billing-mode PAY_PER_REQUEST \
		--region us-east-1

setup-dynamodb-local: ## Crea tabla en DynamoDB local
	@echo "Creando tabla en DynamoDB local..."
	aws dynamodb create-table \
		--table-name transacciones-blockchain \
		--attribute-definitions AttributeName=idTransaction,AttributeType=S \
		--key-schema AttributeName=idTransaction,KeyType=HASH \
		--billing-mode PAY_PER_REQUEST \
		--endpoint-url http://localhost:8000 \
		--region us-east-1

clean: ## Limpia archivos generados
	@echo "Limpiando archivos generados..."
	rm -rf bin/
	rm -f coverage.out coverage.html
	go clean

fmt: ## Formatea código Go
	@echo "Formateando código..."
	go fmt ./...

vet: ## Ejecuta go vet
	@echo "Ejecutando go vet..."
	go vet ./...

mod-tidy: ## Limpia dependencias no utilizadas
	@echo "Limpiando módulos..."
	go mod tidy

check: fmt vet lint test ## Ejecuta todas las verificaciones

dev: docker-up run ## Inicia dependencias y ejecuta en modo desarrollo

health-check: ## Verifica estado de la API
	@echo "Verificando salud de la API..."
	@curl -s http://localhost:8080/health | jq '.'

ready-check: ## Verifica readiness de la API
	@echo "Verificando readiness..."
	@curl -s http://localhost:8080/ready | jq '.'

