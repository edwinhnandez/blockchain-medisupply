# Smart Contract Deployment Guide

## Descripción del Contrato

El contrato `MediSupplyRegistry` permite:
- Registrar hashes de transacciones médicas junto con CIDs de IPFS
- Verificar hashes registrados
- Consultar registros por cuenta
- Obtener información completa de registros

## Prerrequisitos

1. **Node.js y npm** instalados
2. **Hardhat** o **Truffle** para desarrollo y deployment
3. **MetaMask** o similar para gestión de cuentas
4. **Fondos de ETH** en la red de prueba (Sepolia, Goerli) o mainnet

## Instalación de Herramientas

### Opción 1: Usar Hardhat (Recomendado)

```bash
# Instalar Hardhat
npm install --save-dev hardhat

# Inicializar proyecto Hardhat (si es nuevo)
npx hardhat init

# Instalar dependencias adicionales
npm install @openzeppelin/contracts
npm install @nomicfoundation/hardhat-toolbox
```

### Opción 2: Usar Truffle

```bash
# Instalar Truffle globalmente
npm install -g truffle

# Instalar Ganache para red local
npm install -g ganache-cli
```

## Deployment

### 1. Configurar Hardhat (hardhat.config.js)

```javascript
require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: {
    version: "0.8.20",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200
      }
    }
  },
  networks: {
    sepolia: {
      url: `https://eth-sepolia.g.alchemy.com/v2/${process.env.ALCHEMY_API_KEY}`,
      accounts: [process.env.PRIVATE_KEY]
    },
    goerli: {
      url: `https://eth-goerli.g.alchemy.com/v2/${process.env.ALCHEMY_API_KEY}`,
      accounts: [process.env.PRIVATE_KEY]
    }
  }
};
```

### 2. Crear Script de Deployment (scripts/deploy.js)

```javascript
async function main() {
  const [deployer] = await ethers.getSigners();
  
  console.log("Desplegando contrato con la cuenta:", deployer.address);
  console.log("Balance de cuenta:", (await deployer.getBalance()).toString());

  const MediSupplyRegistry = await ethers.getContractFactory("MediSupplyRegistry");
  const registry = await MediSupplyRegistry.deploy();

  await registry.deployed();

  console.log("Contrato desplegado en:", registry.address);
  console.log("Guarda esta dirección para CONTRACT_ADDRESS en tu .env");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
```

### 3. Desplegar

```bash
# Configurar variables de entorno
export ALCHEMY_API_KEY="tu-api-key"
export PRIVATE_KEY="0xtu-private-key-sin-0x"

# Compilar contrato
npx hardhat compile

# Desplegar en Sepolia testnet
npx hardhat run scripts/deploy.js --network sepolia
```

### 4. Verificar Contrato (Opcional)

```bash
npx hardhat verify --network sepolia <CONTRACT_ADDRESS>
```

## Uso del Contrato

Una vez desplegado, copia la dirección del contrato y actualiza tu `.env`:

```bash
CONTRACT_ADDRESS=0x1234567890123456789012345678901234567890
```

## Generar ABI para Go

Después de compilar, el ABI estará en:
```
artifacts/contracts/MediSupplyRegistry.sol/MediSupplyRegistry.json
```

Copia el ABI a `contracts/abis/MediSupplyRegistry.json` para que Go pueda usarlo.

