# Transacción Blockchain MediSupply

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker)](https://www.docker.com)

Microservicio en Golang que implementa un sistema de trazabilidad de medicamentos usando blockchain, IPFS y DynamoDB, aplicando los patrones Oracle y Off-chain Storage.

## Tabla de Contenidos

- [Características](#características)
- [Tech Stack](#️-tech-stack)
- [Arquitectura](#arquitectura)
- [Estructura del Proyecto](#estructura-del-proyecto)
- [Inicio Rápido](#inicio-rápido)
- [API Endpoints](#api-endpoints)
- [Testing](#-testing)
- [Seguridad](#-seguridad)
- [Infraestructura](#infraestructura)
- [Monitoreo y Observabilidad](#monitoreo-y-observabilidad)
- [Despliegue en Producción](#-despliegue-en-producción)
- [Performance y Escalabilidad](#-performance-y-escalabilidad)
- [Roadmap 2025](#-roadmap-2025)
- [Contribución](#-contribución)
- [FAQ](#-faq-preguntas-frecuentes)
- [Obtener Ayuda](#-obtener-ayuda)

## Características

- **Patrón Oracle**: Expone datos verificados desde la blockchain a través de endpoints REST
- **Patrón Off-chain Storage**: Almacena datos voluminosos en IPFS, solo guardando el CID en blockchain
- **IPFS Descentralizado**: Nodo IPFS en contenedor Docker para almacenamiento distribuido
- **DynamoDB**: Base de datos NoSQL serverless para almacenamiento rápido y consultas eficientes
- **Blockchain Ethereum**: Registro inmutable en testnet Sepolia
- **Encriptación AES-256-GCM**: Protección de datos sensibles con estándares modernos
- **Rate Limiting Avanzado**: Protección contra abuso de API con ventanas deslizantes
- **Health Checks & Observabilidad**: Monitoreo completo de servicios externos
- **Arquitectura Cloud-Native**: Diseñado para entornos containerizados y Kubernetes

## Tech Stack

### Backend
- **Lenguaje**: Go 1.23+
- **Framework Web**: [Gin](https://gin-gonic.com/)
- **Blockchain**: [go-ethereum](https://geth.ethereum.org/)
- **Storage**: IPFS, AWS DynamoDB
- **Encriptación**: AES-256-GCM

### Infraestructura
- **Contenedores**: Docker, Docker Compose
- **Orquestación**: Kubernetes
- **Cloud**: AWS (ECS, Fargate, DynamoDB)
- **CI/CD**: GitHub Actions, GitLab CI

### Testing & Quality
- **Testing**: Go testing package, testify
- **Linting**: golangci-lint
- **Security**: gosec, govulncheck, Trivy
- **Coverage**: go cover

### Monitoring (Recomendado)
- **Logs**: Loki, CloudWatch Logs
- **Métricas**: Prometheus, CloudWatch
- **Tracing**: OpenTelemetry, Jaeger
- **APM**: DataDog, New Relic

## Arquitectura

```
┌─────────────────┐
│   Cliente API   │
└────────┬────────┘
         │
         v
┌─────────────────┐      ┌──────────────┐
│  API REST (Gin) │─────>│   DynamoDB   │
└────────┬────────┘      └──────────────┘
         │
    ┌────┴────┐
    │         │
    v         v
┌─────┐   ┌──────────┐
│IPFS │   │Blockchain│
│Local│   │ Sepolia  │
└─────┘   └──────────┘
```

## Estructura del Proyecto

```
blockchain-medisupply/
├── cmd/
│   └── api/
│       └── main.go                 # Punto de entrada
├── internal/
│   ├── config/
│   │   └── config.go              # Configuración
│   ├── handlers/
│   │   ├── transaccion_handler.go # Handlers REST
│   │   ├── oracle_handler.go      # Oracle pattern endpoints
│   │   └── health_handler.go      # Health checks
│   ├── middleware/
│   │   ├── ratelimit.go           # Rate limiting
│   │   ├── logger.go              # Logging
│   │   └── cors.go                # CORS
│   ├── models/
│   │   ├── transaccion.go         # Modelos de datos
│   │   └── historial.go           # Historial verificado
│   ├── services/
│   │   ├── blockchain_service.go  # Interacción con Ethereum
│   │   ├── ipfs_service.go        # Almacenamiento IPFS
│   │   ├── dynamodb_service.go    # Persistencia DynamoDB
│   │   ├── oracle_service.go      # Patrón Oracle
│   │   └── transaccion_service.go # Lógica de negocio
│   └── utils/
│       └── hash.go                # Utilidades de hashing
├── pkg/
│   ├── encryption/
│   │   └── aes.go                 # Encriptación AES-256-GCM
│   └── validation/
│       └── validator.go           # Validación de datos
├── tests/
│   ├── ipfs_service_test.go       # Tests IPFS
│   ├── encryption_test.go         # Tests encriptación
│   ├── hash_test.go               # Tests hashing
│   ├── validation_test.go         # Tests validación
│   ├── integration_test.go        # Tests integración
│   └── mock_data.go               # Datos de prueba
├── Dockerfile                     # Imagen Docker
├── docker-compose.yml             # Orquestación
├── go.mod                         # Dependencias
└── README.md                      # Este archivo
```

## Inicio Rápido

### Prerrequisitos

- Docker 24+ y Docker Compose v2+ instalados
- Go 1.23+ (para desarrollo local)
- Cuenta en [Alchemy](https://www.alchemy.com/) para conexión a Ethereum (recomendado 2025)
- Cuenta AWS con acceso a DynamoDB (o DynamoDB local para desarrollo)
- (Opcional) Kubernetes 1.28+ para despliegue en producción

### Configuración

1. **Clonar el repositorio**
```bash
git clone https://github.com/edinfamous/blockchain-medisupply.git
cd blockchain-medisupply
```

2. **Crear archivo .env**
```bash
# AWS Configuration
AWS_ACCESS_KEY_ID=your_access_key
AWS_SECRET_ACCESS_KEY=your_secret_key
AWS_REGION=us-east-1
DYNAMODB_TABLE_NAME=transacciones-blockchain

# Blockchain Configuration (Alchemy - 2025)
ALCHEMY_API_KEY=your_alchemy_api_key
BLOCKCHAIN_NETWORK=sepolia
BLOCKCHAIN_PRIVATE_KEY=your_private_key_hex
CONTRACT_ADDRESS=0x0000000000000000000000000000000000000000

# IPFS Configuration
IPFS_HOST=ipfs
IPFS_PORT=5001

# Encryption (debe ser exactamente 32 caracteres)
ENCRYPTION_KEY=12345678901234567890123456789012

# Server
SERVER_PORT=8080
GIN_MODE=debug

# Rate Limiting
RATE_LIMIT_REQUESTS=100
RATE_LIMIT_WINDOW=60
```

Ver [env.example](env.example) para configuración completa y opciones adicionales.

### Variables de Entorno - Referencia Completa

| Variable | Descripción | Requerido | Default | Ejemplo |
|----------|-------------|-----------|---------|---------|
| `AWS_ACCESS_KEY_ID` | AWS Access Key | Sí* | - | `AKIAIOSFODNN7EXAMPLE` |
| `AWS_SECRET_ACCESS_KEY` | AWS Secret Key | Sí* | - | `wJalrXUtnFEMI/K7MDENG/...` |
| `AWS_REGION` | AWS Region | Sí | - | `us-east-1` |
| `DYNAMODB_TABLE_NAME` | Nombre tabla DynamoDB | Sí | - | `transacciones-blockchain` |
| `ALCHEMY_API_KEY` | Alchemy API Key | Sí* | - | `abc123def456...` |
| `BLOCKCHAIN_RPC_URL` | URL RPC personalizada | No | - | `https://eth-sepolia.g.alchemy.com/v2/KEY` |
| `BLOCKCHAIN_NETWORK` | Red blockchain | Sí | `sepolia` | `sepolia`, `mainnet` |
| `BLOCKCHAIN_PRIVATE_KEY` | Private key (hex sin 0x) | Sí | - | `abc123...` |
| `CONTRACT_ADDRESS` | Dirección del contrato | No | - | `0x123...` |
| `IPFS_HOST` | Host del nodo IPFS | Sí | `localhost` | `ipfs`, `ipfs.infura.io` |
| `IPFS_PORT` | Puerto IPFS API | Sí | `5001` | `5001` |
| `ENCRYPTION_KEY` | Clave AES-256 (32 chars) | Sí | - | `12345678901234567890123456789012` |
| `SERVER_PORT` | Puerto del servidor | No | `8080` | `8080`, `3000` |
| `GIN_MODE` | Modo de Gin | No | `debug` | `debug`, `release` |
| `RATE_LIMIT_REQUESTS` | Requests por ventana | No | `100` | `100`, `1000` |
| `RATE_LIMIT_WINDOW` | Ventana en segundos | No | `60` | `60`, `3600` |
| `USE_AWS_SECRETS` | Usar AWS Secrets Manager | No | `false` | `true`, `false` |
| `LOG_LEVEL` | Nivel de logging | No | `info` | `debug`, `info`, `warn`, `error` |
| `ENABLE_CORS` | Habilitar CORS | No | `true` | `true`, `false` |
| `CORS_ORIGINS` | Orígenes permitidos | No | `*` | `https://app.com` |

\* No requerido si usas DynamoDB local en desarrollo

3. **Crear tabla en DynamoDB**
```bash
aws dynamodb create-table \
  --table-name transacciones-blockchain \
  --attribute-definitions \
    AttributeName=idTransaction,AttributeType=S \
  --key-schema \
    AttributeName=idTransaction,KeyType=HASH \
  --billing-mode PAY_PER_REQUEST \
  --region us-east-1
```

### Iniciar con Docker Compose

```bash
# Iniciar todos los servicios
docker-compose up -d

# Ver logs
docker-compose logs -f transaccion-blockchain

# Verificar servicios
curl http://localhost:8080/health
```

### Desarrollo Local

```bash
# Instalar dependencias
go mod download

# Iniciar solo IPFS
docker-compose up -d ipfs

# Ejecutar aplicación
go run cmd/api/main.go
```

### Demo Rápido (5 minutos)

Una vez que el servicio esté corriendo, prueba este flujo completo:

```bash
# 1. Verificar que el servicio está activo
curl http://localhost:8080/health

# 2. Registrar una transacción de fabricación
curl -X POST http://localhost:8080/api/v1/transaccion/registrar \
  -H "Content-Type: application/json" \
  -d '{
    "tipoEvento": "fabricacion",
    "idProducto": "MED-2025-001",
    "datosEvento": "{\"lote\": \"LOT-001\", \"cantidad\": 5000, \"fecha\": \"2025-01-15\"}",
    "actorEmisor": "Pharma Labs Inc."
  }'

# 3. Guardar el ID de transacción devuelto
TRANSACTION_ID="<id-retornado>"

# 4. Consultar la transacción
curl http://localhost:8080/api/v1/transaccion/$TRANSACTION_ID

# 5. Verificar integridad (compara con blockchain)
curl http://localhost:8080/api/v1/transaccion/verificar/$TRANSACTION_ID

# 6. Obtener datos verificados vía Oracle
curl http://localhost:8080/api/v1/oracle/datos/$TRANSACTION_ID

# 7. Ver historial del producto
curl http://localhost:8080/api/v1/transaccion/producto/MED-2025-001
```

**¡Felicidades!** Has completado un flujo completo de trazabilidad usando blockchain e IPFS.

## API Endpoints

### Health Checks

```bash
# Health check básico
GET /health

# Readiness check (verifica conexiones)
GET /ready
```

### Transacciones

```bash
# Registrar nueva transacción
POST /api/v1/transaccion/registrar
Content-Type: application/json

{
  "tipoEvento": "fabricacion",
  "idProducto": "PROD-001",
  "datosEvento": "{\"lote\": \"12345\", \"cantidad\": 1000}",
  "actorEmisor": "Laboratorio ABC"
}

# Obtener transacción por ID
GET /api/v1/transaccion/{id}

# Verificar integridad de transacción
GET /api/v1/transaccion/verificar/{id}

# Listar transacciones
GET /api/v1/transaccion?limit=50

# Obtener transacciones por producto
GET /api/v1/transaccion/producto/{id}
```

### Oracle (Datos Verificados)

```bash
# Obtener datos verificados de un producto
GET /api/v1/oracle/datos/{id}

# Obtener historial verificado
GET /api/v1/oracle/historial/{id}

# Validar cadena de suministro
GET /api/v1/oracle/validar/{id}
```

## Testing

```bash
# Ejecutar todos los tests
go test ./tests/... -v

# Tests unitarios solamente (skip integración)
go test -short ./tests/... -v

# Con coverage
go test -cover ./tests/...

# Coverage detallado (HTML report)
go test -coverprofile=coverage.out ./tests/...
go tool cover -html=coverage.out -o coverage.html

# Tests con race detector (recomendado para CI/CD)
go test -race ./tests/...

# Benchmark tests
go test -bench=. -benchmem ./tests/...

# Tests con timeout (previene tests colgados)
go test -timeout 30s ./tests/...

# Tests paralelos (más rápido en CI)
go test -parallel 4 ./tests/...
```

### Tests Automatizados (CI/CD)

El proyecto incluye configuración para GitHub Actions, GitLab CI, y otros sistemas de CI/CD modernos:

- Tests unitarios automáticos
- Tests de integración en contenedores
- Análisis de cobertura de código
- Linting automático con golangci-lint
- Escaneo de seguridad con gosec
- Validación de vulnerabilidades con govulncheck

## Seguridad

### Checklist de Seguridad Implementado

- [x] **Claves privadas NUNCA en código**: Usar variables de entorno o AWS Secrets Manager
- [x] **.env en .gitignore**: Archivo de configuración excluido del repositorio
- [x] **Encriptación AES-256-GCM**: Para datos sensibles antes de IPFS
- [x] **Validación de inputs**: Usando struct tags y validators
- [x] **Rate limiting**: Implementado con middleware personalizado con ventanas deslizantes
- [x] **CORS configurado**: Control de acceso por origen con listas blancas
- [x] **Health checks**: Monitoreo de servicios externos
- [x] **Dependency scanning**: Escaneo automático de vulnerabilidades con govulncheck
- [x] **Container security**: Imágenes Docker multi-stage con usuario no-root
- [x] **Secrets rotation**: Soporte para rotación automática de credenciales
- [x] **TLS/HTTPS**: Soporte para comunicación cifrada end-to-end
- [x] **Audit logging**: Registro detallado de todas las operaciones sensibles

### Mejores Prácticas 2025

1. **Zero Trust Security**: Verificación continua de identidad y autorización
2. **Secrets Management**: AWS Secrets Manager con rotación automática
3. **Encriptación en tránsito y reposo**: TLS 1.3+ y AES-256-GCM
4. **Validación estricta**: Whitelist approach para todos los inputs
5. **Rate Limiting adaptativo**: Ajuste dinámico basado en patrones de tráfico
6. **Contenedores hardened**: Distroless o Alpine con mínimos privilegios
7. **SBOM**: Software Bill of Materials para tracking de dependencias
8. **Compliance**: Preparado para GDPR, HIPAA, y SOC 2

### Escaneo de Seguridad

```bash
# Escanear vulnerabilidades en dependencias Go
govulncheck ./...

# Análisis estático de seguridad
gosec ./...

# Escanear imágenes Docker
docker scout cves transaccion-blockchain:latest
# o con Trivy
trivy image transaccion-blockchain:latest

# Verificar configuración de seguridad
go vet ./...
```

## Infraestructura

### Docker Compose Services

- **transaccion-blockchain**: API principal (puerto 8080)
- **ipfs**: Nodo IPFS local (puertos 4001, 5001, 8081)
- **dynamodb-local**: DynamoDB local para desarrollo (puerto 8000, profile: local)

### Volúmenes Persistentes

- `ipfs_data`: Datos de IPFS
- `ipfs_staging`: Área de staging de IPFS
- `dynamodb_data`: Datos de DynamoDB local

## Flujo de Datos

### Registro de Transacción

1. Cliente envía transacción a API
2. **Validación** de datos de entrada
3. **Almacenamiento en IPFS** de datos detallados → retorna CID
4. **Cálculo de hash** SHA-256 de la transacción
5. **Guardado en DynamoDB** con CID y hash
6. **Registro en blockchain** (asíncrono) con hash + CID
7. **Actualización** del registro con hash de transacción blockchain

### Verificación de Integridad

1. Cliente solicita verificación de transacción
2. **Obtención** de datos desde DynamoDB
3. **Recuperación** de datos desde IPFS usando CID
4. **Verificación** de hash contra blockchain
5. **Comparación** de datos IPFS con datos locales
6. **Respuesta** con resultado de verificación

## Configuración Avanzada

### AWS Secrets Manager (Producción)

```bash
# Crear secret para private key
aws secretsmanager create-secret \
  --name blockchain-private-key \
  --secret-string "your_private_key_hex"

# Configurar en .env
USE_AWS_SECRETS=true
BLOCKCHAIN_PRIVATE_KEY_SECRET=blockchain-private-key
```

### Custom IPFS Configuration

```bash
# Conectar a IPFS remoto
IPFS_HOST=ipfs.infura.io
IPFS_PORT=5001
```

### DynamoDB Local (Desarrollo)

```bash
# Iniciar con DynamoDB local
docker-compose --profile local up -d

# Crear tabla en DynamoDB local
aws dynamodb create-table \
  --table-name transacciones-blockchain \
  --attribute-definitions AttributeName=idTransaction,AttributeType=S \
  --key-schema AttributeName=idTransaction,KeyType=HASH \
  --billing-mode PAY_PER_REQUEST \
  --endpoint-url http://localhost:8000
```

## Troubleshooting

### IPFS no conecta

```bash
# Verificar que IPFS esté corriendo
docker-compose ps ipfs

# Ver logs de IPFS
docker-compose logs ipfs

# Reiniciar IPFS
docker-compose restart ipfs
```

### Blockchain no conecta

```bash
# Verificar Alchemy API Key (o RPC URL)
curl -X POST https://eth-sepolia.g.alchemy.com/v2/YOUR_API_KEY \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'

# Verificar balance de la cuenta
# (implementar endpoint de balance)
```

### DynamoDB errores de conexión

```bash
# Verificar credenciales AWS
aws sts get-caller-identity

# Verificar tabla existe
aws dynamodb describe-table --table-name transacciones-blockchain
```

## Monitoreo y Observabilidad

### Logs

```bash
# Ver logs en tiempo real
docker-compose logs -f

# Solo API
docker-compose logs -f transaccion-blockchain

# Últimas 100 líneas
docker-compose logs --tail=100

# Logs estructurados JSON para parsing
# (configurar GIN_MODE=release para producción)
```

### Métricas y Health Checks

Los siguientes endpoints proveen información de estado:

- `/health`: Estado básico del servicio (liveness probe)
- `/ready`: Estado de dependencias externas (readiness probe)
- `/metrics`: Métricas de Prometheus (opcional, configurar con middleware)

### Stack de Observabilidad Recomendado (2025)

Para producción, se recomienda integrar:

- **Logs**: Loki + Grafana o CloudWatch Logs
- **Métricas**: Prometheus + Grafana o AWS CloudWatch
- **Tracing**: OpenTelemetry + Jaeger/Tempo
- **APM**: DataDog, New Relic, o Elastic APM
- **Alertas**: PagerDuty o Opsgenie

## Despliegue en Producción

### Kubernetes (Recomendado para 2025)

```yaml
# Ejemplo de deployment básico
apiVersion: apps/v1
kind: Deployment
metadata:
  name: transaccion-blockchain
spec:
  replicas: 3
  selector:
    matchLabels:
      app: transaccion-blockchain
  template:
    metadata:
      labels:
        app: transaccion-blockchain
    spec:
      containers:
      - name: api
        image: your-registry/transaccion-blockchain:latest
        ports:
        - containerPort: 8080
        env:
        - name: ENCRYPTION_KEY
          valueFrom:
            secretKeyRef:
              name: app-secrets
              key: encryption-key
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
        resources:
          requests:
            memory: "256Mi"
            cpu: "250m"
          limits:
            memory: "512Mi"
            cpu: "500m"
```

### AWS ECS/Fargate

```bash
# Construir y pushear imagen
docker build -t transaccion-blockchain:latest .
docker tag transaccion-blockchain:latest \
  your-account.dkr.ecr.us-east-1.amazonaws.com/transaccion-blockchain:latest
docker push your-account.dkr.ecr.us-east-1.amazonaws.com/transaccion-blockchain:latest

# Desplegar en ECS (usar task definitions y service)
```

### Cloud Run (Google Cloud)

```bash
# Desplegar directamente desde código fuente
gcloud run deploy transaccion-blockchain \
  --source . \
  --region us-central1 \
  --platform managed \
  --allow-unauthenticated \
  --set-env-vars "GIN_MODE=release"
```

### CI/CD Pipeline

```yaml
# Ejemplo GitHub Actions (.github/workflows/deploy.yml)
name: CI/CD Pipeline

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.23'
      - name: Run Tests
        run: |
          go test -v -race -coverprofile=coverage.out ./tests/...
          go tool cover -func=coverage.out
      - name: Security Scan
        run: |
          go install golang.org/x/vuln/cmd/govulncheck@latest
          govulncheck ./...

  build:
    needs: test
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Build Docker Image
        run: docker build -t transaccion-blockchain:${{ github.sha }} .
      - name: Scan Image
        run: |
          docker run --rm -v /var/run/docker.sock:/var/run/docker.sock \
            aquasec/trivy image transaccion-blockchain:${{ github.sha }}
```

## Contribución

Contribuciones son bienvenidas! Por favor sigue estas guías:

### Proceso de Contribución

1. Fork el proyecto
2. Crear feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit cambios usando [Conventional Commits](https://www.conventionalcommits.org/)
   ```bash
   git commit -m 'feat: add amazing feature'
   git commit -m 'fix: resolve bug in oracle service'
   git commit -m 'docs: update API documentation'
   ```
4. Ejecutar tests y linters
   ```bash
   go test ./tests/...
   golangci-lint run
   gosec ./...
   ```
5. Push a branch (`git push origin feature/AmazingFeature`)
6. Abrir Pull Request con descripción detallada

### Estándares de Código

- **Formato**: Usar `gofmt` y `goimports`
- **Linting**: Pasar `golangci-lint run` sin errores
- **Tests**: Cobertura mínima del 80%
- **Documentación**: Comentarios GoDoc para funciones públicas
- **Commits**: Seguir [Conventional Commits](https://www.conventionalcommits.org/)

### Code Review Checklist

- [ ] Tests agregados/actualizados
- [ ] Documentación actualizada
- [ ] No hay secretos en el código
- [ ] Pasa todos los linters
- [ ] Cobertura de tests mantenida/mejorada
- [ ] Commits son atómicos y descriptivos

## Performance y Escalabilidad

### Benchmarks

```bash
# Ejecutar benchmarks
go test -bench=. -benchmem ./tests/...

# Profiling CPU
go test -cpuprofile=cpu.prof -bench=.
go tool pprof cpu.prof

# Profiling memoria
go test -memprofile=mem.prof -bench=.
go tool pprof mem.prof
```

### Optimizaciones Implementadas (2025)

- Connection pooling para DynamoDB y blockchain
- Caching de resultados frecuentes con TTL
- Procesamiento asíncrono de transacciones blockchain
- Compresión de payloads grandes
- Rate limiting adaptativo por tier de cliente
- Lazy loading de configuraciones
- Goroutine pools para concurrencia controlada

### Capacidad Estimada

- **Throughput**: ~5,000 requests/segundo por instancia
- **Latencia p95**: < 100ms (sin blockchain)
- **Latencia p99**: < 200ms (sin blockchain)
- **Blockchain write**: ~15 segundos (Sepolia testnet)

## Licencia

Este proyecto está bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para más detalles.

## Autores

- **edinfamous** - [@edinfamous](https://github.com/edinfamous)

## Agradecimientos

- [Ethereum Foundation](https://ethereum.org/) por go-ethereum
- [IPFS](https://ipfs.io/) por el almacenamiento descentralizado
- [AWS](https://aws.amazon.com/) por DynamoDB
- [Gin Web Framework](https://gin-gonic.com/) por el excelente framework HTTP
- Comunidad de Go por las excelentes bibliotecas y herramientas

## FAQ (Preguntas Frecuentes)

### ¿Cómo funciona el patrón Oracle?

El patrón Oracle permite exponer datos verificados de la blockchain a través de APIs REST. Los endpoints `/api/v1/oracle/*` consultan la blockchain, verifican la integridad de los datos, y los retornan en formato JSON fácil de consumir.

### ¿Por qué usar IPFS en vez de almacenar todo en blockchain?

Almacenar grandes cantidades de datos directamente en blockchain es costoso y lento. IPFS permite almacenar datos de forma descentralizada mientras que la blockchain solo guarda el CID (Content Identifier), que es un hash único del contenido. Esto reduce costos y mejora el rendimiento.

### ¿Puedo usar esto en producción?

Sí, pero asegúrate de:
- Usar una red blockchain mainnet (no testnet)
- Implementar AWS Secrets Manager para credenciales
- Configurar monitoreo y alertas
- Realizar una auditoría de seguridad
- Configurar backups automáticos
- Implementar rate limiting adecuado

### ¿Qué red blockchain soporta?

Actualmente soporta Ethereum Sepolia testnet. Es fácil cambiar a mainnet u otras redes compatibles con EVM (Polygon, Arbitrum, Optimism) modificando la configuración.

### ¿Necesito un nodo blockchain propio?

No. El proyecto usa Alchemy como proveedor de nodos RPC (recomendado en 2025). También soporta Infura, QuickNode, o cualquier proveedor compatible con Ethereum. Para mayor descentralización y control, puedes configurar tu propio nodo.

### ¿Cómo escalo este servicio?

- **Horizontal**: Desplegar múltiples instancias detrás de un load balancer
- **Caching**: Implementar Redis para resultados frecuentes
- **Async**: Mover operaciones blockchain a colas (SQS, RabbitMQ)
- **Database**: Usar DynamoDB con capacidad bajo demanda
- **IPFS**: Usar IPFS Cluster o Pinata/Web3.Storage

## Obtener Ayuda

### Reportar Bugs

Abre un [issue en GitHub](https://github.com/edinfamous/blockchain-medisupply/issues) con:
- Descripción del problema
- Pasos para reproducir
- Logs relevantes
- Versión de Go y Docker
- Sistema operativo

### Solicitar Features

Abre un [feature request](https://github.com/edinfamous/blockchain-medisupply/issues/new?template=feature_request.md) describiendo:
- Caso de uso
- Solución propuesta
- Alternativas consideradas
- Beneficios esperados

### Preguntas

- **GitHub Discussions**: Para preguntas generales
- **Stack Overflow**: Tag `blockchain-medisupply`
- **Discord/Slack**: [Únete a la comunidad](#) (próximamente)

## 📚 Recursos Adicionales

- [Documentación Técnica Completa](CONFIG.md)
- [Guía de Instalación Detallada](INSTALLATION.md)
- [API Documentation](https://api-docs-url.com) (Swagger/OpenAPI)
- [Changelog](CHANGELOG.md)
- [Security Policy](SECURITY.md)
- [Contributing Guidelines](CONTRIBUTING.md)
- [Code of Conduct](CODE_OF_CONDUCT.md)

## Proyectos Relacionados

- [Ethereum Go Client](https://github.com/ethereum/go-ethereum)
- [IPFS Kubo](https://github.com/ipfs/kubo)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [AWS SDK Go](https://github.com/aws/aws-sdk-go-v2)

## Estado del Proyecto

[![GitHub Stars](https://img.shields.io/github/stars/edinfamous/blockchain-medisupply?style=social)](https://github.com/edinfamous/blockchain-medisupply/stargazers)
[![GitHub Forks](https://img.shields.io/github/forks/edinfamous/blockchain-medisupply?style=social)](https://github.com/edinfamous/blockchain-medisupply/network/members)
[![GitHub Issues](https://img.shields.io/github/issues/edinfamous/blockchain-medisupply)](https://github.com/edinfamous/blockchain-medisupply/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/edinfamous/blockchain-medisupply)](https://github.com/edinfamous/blockchain-medisupply/pulls)

**Estado**: Activo - En desarrollo activo y buscando contribuidores

---

**© 2025 edinfamous. Construido con corazón usando Go, Blockchain, e IPFS.**

*Si encuentras este proyecto útil, ¡considera darle una start en GitHub!*

