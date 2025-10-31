# Guía de Instalación Detallada

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://golang.org)
[![Docker](https://img.shields.io/badge/Docker-24+-2496ED?style=flat&logo=docker)](https://www.docker.com)
[![Last Updated](https://img.shields.io/badge/Updated-2025-brightgreen)](https://github.com/edinfamous/blockchain-medisupply)

Esta guía te llevará paso a paso a través de la instalación completa del sistema de trazabilidad blockchain para medicamentos. Actualizada para 2025 con las mejores prácticas de la industria.

## Tabla de Contenidos

- [Requisitos del Sistema](#requisitos-del-sistema)
- [Instalación Rápida (5 minutos)](#-instalación-rápida-5-minutos)
- [Instalación Paso a Paso](#instalación-paso-a-paso)
- [Configuración por Entorno](#-configuración-por-entorno)
- [Verificación y Validación](#-verificación-y-validación)
- [Solución de Problemas](#-solución-de-problemas)
- [Instalación Avanzada](#-instalación-avanzada)
- [Actualización de Versiones](#-actualización-de-versiones)
- [Desinstalación](#-desinstalación)

## Requisitos del Sistema

### Software Necesario (2025)

| Software | Versión Mínima | Versión Recomendada | Propósito |
|----------|---------------|---------------------|-----------|
| **Go** | 1.23+ | 1.23+ | Runtime y desarrollo |
| **Docker** | 24.0+ | 27.0+ | Containerización |
| **Docker Compose** | 2.20+ | 2.29+ | Orquestación |
| **AWS CLI** | 2.13+ | 2.20+ | Gestión de AWS |
| **Make** | 4.0+ | 4.4+ | Build automation |
| **Git** | 2.40+ | 2.47+ | Control de versiones |

**Enlaces de descarga:**
- Go: [https://go.dev/dl/](https://go.dev/dl/)
- Docker Desktop: [https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/)
- AWS CLI: [https://aws.amazon.com/cli/](https://aws.amazon.com/cli/)

### Herramientas Opcionales (Recomendadas)

| Herramienta | Propósito | Instalación |
|-------------|-----------|-------------|
| **jq** | Formateo JSON | `brew install jq` (Mac) / `apt install jq` (Linux) |
| **curl** | Testing API | Incluido en la mayoría de OS |
| **wget** | Descarga de archivos | `brew install wget` / `apt install wget` |
| **golangci-lint** | Linting de código | [https://golangci-lint.run/](https://golangci-lint.run/) |
| **gosec** | Security scanning | `go install github.com/securego/gosec/v2/cmd/gosec@latest` |
| **govulncheck** | Vulnerability check | `go install golang.org/x/vuln/cmd/govulncheck@latest` |

### Cuentas y Servicios

#### Requeridos

1. **AWS Account** - Para DynamoDB
   - Crear cuenta en [aws.amazon.com](https://aws.amazon.com)
   - Configurar IAM user con permisos de DynamoDB
   - Habilitar MFA para seguridad
   - Crear access keys para desarrollo

#### Opcionales (Blockchain)

2. **Alchemy Account** (Recomendado 2025)
   - [Alchemy](https://www.alchemy.com/) - Opción recomendada 2025 ⭐
   - [Infura](https://infura.io) - También soportado
   - [QuickNode](https://www.quicknode.com/) - Alternativa rápida

3. **Ethereum Testnet Funds**
   - [Sepolia Faucet](https://sepoliafaucet.com/)
   - [Alchemy Sepolia Faucet](https://sepoliafaucet.com/)
   - [QuickNode Faucet](https://faucet.quicknode.com/)

### Requisitos de Hardware

| Entorno | CPU | RAM | Disco | Notas |
|---------|-----|-----|-------|-------|
| **Desarrollo** | 2 cores | 4 GB | 20 GB | Mínimo para local |
| **Testing** | 4 cores | 8 GB | 50 GB | CI/CD pipelines |
| **Producción** | 4+ cores | 8+ GB | 100+ GB | Depende de carga |

### Sistemas Operativos Soportados

- macOS 12+ (Monterey, Ventura, Sonoma)
- Ubuntu 20.04/22.04/24.04 LTS
- Debian 11/12
- Windows 10/11 + WSL2
- Amazon Linux 2023
- Windows nativo (con limitaciones)

## ⚡ Instalación Rápida (5 minutos)

Para usuarios experimentados que quieren empezar rápidamente:

```bash
# 1. Clonar y entrar al proyecto
git clone https://github.com/edinfamous/blockchain-medisupply.git
cd blockchain-medisupply

# 2. Copiar y configurar variables de entorno
cp env.example .env
# Editar .env con tus credenciales

# 3. Iniciar servicios con Docker
docker-compose up -d

# 4. Verificar que todo funciona
curl http://localhost:8080/health

# 5. Ejecutar tests
go test -short ./tests/...
```

**¡Listo!** Tu instalación está completa. Ir a [Verificación](#-verificación-y-validación) para probar más funcionalidades.

## Instalación Paso a Paso

### 1. Verificar Prerrequisitos

```bash
# Verificar versiones instaladas
go version          # debe ser 1.23+
docker --version    # debe ser 24.0+
docker compose version  # debe ser 2.20+
aws --version       # debe ser 2.13+
make --version      # debe ser 4.0+

# Verificar Docker está corriendo
docker ps

# Verificar permisos (Linux/Mac)
docker run hello-world
```

**Salida esperada:**
```
go version go1.23.0 darwin/arm64
Docker version 27.0.3, build 7d4bcd8
Docker Compose version 2.29.0
aws-cli/2.20.0 Python/3.11.8
GNU Make 4.4.1
```

### 2. Clonar el Repositorio

```bash
# Clonar vía HTTPS
git clone https://github.com/edinfamous/blockchain-medisupply.git
cd blockchain-medisupply

# O vía SSH (recomendado para colaboradores)
git clone git@github.com:edinfamous/blockchain-medisupply.git
cd blockchain-medisupply

# Verificar estructura
ls -la

# Verificar rama actual
git branch
```

### 3. Instalar Dependencias de Go

#### Método Recomendado (2025):

```bash
# Limpiar caché si hay problemas
go clean -modcache

# Descargar dependencias
go mod download

# Verificar y limpiar dependencias
go mod tidy

# Verificar que todas las dependencias estén correctas
go mod verify
```

**Salida esperada:**
```
all modules verified
```

#### Solución de Problemas de Red:

```bash
# Opción 1: Usar proxy alternativo (China, etc.)
export GOPROXY=https://goproxy.io,direct
go mod download

# Opción 2: Usar proxy de Google
export GOPROXY=https://proxy.golang.org,direct
go mod download

# Opción 3: Múltiples intentos con retry
for i in {1..3}; do
  echo "Intento $i..."
  go mod download && break || sleep 5
done

# Opción 4: Saltar (usar Docker que descargará internamente)
echo "Las dependencias se descargarán en el build de Docker"
```

#### Verificar Instalación de Dependencias:

```bash
# Listar dependencias principales
go list -m all | head -20

# Verificar dependencias críticas
go list -m github.com/gin-gonic/gin
go list -m github.com/ethereum/go-ethereum
go list -m github.com/aws/aws-sdk-go-v2

# Compilar para verificar
go build -o /tmp/test-build cmd/api/main.go
echo "Compilación exitosa!"
rm /tmp/test-build
```

### 4. Configurar Variables de Entorno

#### Generar Encryption Key Segura (2025)

```bash
# Método 1: OpenSSL (recomendado)
openssl rand -base64 32 | cut -c1-32

# Método 2: Python
python3 -c "import secrets; print(secrets.token_urlsafe(32)[:32])"

# Método 3: Go
go run -<<'EOF'
package main
import ("crypto/rand"; "encoding/base64"; "fmt")
func main() {
    b := make([]byte, 32)
    rand.Read(b)
    fmt.Println(base64.URLEncoding.EncodeToString(b)[:32])
}
EOF

# Guardar la clave generada para usarla más adelante
```

#### Copiar y Configurar .env

```bash
# Copiar template
cp env.example .env

# Abrir con tu editor preferido
code .env        # VS Code
nano .env        # Terminal
vim .env         # Vim
subl .env        # Sublime Text
```

#### Configuración Mínima (Desarrollo Local)

**Ideal para empezar sin blockchain:**

```bash
# ================== ENCRYPTION ==================
# CRÍTICO: Debe ser exactamente 32 caracteres
ENCRYPTION_KEY=YOUR_32_CHARACTER_KEY_HERE!!!

# ================== AWS ==================
AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
AWS_REGION=us-east-1
DYNAMODB_TABLE_NAME=transacciones-blockchain-dev

# ================== IPFS ==================
IPFS_HOST=ipfs
IPFS_PORT=5001

# ================== SERVER ==================
SERVER_PORT=8080
GIN_MODE=debug
LOG_LEVEL=debug

# ================== RATE LIMITING ==================
RATE_LIMIT_REQUESTS=100
RATE_LIMIT_WINDOW=60
```

#### Configuración Completa (Con Blockchain)

**Para funcionalidad completa incluyendo blockchain:**

```bash
# Todo lo anterior +

# ================== BLOCKCHAIN (ALCHEMY - 2025) ==================
ALCHEMY_API_KEY=your_alchemy_api_key

# O usar URL RPC personalizada (opcional):
# BLOCKCHAIN_RPC_URL=https://eth-sepolia.g.alchemy.com/v2/YOUR_API_KEY
# También soporta Infura:
# BLOCKCHAIN_RPC_URL=https://sepolia.infura.io/v3/YOUR_PROJECT_ID

BLOCKCHAIN_NETWORK=sepolia
BLOCKCHAIN_PRIVATE_KEY=your_private_key_without_0x_prefix
CONTRACT_ADDRESS=0x0000000000000000000000000000000000000000

# Opcionales blockchain
GAS_LIMIT=300000
MAX_GAS_PRICE=50  # en Gwei
BLOCKCHAIN_TIMEOUT=30  # segundos
```

#### Configuración Producción (2025)

**Para entornos de producción:**

```bash
# ================== SECURITY ==================
USE_AWS_SECRETS=true
BLOCKCHAIN_PRIVATE_KEY_SECRET=prod/blockchain/private-key
ENCRYPTION_KEY_SECRET=prod/app/encryption-key

# ================== SERVER ==================
GIN_MODE=release
LOG_LEVEL=info
SERVER_PORT=8080
ENABLE_CORS=true
CORS_ORIGINS=https://app.example.com,https://dashboard.example.com

# ================== MONITORING ==================
ENABLE_METRICS=true
METRICS_PORT=9090
OTEL_EXPORTER_OTLP_ENDPOINT=https://otel-collector:4317
SENTRY_DSN=https://your-sentry-dsn

# ================== PERFORMANCE ==================
MAX_CONCURRENT_REQUESTS=1000
DB_CONNECTION_POOL_SIZE=50
IPFS_TIMEOUT=30
CACHE_TTL=300  # segundos

# ================== HIGH AVAILABILITY ==================
HEALTH_CHECK_INTERVAL=10
GRACEFUL_SHUTDOWN_TIMEOUT=30
ENABLE_PROFILING=false
```

#### Validar Configuración

```bash
# Verificar que el archivo .env existe
test -f .env && echo ".env existe" || echo ".env no existe"

# Verificar que ENCRYPTION_KEY tiene 32 caracteres
KEY=$(grep ENCRYPTION_KEY .env | cut -d'=' -f2)
if [ ${#KEY} -eq 32 ]; then
    echo "ENCRYPTION_KEY tiene 32 caracteres"
else
    echo "ENCRYPTION_KEY debe tener exactamente 32 caracteres (actual: ${#KEY})"
fi

# Verificar variables críticas están definidas
for var in ENCRYPTION_KEY AWS_ACCESS_KEY_ID AWS_REGION DYNAMODB_TABLE_NAME; do
    grep -q "^${var}=" .env && echo "$var configurado" || echo " $var falta"
done
```

### 5. Configurar AWS

#### A. Crear Usuario IAM

```bash
# 1. Crear usuario IAM con CLI
aws iam create-user --user-name blockchain-medisupply-dev

# 2. Crear política personalizada (mínimo privilegio)
cat > /tmp/dynamodb-policy.json <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "dynamodb:GetItem",
        "dynamodb:PutItem",
        "dynamodb:UpdateItem",
        "dynamodb:Query",
        "dynamodb:Scan",
        "dynamodb:DescribeTable"
      ],
      "Resource": "arn:aws:dynamodb:*:*:table/transacciones-blockchain*"
    }
  ]
}
EOF

# 3. Crear política
aws iam create-policy \
  --policy-name BlockchainMediSupplyDynamoDBPolicy \
  --policy-document file:///tmp/dynamodb-policy.json

# 4. Adjuntar política al usuario
aws iam attach-user-policy \
  --user-name blockchain-medisupply-dev \
  --policy-arn arn:aws:iam::YOUR_ACCOUNT_ID:policy/BlockchainMediSupplyDynamoDBPolicy

# 5. Crear access keys
aws iam create-access-key --user-name blockchain-medisupply-dev

# 6. Habilitar MFA (recomendado para producción)
aws iam enable-mfa-device --user-name blockchain-medisupply-dev \
  --serial-number arn:aws:iam::YOUR_ACCOUNT_ID:mfa/device \
  --authentication-code1 123456 \
  --authentication-code2 789012
```

#### B. Configurar Credenciales

```bash
# Método 1: AWS CLI Configure (Interactivo)
aws configure
# AWS Access Key ID: [Ingresar]
# AWS Secret Access Key: [Ingresar]
# Default region: us-east-1
# Default output format: json

# Método 2: Configurar perfil específico
aws configure --profile blockchain-medisupply
# Usar este perfil: export AWS_PROFILE=blockchain-medisupply

# Método 3: Variables de entorno (ya en .env)
# Las variables en .env se cargarán automáticamente

# Verificar configuración
aws sts get-caller-identity
aws sts get-caller-identity --profile blockchain-medisupply  # si usas perfil
```

**Salida esperada:**
```json
{
    "UserId": "AIDAXXXXXXXXXXXXX",
    "Account": "123456789012",
    "Arn": "arn:aws:iam::123456789012:user/blockchain-medisupply-dev"
}
```

#### C. Crear Tabla DynamoDB

**Opción 1: DynamoDB en AWS (Producción/Staging)**

```bash
# Crear tabla con PAY_PER_REQUEST (recomendado 2025)
aws dynamodb create-table \
  --table-name transacciones-blockchain \
  --attribute-definitions \
    AttributeName=idTransaction,AttributeType=S \
    AttributeName=idProducto,AttributeType=S \
    AttributeName=timestamp,AttributeType=N \
  --key-schema \
    AttributeName=idTransaction,KeyType=HASH \
  --global-secondary-indexes \
    "IndexName=ProductoIndex,KeySchema=[{AttributeName=idProducto,KeyType=HASH},{AttributeName=timestamp,KeyType=RANGE}],Projection={ProjectionType=ALL}" \
  --billing-mode PAY_PER_REQUEST \
  --table-class STANDARD \
  --tags \
    Key=Environment,Value=production \
    Key=Project,Value=blockchain-medisupply \
  --region us-east-1

# Habilitar Point-in-Time Recovery (backup automático)
aws dynamodb update-continuous-backups \
  --table-name transacciones-blockchain \
  --point-in-time-recovery-specification PointInTimeRecoveryEnabled=true

# Habilitar encryption at rest
aws dynamodb update-table \
  --table-name transacciones-blockchain \
  --sse-specification Enabled=true,SSEType=KMS
```

**Opción 2: DynamoDB Local (Desarrollo)**

```bash
# 1. Iniciar DynamoDB local con Docker
docker-compose --profile local up -d dynamodb-local

# 2. Esperar a que esté listo
sleep 5

# 3. Crear tabla en local
aws dynamodb create-table \
  --table-name transacciones-blockchain \
  --attribute-definitions \
    AttributeName=idTransaction,AttributeType=S \
  --key-schema \
    AttributeName=idTransaction,KeyType=HASH \
  --billing-mode PAY_PER_REQUEST \
  --endpoint-url http://localhost:8000 \
  --region us-east-1

# 4. Configurar .env para usar local (agregar si es necesario)
echo "AWS_ENDPOINT=http://localhost:8000" >> .env
```

**Opción 3: Usar Makefile (Más Rápido)**

```bash
# Para AWS (usa variables de entorno de .env)
make setup-dynamodb

# Para desarrollo local
make setup-dynamodb-local

# Ver comandos disponibles
make help
```

#### D. Verificar Tabla Creada

```bash
# AWS Cloud
aws dynamodb describe-table \
  --table-name transacciones-blockchain \
  --region us-east-1 \
  --query 'Table.[TableName,TableStatus,ItemCount,TableSizeBytes]' \
  --output table

# Local
aws dynamodb describe-table \
  --table-name transacciones-blockchain \
  --endpoint-url http://localhost:8000 \
  --output table

# Listar todas las tablas
aws dynamodb list-tables --region us-east-1

# Verificar índices
aws dynamodb describe-table \
  --table-name transacciones-blockchain \
  --query 'Table.GlobalSecondaryIndexes[*].[IndexName,IndexStatus]' \
  --output table
```

**Salida esperada:**
```
-------------------------
|    DescribeTable      |
+---------------------+
| transacciones-blockchain |
| ACTIVE              |
| 0                   |
| 0                   |
+---------------------+
```

### 6. Configurar Blockchain (Opcional pero Recomendado)

#### A. Elegir Proveedor RPC (2025)

| Proveedor | Pros | Cons | Recomendado Para |
|-----------|------|------|------------------|
| **Alchemy** | Dashboard excelente, generoso | - | Desarrollo y Producción ⭐ |
| **Infura** | Confiable, establecido | Rate limits estrictos | Producción |
| **QuickNode** | Muy rápido | Más caro | Alta performance |
| **Ankr** | Gratis, descentralizado | Menos features | Desarrollo |

#### B. Crear Cuenta en Proveedor (Ejemplo: Alchemy)

```bash
# 1. Registrarse en https://www.alchemy.com/
# 2. Crear nuevo app
#    - Name: blockchain-medisupply-dev
#    - Chain: Ethereum
#    - Network: Sepolia (testnet)

# 3. Copiar API Key

# 4. Configurar en .env (Recomendado: Alchemy)
echo "ALCHEMY_API_KEY=your_api_key_here" >> .env
echo "BLOCKCHAIN_NETWORK=sepolia" >> .env

# O usar Infura (también soportado):
# echo "BLOCKCHAIN_RPC_URL=https://sepolia.infura.io/v3/your_project_id" >> .env
```

#### C. Crear Wallet Ethereum

```bash
# Método 1: Con cast (foundry) - MÁS SEGURO
curl -L https://foundry.paradigm.xyz | bash
foundryup
cast wallet new

# Método 2: Con geth
geth account new

# Método 3: Programático con Go
cat > /tmp/generate-wallet.go <<'EOF'
package main
import (
    "crypto/ecdsa"
    "fmt"
    "github.com/ethereum/go-ethereum/crypto"
)
func main() {
    privateKey, _ := crypto.GenerateKey()
    privateKeyBytes := crypto.FromECDSA(privateKey)
    publicKey := privateKey.Public().(*ecdsa.PublicKey)
    address := crypto.PubkeyToAddress(*publicKey).Hex()
    
    fmt.Println(" GUARDAR DE FORMA SEGURA Y NUNCA COMPARTIR")
    fmt.Println("Address:", address)
    fmt.Printf("Private Key: %x\n", privateKeyBytes)
}
EOF
go run /tmp/generate-wallet.go
rm /tmp/generate-wallet.go

# ⚠️  IMPORTANTE: Guardar private key en 1Password, LastPass, o similar
# NUNCA en .git o archivos de texto plano
```

#### D. Obtener Fondos en Sepolia Testnet

```bash
# Múltiples faucets (probar varios si uno falla)

# 1. Alchemy Faucet (1 ETH/día) RECOMENDADO
# https://sepoliafaucet.com/

# 2. Chainlink Faucet (20 ETH)
# https://faucets.chain.link/sepolia

# 3. QuickNode Faucet (0.1 ETH)
# https://faucet.quicknode.com/ethereum/sepolia

# 4. Chainlink Faucet (20 ETH)
# https://faucets.chain.link/sepolia

# 5. Verificar balance
# Ir a https://sepolia.etherscan.io/ y buscar tu dirección
```

#### E. Configurar en .env

```bash
# Agregar tu private key (SIN el prefijo 0x)
nano .env

# Agregar estas líneas:
BLOCKCHAIN_NETWORK=sepolia
BLOCKCHAIN_PRIVATE_KEY=tu_private_key_sin_0x
ALCHEMY_API_KEY=tu_api_key  # Recomendado 2025
CONTRACT_ADDRESS=0x0000000000000000000000000000000000000000

# Guardar y cerrar
```

#### F. Verificar Configuración Blockchain

```bash
# Script para verificar conexión y balance
cat > /tmp/check-blockchain.sh <<'EOF'
#!/bin/bash
source .env

echo "Verificando configuración blockchain..."
echo ""
echo "Network: $BLOCKCHAIN_NETWORK"
echo "Provider: ${ALCHEMY_API_KEY:+Alchemy}${BLOCKCHAIN_RPC_URL:+Custom RPC}"
echo ""

# Obtener address de la private key
# (requiere ethers o similar - implementar según sea necesario)
echo "Configuración lista"
EOF

chmod +x /tmp/check-blockchain.sh
/tmp/check-blockchain.sh
rm /tmp/check-blockchain.sh
```

### 7. Iniciar Aplicación

#### Opción A: Docker Compose (Recomendado para Producción)

```bash
# 1. Build imágenes
docker-compose build

# 2. Iniciar todos los servicios en background
docker-compose up -d

# 3. Ver logs en tiempo real
docker-compose logs -f

# 4. Ver solo logs de API
docker-compose logs -f transaccion-blockchain

# 5. Verificar servicios corriendo
docker-compose ps

# 6. Verificar health
curl http://localhost:8082/health
```

**Salida esperada de `docker-compose ps`:**
```
NAME                          STATUS              PORTS
transaccion-blockchain-api    Up 30 seconds       0.0.0.0:8080->8080/tcp
ipfs-node                     Up 30 seconds       0.0.0.0:4001->4001/tcp, 0.0.0.0:5001->5001/tcp, 0.0.0.0:8081->8080/tcp
```

#### Opción B: Desarrollo Local (Hot Reload)

```bash
# Terminal 1: Iniciar IPFS solamente
docker-compose up ipfs

# Terminal 2: Iniciar API con hot reload
# Instalar air primero (si no lo tienes)
go install github.com/air-verse/air@latest

# Iniciar con hot reload
air

# O sin hot reload:
go run cmd/api/main.go

# La API iniciará en http://localhost:8080
# Los cambios en el código recargarán automáticamente con air
```

#### Opción C: Con Makefile (Más Conveniente)

```bash
# Ver todos los comandos disponibles
make help

# Iniciar todo (build + up)
make docker-up

# Ver logs
make docker-logs

# Health check automático
make health-check

# Detener todo
make docker-down

# Rebuild y restart
make docker-restart

# Limpiar todo (incluyendo volúmenes)
make docker-clean
```

#### Opción D: Build desde Código (Sin Docker)

```bash
# Compilar binario
go build -o bin/medisupply-api cmd/api/main.go

# Ejecutar
./bin/medisupply-api

# O en un solo comando
go run cmd/api/main.go
```

#### Verificación Inicial

```bash
# 1. Health check
curl http://localhost:8080/health

# 2. Readiness check
curl http://localhost:8080/ready

# 3. Ver info de la API
curl http://localhost:8080/

# 4. Verificar IPFS
curl -X POST http://localhost:5001/api/v0/version

# 5. Ver logs en tiempo real
docker-compose logs -f transaccion-blockchain
```

**Respuestas esperadas:**
```json
// /health
{"status":"ok","timestamp":"2025-01-15T10:30:00Z","service":"transaccion-blockchain-service"}

// /ready
{"status":"ready","checks":{"dynamodb":"ok","ipfs":"ok"},"timestamp":"2025-01-15T10:30:00Z"}

// IPFS version
{"Version":"0.28.0","Commit":"...","Repo":"..."}
```

## Verificación y Validación

### Checklist de Verificación Completa

```bash
# 1. Verificar versiones
echo "=== Verificando versiones ==="
go version | grep "go1.23"
docker --version | grep "2[4-7]"
docker compose version | grep "2\.[2-9]"

# 2. Verificar servicios Docker
echo "=== Verificando servicios Docker ==="
docker-compose ps | grep "Up"

# 3. Health checks
echo "=== Health checks ==="
curl -f http://localhost:8080/health || echo "Health check falló"
curl -f http://localhost:8080/ready || echo "Ready check falló"

# 4. Verificar IPFS
echo "=== Verificando IPFS ==="
curl -X POST http://localhost:5001/api/v0/version || echo "IPFS no responde"

# 5. Verificar DynamoDB
echo "=== Verificando DynamoDB ==="
aws dynamodb describe-table --table-name transacciones-blockchain --region us-east-1 || echo " DynamoDB no accesible"

# 6. Ejecutar tests
echo "=== Ejecutando tests ==="
go test -short ./tests/... || echo "Tests fallaron"

echo ""
echo "¡Verificación completa!"
```

### Ejecutar Tests

```bash
# Tests unitarios (rápidos, sin dependencias externas)
go test -short -v ./tests/...

# Tests de integración (requiere servicios corriendo)
docker-compose up -d
sleep 5
go test -v ./tests/...

# Tests con coverage (recomendado)
go test -cover -coverprofile=coverage.out ./tests/...
go tool cover -html=coverage.out -o coverage.html
open coverage.html  # Ver reporte HTML

# Tests con race detector (encuentra race conditions)
go test -race ./tests/...

# Benchmark tests
go test -bench=. -benchmem ./tests/...

# Tests en modo verbose con timeout
go test -v -timeout 30s ./tests/...
```

### Prueba End-to-End Completa

```bash
# 1. Ejecutar script de prueba automatizado
chmod +x scripts/test-api.sh
./scripts/test-api.sh

# 2. O prueba manual completa
echo "Iniciando prueba end-to-end..."

# Registrar transacción
RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/transaccion/registrar \
  -H "Content-Type: application/json" \
  -d '{
    "tipoEvento": "fabricacion",
    "idProducto": "TEST-001",
    "datosEvento": "{\"lote\": \"TEST-LOT\", \"cantidad\": 1000}",
    "actorEmisor": "Test Lab"
  }')

echo "Respuesta: $RESPONSE"

# Extraer ID de transacción
TX_ID=$(echo $RESPONSE | jq -r '.idTransaction')
echo "Transaction ID: $TX_ID"

# Consultar transacción
curl -s http://localhost:8080/api/v1/transaccion/$TX_ID | jq

# Verificar integridad
curl -s http://localhost:8080/api/v1/transaccion/verificar/$TX_ID | jq

# Oracle endpoint
curl -s http://localhost:8080/api/v1/oracle/datos/$TX_ID | jq

echo "Prueba end-to-end completada!"
```

### Validación de Seguridad

```bash
# 1. Verificar que .env no está en git
if git ls-files --error-unmatch .env 2>/dev/null; then
    echo " PELIGRO: .env está tracked en git!"
    echo "Ejecuta: git rm --cached .env"
else
    echo ".env no está en git"
fi

# 2. Verificar permisos de archivos
chmod 600 .env
echo "Permisos de .env configurados (solo lectura/escritura para usuario)"

# 3. Escanear vulnerabilidades
govulncheck ./...

# 4. Security lint
gosec ./...

# 5. Verificar secrets no hardcodeados
grep -r "AKIA" --include="*.go" . || echo "No se encontraron AWS keys en código"
grep -r "password.*=.*\"" --include="*.go" . || echo "No se encontraron passwords en código"
```

## 🔧 Configuración por Entorno

### Desarrollo Local

```bash
# .env.development
GIN_MODE=debug
LOG_LEVEL=debug
RATE_LIMIT_REQUESTS=1000
RATE_LIMIT_WINDOW=60
ENABLE_PROFILING=true
USE_AWS_SECRETS=false
```

### Staging

```bash
# .env.staging
GIN_MODE=release
LOG_LEVEL=info
RATE_LIMIT_REQUESTS=500
RATE_LIMIT_WINDOW=60
ENABLE_PROFILING=false
USE_AWS_SECRETS=true
DYNAMODB_TABLE_NAME=transacciones-blockchain-staging
BLOCKCHAIN_NETWORK=sepolia
```

### Producción

```bash
# .env.production
GIN_MODE=release
LOG_LEVEL=warn
RATE_LIMIT_REQUESTS=100
RATE_LIMIT_WINDOW=60
ENABLE_PROFILING=false
USE_AWS_SECRETS=true
DYNAMODB_TABLE_NAME=transacciones-blockchain-prod
BLOCKCHAIN_NETWORK=mainnet  # ⚠️  Usar mainnet con cuidado
ENABLE_METRICS=true
OTEL_EXPORTER_OTLP_ENDPOINT=https://otel-collector:4317
```

## Solución de Problemas

### Problema: "go mod tidy" timeout

**Causa:** Problemas de red o firewall bloqueando proxy.golang.org

**Solución:**
```bash
# Opción 1: Cambiar proxy
export GOPROXY=https://goproxy.io,direct
go mod download

# Opción 2: Usar proxy de Google
export GOPROXY=https://proxy.golang.org,direct
go mod download

# Opción 3: Limpiar caché y reintentar
go clean -modcache
go mod download

# Opción 4: Configurar timeout más largo
export GOPROXY_TIMEOUT=300
go mod download
```

### Problema: Docker no inicia o no responde

**Solución:**
```bash
# 1. Verificar que Docker esté instalado y corriendo
docker --version
docker info

# 2. Reiniciar Docker
# macOS:
killall Docker && open /Applications/Docker.app

# Linux:
sudo systemctl restart docker

# Windows (PowerShell):
Restart-Service docker

# 3. Verificar recursos
docker system df
docker system prune -a  # Limpiar recursos no usados

# 4. Verificar permisos (Linux)
sudo usermod -aG docker $USER
newgrp docker
```

### Problema: IPFS no conecta

**Causa:** Puerto ocupado o contenedor no iniciado

**Solución:**
```bash
# 1. Verificar contenedor
docker-compose ps ipfs

# 2. Ver logs detallados
docker-compose logs ipfs --tail=50

# 3. Verificar puertos
netstat -an | grep 5001  # Linux/Mac
netsh interface ipv4 show excludedportrange protocol=tcp  # Windows

# 4. Reiniciar IPFS
docker-compose restart ipfs

# 5. Recrear contenedor
docker-compose down ipfs
docker-compose up -d ipfs

# 6. Verificar conectividad
curl -X POST http://localhost:5001/api/v0/version

# 7. Verificar volumen
docker volume inspect blockchain-medisupply_ipfs_data
```

### Problema: DynamoDB access denied

**Causa:** Credenciales incorrectas o permisos IAM insuficientes

**Solución:**
```bash
# 1. Verificar credenciales
aws sts get-caller-identity

# 2. Verificar permisos específicos
aws iam get-user-policy --user-name blockchain-medisupply-dev --policy-name DynamoDBPolicy

# 3. Verificar que la tabla existe
aws dynamodb list-tables --region us-east-1

# 4. Test de conexión
aws dynamodb describe-table \
  --table-name transacciones-blockchain \
  --region us-east-1

# 5. Si usa DynamoDB local, verificar endpoint
export AWS_ENDPOINT=http://localhost:8000
# Agregar a .env si es persistente
```

### Problema: "ENCRYPTION_KEY debe tener 32 caracteres"

**Solución:**
```bash
# Generar clave válida
ENCRYPTION_KEY=$(openssl rand -base64 32 | cut -c1-32)
echo "ENCRYPTION_KEY=$ENCRYPTION_KEY" >> .env

# Verificar longitud
grep ENCRYPTION_KEY .env | cut -d'=' -f2 | wc -c
# Debe retornar 33 (32 caracteres + newline)
```

### Problema: "Port already in use"

**Solución:**
```bash
# Encontrar proceso usando el puerto
# macOS/Linux:
lsof -i :8080
kill -9 <PID>

# Windows:
netstat -ano | findstr :8080
taskkill /PID <PID> /F

# Cambiar puerto en .env
echo "SERVER_PORT=8081" >> .env
```

### Problema: "Cannot connect to blockchain"

**Solución:**
```bash
# 1. Verificar Alchemy API Key
curl -X POST https://eth-sepolia.g.alchemy.com/v2/$ALCHEMY_API_KEY \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'

# 2. Verificar balance de testnet
# Ir a https://sepolia.etherscan.io/address/YOUR_ADDRESS

# 3. Verificar private key (sin 0x prefix)
grep BLOCKCHAIN_PRIVATE_KEY .env

# 4. Test de conexión programático
cat > /tmp/test-rpc.go <<'EOF'
package main
import (
    "context"
    "fmt"
    "github.com/ethereum/go-ethereum/ethclient"
)
func main() {
    client, _ := ethclient.Dial("https://sepolia.infura.io/v3/YOUR_PROJECT_ID")
    blockNumber, _ := client.BlockNumber(context.Background())
    fmt.Printf("Latest block: %d\n", blockNumber)
}
EOF
go run /tmp/test-rpc.go
```

### Problema: Tests fallan

**Solución:**
```bash
# 1. Ejecutar tests con verbose para ver detalles
go test -v ./tests/...

# 2. Ejecutar test específico
go test -v -run TestNombreEspecifico ./tests/...

# 3. Limpiar caché de tests
go clean -testcache

# 4. Verificar dependencias de tests
docker-compose up -d ipfs
sleep 5
go test ./tests/...

# 5. Ver logs de test
go test -v ./tests/... 2>&1 | tee test.log
```

### Problema: Memoria o CPU alta

**Solución:**
```bash
# 1. Ver uso de recursos Docker
docker stats

# 2. Limitar recursos en docker-compose.yml
# deploy:
#   resources:
#     limits:
#       cpus: '1.0'
#       memory: 512M

# 3. Ver goroutines (si API está corriendo)
curl http://localhost:8080/debug/pprof/goroutine

# 4. Profiling
go test -cpuprofile=cpu.prof -memprofile=mem.prof ./tests/...
go tool pprof cpu.prof
```

## Instalación Avanzada

### Usando AWS Secrets Manager (Producción)

```bash
# 1. Crear secrets en AWS
aws secretsmanager create-secret \
  --name prod/blockchain-medisupply/encryption-key \
  --secret-string "$(openssl rand -base64 32 | cut -c1-32)"

aws secretsmanager create-secret \
  --name prod/blockchain-medisupply/blockchain-private-key \
  --secret-string "your_private_key_here"

# 2. Configurar en .env
USE_AWS_SECRETS=true
ENCRYPTION_KEY_SECRET=prod/blockchain-medisupply/encryption-key
BLOCKCHAIN_PRIVATE_KEY_SECRET=prod/blockchain-medisupply/blockchain-private-key

# 3. Dar permisos al usuario IAM
aws iam attach-user-policy \
  --user-name blockchain-medisupply-prod \
  --policy-arn arn:aws:iam::aws:policy/SecretsManagerReadWrite
```

### Deploy en Kubernetes

```bash
# 1. Crear namespace
kubectl create namespace blockchain-medisupply

# 2. Crear secrets
kubectl create secret generic app-secrets \
  --from-literal=encryption-key=$(openssl rand -base64 32 | cut -c1-32) \
  --from-literal=aws-access-key=YOUR_ACCESS_KEY \
  --from-literal=aws-secret-key=YOUR_SECRET_KEY \
  -n blockchain-medisupply

# 3. Aplicar manifests (ver README para ejemplos)
kubectl apply -f k8s/ -n blockchain-medisupply

# 4. Verificar deployment
kubectl get pods -n blockchain-medisupply
kubectl logs -f deployment/transaccion-blockchain -n blockchain-medisupply
```

### Multi-región Setup

```bash
# Configurar réplicas en múltiples regiones
# .env.us-east-1
AWS_REGION=us-east-1
DYNAMODB_TABLE_NAME=transacciones-blockchain-global

# .env.eu-west-1
AWS_REGION=eu-west-1
DYNAMODB_TABLE_NAME=transacciones-blockchain-global

# Habilitar DynamoDB Global Tables
aws dynamodb create-global-table \
  --global-table-name transacciones-blockchain-global \
  --replication-group RegionName=us-east-1 RegionName=eu-west-1
```

## Actualización de Versiones

### Actualizar Dependencias Go

```bash
# Ver dependencias desactualizadas
go list -u -m all

# Actualizar todas
go get -u ./...
go mod tidy

# Actualizar específica
go get -u github.com/gin-gonic/gin@latest
go mod tidy

# Verificar vulnerabilidades
govulncheck ./...
```

### Actualizar Imágenes Docker

```bash
# Pull últimas imágenes
docker-compose pull

# Rebuild con --no-cache
docker-compose build --no-cache

# Restart servicios
docker-compose down
docker-compose up -d
```

## Desinstalación

### Desinstalación Completa

```bash
# 1. Detener todos los servicios
docker-compose down

# 2. Eliminar volúmenes (⚠️  BORRA TODOS LOS DATOS)
docker-compose down -v

# 3. Eliminar imágenes
docker rmi $(docker images 'blockchain-medisupply*' -q)

# 4. Eliminar tabla DynamoDB (⚠️  PERMANENTE)
aws dynamodb delete-table \
  --table-name transacciones-blockchain \
  --region us-east-1

# 5. Eliminar archivos locales
cd ..
rm -rf blockchain-medisupply

# 6. Limpiar AWS resources
aws iam detach-user-policy \
  --user-name blockchain-medisupply-dev \
  --policy-arn arn:aws:iam::YOUR_ACCOUNT:policy/BlockchainMediSupplyDynamoDBPolicy
aws iam delete-user --user-name blockchain-medisupply-dev
```

### Desinstalación Parcial (Mantener Datos)

```bash
# Solo detener servicios
docker-compose stop

# Los volúmenes y datos se mantienen
# Reiniciar con: docker-compose start
```

## Siguientes Pasos

Una vez completada la instalación:

1. **Leer documentación completa**: [README.md](./README.md)
2. **Quick start guide**: Ver sección "Demo Rápido" en README
3. **Configuración avanzada**: [CONFIG.md](./CONFIG.md)
4. **Security best practices**: Ver sección de seguridad en README
5. **Troubleshooting adicional**: Ver sección en README.md

## Soporte

Si encuentras problemas durante la instalación:

### Opciones de Ayuda

1. **Documentación**: Revisa [README.md](./README.md) y [CONFIG.md](./CONFIG.md)
2. **Logs**: Siempre revisa los logs primero
   ```bash
   docker-compose logs -f
   ```
3. **GitHub Issues**: [Reportar bug](https://github.com/edinfamous/blockchain-medisupply/issues)
4. **Discussions**: [Hacer preguntas](https://github.com/edinfamous/blockchain-medisupply/discussions)

### Información Útil para Reportar Issues

```bash
# Recopilar información del sistema
cat > /tmp/system-info.txt <<EOF
OS: $(uname -a)
Go version: $(go version)
Docker version: $(docker --version)
Docker Compose: $(docker compose version)
AWS CLI: $(aws --version)

Services status:
$(docker-compose ps)

Recent logs:
$(docker-compose logs --tail=50)
EOF

cat /tmp/system-info.txt
```

---

**¡Felicitaciones! 🎉** Has completado la instalación de blockchain-medisupply.

Para comenzar a usar el sistema, continúa con la [documentación principal](./README.md).

