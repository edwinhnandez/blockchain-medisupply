# Guía de Configuración - Variables de Entorno

## Archivo de Configuración

Copia el archivo `env.example` a `.env` en la raíz del proyecto:

```bash
cp env.example .env
```

Luego edita `.env` con tus valores reales:

```bash
nano .env
# o
vim .env
# o
code .env
```

## Variables de Entorno Requeridas

### AWS Configuration (OBLIGATORIO)

```bash
# Access Key ID de tu usuario IAM
AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE

# Secret Access Key de tu usuario IAM
AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY

# Región de AWS donde está tu DynamoDB
AWS_REGION=us-east-1

# Nombre de la tabla DynamoDB
DYNAMODB_TABLE_NAME=transacciones-blockchain
```

**Cómo obtener:**
1. Ir a [AWS Console](https://console.aws.amazon.com)
2. IAM > Users > Tu usuario > Security Credentials
3. Create Access Key
4. Copiar Access Key ID y Secret Access Key

**Permisos IAM necesarios:**
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "dynamodb:PutItem",
        "dynamodb:GetItem",
        "dynamodb:UpdateItem",
        "dynamodb:Scan",
        "dynamodb:Query"
      ],
      "Resource": "arn:aws:dynamodb:us-east-1:*:table/transacciones-blockchain"
    }
  ]
}
```

### Encryption (OBLIGATORIO)

```bash
# Clave de 32 caracteres para AES-256-GCM
ENCRYPTION_KEY=12345678901234567890123456789012
```

**Generar clave segura:**
```bash
# Opción 1: Con OpenSSL
openssl rand -base64 32 | cut -c1-32

# Opción 2: Con Python
python3 -c "import secrets; print(secrets.token_urlsafe(24)[:32])"

# Opción 3: Manual
# Crear string de exactamente 32 caracteres
```

**IMPORTANTE:**
- Debe tener **exactamente 32 caracteres**
- Cambiar en producción
- Nunca compartir ni subir a git
- Guardar de forma segura

### IPFS Configuration (OBLIGATORIO)

```bash
# Host del nodo IPFS
IPFS_HOST=ipfs              # Con docker-compose
# IPFS_HOST=localhost       # Si IPFS corre local

# Puerto del API
IPFS_PORT=5001

# Puerto del Gateway
IPFS_GATEWAY_PORT=8081
```

**Valores según entorno:**
- **Docker Compose**: `IPFS_HOST=ipfs` (nombre del servicio)
- **Local**: `IPFS_HOST=localhost`
- **Remoto**: `IPFS_HOST=ipfs.infura.io` o usar servicio IPFS de Alchemy/Pinata

### Server Configuration (OBLIGATORIO)

```bash
# Puerto del servidor
SERVER_PORT=8080

# Modo de Gin (debug, release, test)
GIN_MODE=debug
```

**Valores de GIN_MODE:**
- `debug` - Desarrollo (logs verbosos)
- `release` - Producción (optimizado)
- `test` - Testing

### Rate Limiting (OPCIONAL pero recomendado)

```bash
# Máximo de requests por ventana
RATE_LIMIT_REQUESTS=100

# Ventana en segundos
RATE_LIMIT_WINDOW=60
```

**Ejemplos:**
- `100/60` = 100 requests por minuto
- `1000/3600` = 1000 requests por hora
- `10/1` = 10 requests por segundo

## Variables Opcionales (Blockchain)

### Blockchain Configuration (OPCIONAL)

```bash
# Alchemy API Key (Recomendado 2025)
ALCHEMY_API_KEY=abc123def456ghi789

# O usar URL RPC personalizada (opcional)
BLOCKCHAIN_RPC_URL=https://eth-sepolia.g.alchemy.com/v2/YOUR_KEY

# Red blockchain
BLOCKCHAIN_NETWORK=sepolia

# Private key de tu wallet (sin 0x)
BLOCKCHAIN_PRIVATE_KEY=abcdef1234567890...

# Dirección del contrato
CONTRACT_ADDRESS=0x0000000000000000000000000000000000000000
```

**Cómo obtener Alchemy API Key:**
1. Ir a [dashboard.alchemy.com](https://dashboard.alchemy.com)
2. Crear cuenta gratuita (Sign Up)
3. Click "Create new app"
4. Configurar: Name, Chain: Ethereum, Network: Sepolia
5. Copiar API Key desde el dashboard

**Cómo obtener Private Key:**
1. Instalar [MetaMask](https://metamask.io)
2. Crear nueva wallet
3. Account details > Export Private Key
4. **Quitar el prefijo `0x`**
5. **NUNCA compartir esta clave**

**Cómo obtener fondos testnet:**
1. Copiar tu dirección Ethereum
2. Ir a [Sepolia Faucet](https://sepoliafaucet.com/)
3. Pegar dirección y solicitar fondos
4. Esperar confirmación (~1 minuto)

**Redes disponibles:**
- `sepolia` - Testnet recomendada
- `goerli` - Testnet alternativa
- `mainnet` - Producción (requiere ETH real)

### AWS Secrets Manager (PRODUCCIÓN)

```bash
# Habilitar AWS Secrets Manager
USE_AWS_SECRETS=true

# Nombre del secret
BLOCKCHAIN_PRIVATE_KEY_SECRET=blockchain-private-key
```

**Crear secret en AWS:**
```bash
aws secretsmanager create-secret \
  --name blockchain-private-key \
  --secret-string "tu_private_key_sin_0x" \
  --region us-east-1
```

## Perfiles de Configuración

### Desarrollo Local (sin blockchain)

```bash
# Mínimo necesario
AWS_ACCESS_KEY_ID=tu_access_key
AWS_SECRET_ACCESS_KEY=tu_secret_key
AWS_REGION=us-east-1
DYNAMODB_TABLE_NAME=transacciones-blockchain
ENCRYPTION_KEY=12345678901234567890123456789012
IPFS_HOST=ipfs
IPFS_PORT=5001
SERVER_PORT=8080
GIN_MODE=debug
```

### Desarrollo con Blockchain

Agregar a lo anterior:

```bash
ALCHEMY_API_KEY=tu_api_key
BLOCKCHAIN_NETWORK=sepolia
BLOCKCHAIN_PRIVATE_KEY=tu_private_key
```

### Producción

```bash
# Todas las anteriores +
GIN_MODE=release
USE_AWS_SECRETS=true
RATE_LIMIT_REQUESTS=1000
RATE_LIMIT_WINDOW=60

# Opcional
LOG_LEVEL=info
LOG_FORMAT=json
```

## Verificación de Configuración

### Verificar que .env esté correcto:

```bash
# 1. Verificar que existe
ls -la .env

# 2. Ver variables (¡cuidado en producción!)
cat .env

# 3. Probar configuración con docker-compose
docker-compose config

# 4. Iniciar y verificar logs
docker-compose up -d
docker-compose logs transaccion-blockchain
```

### Verificar servicios:

```bash
# Health check
curl http://localhost:8082/health

# Readiness (verifica conexiones)
curl http://localhost:8082/ready

# Debería retornar algo como:
# {
#   "status": "ready",
#   "checks": {
#     "ipfs": "healthy",
#     "blockchain": "healthy",
#     "dynamodb": "healthy"
#   }
# }
```

## Solución de Problemas

### Error: "ENCRYPTION_KEY debe tener 32 caracteres"

```bash
# Verificar longitud
echo -n "tu_clave" | wc -c

# Debe ser exactamente 32
# Generar nueva:
openssl rand -base64 32 | cut -c1-32
```

### Error: DynamoDB access denied

```bash
# Verificar credenciales
aws sts get-caller-identity

# Verificar tabla existe
aws dynamodb describe-table --table-name transacciones-blockchain

# Verificar permisos IAM
aws iam get-user-policy --user-name tu-usuario --policy-name DynamoDBAccess
```

### Error: IPFS no conecta

```bash
# Verificar IPFS corriendo
docker-compose ps ipfs

# Verificar puerto
curl -X POST http://localhost:5001/api/v0/version

# Si usas IPFS local
IPFS_HOST=localhost
```

### Error: Blockchain timeout

```bash
# Verificar Alchemy API Key
curl -X POST https://eth-sepolia.g.alchemy.com/v2/TU_API_KEY \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'

# Verificar fondos
# Usar Etherscan para tu dirección
```

## Ejemplo Completo de .env

```bash
# AWS
AWS_ACCESS_KEY_ID=AKIAI44QH8DHBEXAMPLE
AWS_SECRET_ACCESS_KEY=je7MtGbClwBF/2Zp9Utk/h3yCo8nvbEXAMPLEKEY
AWS_REGION=us-east-1
DYNAMODB_TABLE_NAME=transacciones-blockchain

# Blockchain (opcional)
ALCHEMY_API_KEY=abc123def456ghi789jkl012mno345pq
BLOCKCHAIN_NETWORK=sepolia
BLOCKCHAIN_PRIVATE_KEY=0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef
CONTRACT_ADDRESS=0x0000000000000000000000000000000000000000

# IPFS
IPFS_HOST=ipfs
IPFS_PORT=5001
IPFS_GATEWAY_PORT=8081

# Security
ENCRYPTION_KEY=MySecure32CharacterEncryptKey!!

# Server
SERVER_PORT=8080
GIN_MODE=debug

# Rate Limiting
RATE_LIMIT_REQUESTS=100
RATE_LIMIT_WINDOW=60

# AWS Secrets (production)
USE_AWS_SECRETS=false
BLOCKCHAIN_PRIVATE_KEY_SECRET=blockchain-private-key
```

## Seguridad Best Practices

1. **Nunca** subir `.env` a git
2. Usar `.env.example` o `env.example` para documentar
3. Rotar claves regularmente
4. Usar AWS Secrets Manager en producción
5. Limitar permisos IAM al mínimo necesario
6. Usar variables de entorno en CI/CD
7. Encriptar backups que contengan secrets
8. Auditar acceso a secrets regularmente

## Referencias

- [AWS IAM Best Practices](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html)
- [Alchemy Documentation](https://docs.alchemy.com/)
- [Infura Documentation](https://docs.infura.io/) (también soportado)
- [IPFS Documentation](https://docs.ipfs.tech/)
- [Go godotenv](https://github.com/joho/godotenv)

