# Guía Completa: Cuentas Ethereum y Generación de Claves

## ¿Qué es una Cuenta Ethereum?

Una cuenta Ethereum es una identidad en la blockchain que te permite:
- Recibir y enviar transacciones
- Interactuar con smart contracts
- Almacenar fondos (ETH y tokens ERC-20)

**Importante:** En Ethereum, las cuentas NO se crean en Solidity. Las cuentas existen independientemente del código del contrato. Lo que puedes hacer en Solidity es:
1. **Identificar** cuentas que interactúan con tu contrato
2. **Almacenar datos** asociados a cuentas
3. **Implementar lógica** basada en quién llama al contrato

## Componentes de una Cuenta Ethereum

### 1. Private Key (Clave Privada)
- **64 caracteres hexadecimales** (256 bits)
- **Mantenerla SECRETA** - quien la posee controla la cuenta
- Se usa para **firmar transacciones**

Ejemplo:
```
0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
```

### 2. Public Key (Clave Pública)
- **130 caracteres hexadecimales** (520 bits)
- Deriva de la private key usando **criptografía ECDSA**
- Se puede compartir públicamente
- Se usa para **verificar firmas**

### 3. Address (Dirección)
- **42 caracteres** (20 bytes = 160 bits)
- Deriva de la public key usando **Keccak-256**
- Formato: `0x` + 40 caracteres hexadecimales
- Es tu **identidad pública** en la blockchain

Ejemplo:
```
0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb0
```

## Generación de Claves

### Método 1: Usando Go (geth)

```go
package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Generar nueva clave privada
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// Obtener clave privada en formato hexadecimal
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("Private Key:", fmt.Sprintf("0x%x", privateKeyBytes))

	// Obtener clave pública
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("Public Key:", fmt.Sprintf("0x%x", publicKeyBytes))

	// Obtener dirección (address)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("Address:", address.Hex())

	// Generar keystore (archivo encriptado)
	ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.ImportECDSA(privateKey, "password123")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Account imported:", account.Address.Hex())
}
```

**Ejecutar:**
```bash
go run generate_key.go
```

### Método 2: Usando Web3.js / Ethers.js

```javascript
const { ethers } = require("ethers");

// Generar nueva wallet
const wallet = ethers.Wallet.createRandom();

console.log("Private Key:", wallet.privateKey);
console.log("Address:", wallet.address);
console.log("Mnemonic:", wallet.mnemonic.phrase);
```

### Método 3: Usando MetaMask

1. Instalar MetaMask extension
2. Crear nueva cuenta
3. **Mostrar clave privada:**
   - Settings → Security & Privacy
   - "Show Secret Recovery Phrase"
   - **ADVERTENCIA:** Nunca compartas esta clave

### Método 4: Usando Geth (Command Line)

```bash
# Crear nueva cuenta
geth account new

# Esto te pedirá una contraseña y generará:
# - Private key (encriptada)
# - Address
```

### Método 5: Usando OpenSSL y herramientas

```bash
# Generar clave privada usando OpenSSL
openssl ecparam -name secp256k1 -genkey -noout | openssl ec -text -noout | grep priv -A 3 | tail -n +2 | tr -d '\n[:space:]' | sed 's/^00//'

# Luego derivar address desde la clave privada
```

## Importar Clave Privada en tu Aplicación

### En tu `.env`:

```bash
# Clave privada SIN el prefijo 0x
BLOCKCHAIN_PRIVATE_KEY=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# O CON el prefijo 0x (ambos funcionan)
BLOCKCHAIN_PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
```

### En Go:

```go
import (
	"github.com/ethereum/go-ethereum/crypto"
)

// Cargar desde variable de entorno
privateKeyHex := os.Getenv("BLOCKCHAIN_PRIVATE_KEY")

// Remover prefijo 0x si existe
if strings.HasPrefix(privateKeyHex, "0x") {
	privateKeyHex = privateKeyHex[2:]
}

// Parsear clave privada
privateKey, err := crypto.HexToECDSA(privateKeyHex)
if err != nil {
	log.Fatal(err)
}

// Obtener address
publicKey := privateKey.Public()
publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
address := crypto.PubkeyToAddress(*publicKeyECDSA)

fmt.Println("Address:", address.Hex())
```

## Uso de Cuentas en Smart Contracts

Aunque las cuentas no se crean en Solidity, puedes usar información de la cuenta llamadora:

```solidity
// MediSupplyRegistry.sol

contract MediSupplyRegistry {
    // Obtener la cuenta que llamó la función
    address public registrador = msg.sender;
    
    // Mapear cuentas a registros
    mapping(address => bytes32[]) public registrosPorCuenta;
    
    function registrarHash(bytes32 hash, string memory cid) public {
        // msg.sender es la cuenta que llamó esta función
        registrosPorCuenta[msg.sender].push(hash);
        
        // Registrar quien lo registró
        registros[hash] = Registro({
            registrador: msg.sender,  // <-- Cuenta del llamador
            timestamp: block.timestamp,
            cid: cid
        });
    }
    
    // Verificar si una cuenta tiene permiso
    function soloOwner() public view {
        require(msg.sender == owner, "Solo el owner puede hacer esto");
    }
}
```

## Variables Especiales en Solidity

En Solidity, tienes acceso a información sobre la cuenta que llama:

- `msg.sender` - Address de la cuenta que llamó la función
- `tx.origin` - Address de la cuenta que inició la transacción
- `msg.value` - Cantidad de ETH enviada con la transacción
- `block.coinbase` - Address del minero del bloque actual
- `block.timestamp` - Timestamp del bloque actual

## Seguridad: Mejores Prácticas

### ✅ HACER:

1. **Nunca commits** tu clave privada en Git
2. Usa **variables de entorno** o **AWS Secrets Manager**
3. Usa **diferentes cuentas** para desarrollo y producción
4. **Encripta** claves privadas cuando las almacenes
5. Usa **hardware wallets** (Ledger, Trezor) para producción

### ❌ NO HACER:

1. ❌ Compartir tu clave privada
2. ❌ Usar la misma cuenta para todo
3. ❌ Guardar claves en texto plano
4. ❌ Usar claves de producción en desarrollo
5. ❌ Confiar en servicios online para generar claves

## Obtener Fondos de Prueba (Testnet)

### Sepolia Testnet (Recomendado):

1. Obtener ETH de faucet:
   - https://sepoliafaucet.com/
   - https://faucet.quicknode.com/ethereum/sepolia
   - https://www.alchemy.com/faucets/ethereum-sepolia

2. Conectar MetaMask a Sepolia:
   - Network Name: Sepolia
   - RPC URL: https://rpc.sepolia.org
   - Chain ID: 11155111
   - Currency Symbol: ETH

3. Solicitar ETH:
   - Ingresa tu address en el faucet
   - Espera confirmación

## Script de Generación Completo

Crea `scripts/generate_account.go`:

```go
package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Generar nueva clave
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// Información de la cuenta
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println("=== NUEVA CUENTA ETHEREUM ===")
	fmt.Println()
	fmt.Printf("Private Key: 0x%x\n", privateKeyBytes)
	fmt.Printf("Public Key:  0x%x\n", publicKeyBytes)
	fmt.Printf("Address:     %s\n", address.Hex())
	fmt.Println()
	fmt.Println("⚠️  ADVERTENCIA:")
	fmt.Println("   - Guarda esta clave privada en un lugar SEGURO")
	fmt.Println("   - NUNCA compartas tu clave privada")
	fmt.Println("   - Usa esta clave solo para desarrollo/testnet")
	fmt.Println()

	// Opcional: guardar en keystore
	ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.ImportECDSA(privateKey, "password123")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("✅ Keystore guardado en: ./keystore/%s\n", account.Address.Hex())
}
```

**Ejecutar:**
```bash
go run scripts/generate_account.go
```

## Resumen

1. **Cuentas NO se crean en Solidity** - se crean antes de usar la blockchain
2. **Private Key** → deriva → **Public Key** → deriva → **Address**
3. Usa herramientas como **geth**, **MetaMask**, o **Go** para generar cuentas
4. **msg.sender** en Solidity te da la cuenta que llama al contrato
5. **Nunca compartas** tu clave privada
6. Usa **testnet** para desarrollo y **hardware wallets** para producción

## Próximos Pasos

1. Genera una cuenta de desarrollo usando el script
2. Obtén ETH de testnet del faucet
3. Despliega el contrato usando esa cuenta
4. Configura `CONTRACT_ADDRESS` en tu `.env`
5. Usa `BLOCKCHAIN_PRIVATE_KEY` con la clave privada generada

