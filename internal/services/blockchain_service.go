package services

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/edinfamous/blockchain-medisupply/pkg/contracts"
)

// BlockchainService maneja las operaciones con la blockchain
type BlockchainService struct {
	client          *ethclient.Client
	privateKey      *ecdsa.PrivateKey
	chainID         *big.Int
	contractAddress common.Address
	contract        *contracts.MediSupplyRegistry
}

// NewBlockchainService crea una nueva instancia de BlockchainService
// rpcURL puede ser Alchemy, Infura, o cualquier otro proveedor RPC compatible con Ethereum
// contractAddress es la dirección del smart contract desplegado (puede ser vacía para modo sin contrato)
func NewBlockchainService(rpcURL string, privateKeyHex string, contractAddress string) (*BlockchainService, error) {
	// Conectar al cliente Ethereum
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("error conectando a Ethereum: %w", err)
	}
	fmt.Println("Conectado a Ethereum")
	fmt.Println("rpcURL:", rpcURL)
	fmt.Println("privateKeyHex:", privateKeyHex)
	fmt.Println("contractAddress:", contractAddress)

	// Normalizar y validar private key
	privateKeyHex = normalizePrivateKey(privateKeyHex)
	
	// Validar longitud
	if len(privateKeyHex) != 64 {
		return nil, fmt.Errorf("private key debe tener 64 caracteres hexadecimales (32 bytes), tiene %d caracteres. Ejemplo: ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", len(privateKeyHex))
	}
	
	// Parsear private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error parseando private key: %w. Asegúrate de que sea un hex válido de 64 caracteres", err)
	}

	// Obtener chainID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error obteniendo chainID: %w", err)
	}

	service := &BlockchainService{
		client:     client,
		privateKey: privateKey,
		chainID:    chainID,
	}

	// Si hay una dirección de contrato, inicializar el contrato
	if contractAddress != "" {
		address := common.HexToAddress(contractAddress)
		if !common.IsHexAddress(contractAddress) {
			return nil, fmt.Errorf("dirección de contrato inválida: %s", contractAddress)
		}

		contract, err := contracts.NewMediSupplyRegistry(address, client)
		if err != nil {
			return nil, fmt.Errorf("error inicializando contrato: %w", err)
		}

		service.contractAddress = address
		service.contract = contract
	}

	return service, nil
}

// RegistrarEnBlockchain registra un hash en la blockchain usando el smart contract
// Si el contrato está configurado, usa el contrato. Si no, usa transacciones simples.
func (s *BlockchainService) RegistrarEnBlockchain(ctx context.Context, hash, cid string) (string, error) {
	// Si tenemos un contrato configurado, usarlo
	if s.contract != nil {
		return s.registrarConContrato(ctx, hash, cid)
	}

	// Fallback: usar transacción simple (modo compatible hacia atrás)
	return s.registrarConTransaccionSimple(ctx, hash, cid)
}

// registrarConContrato registra un hash usando el smart contract
func (s *BlockchainService) registrarConContrato(ctx context.Context, hash, cid string) (string, error) {
	// Preparar opciones de transacción
	opts, err := s.GetTransactionOpts(ctx)
	if err != nil {
		return "", fmt.Errorf("error obteniendo opciones de transacción: %w", err)
	}

	// Convertir hash string a bytes32
	hashBytes, err := hex.DecodeString(hash)
	if err != nil {
		return "", fmt.Errorf("error decodificando hash: %w", err)
	}
	if len(hashBytes) != 32 {
		return "", fmt.Errorf("hash debe tener 32 bytes, tiene %d", len(hashBytes))
	}
	var hashBytes32 [32]byte
	copy(hashBytes32[:], hashBytes)

	// Crear hash de transacción único (hash de hash + timestamp)
	hashTransaccionBytes := crypto.Keccak256([]byte(hash + cid))
	var hashTransaccion [32]byte
	copy(hashTransaccion[:], hashTransaccionBytes[:32])

	// Llamar al contrato
	tx, err := s.contract.RegistrarHash(opts, hashTransaccion, hashBytes32, cid)
	if err != nil {
		return "", fmt.Errorf("error registrando en contrato: %w", err)
	}

	// Esperar confirmación
	receipt, err := bind.WaitMined(ctx, s.client, tx)
	if err != nil {
		return "", fmt.Errorf("error esperando confirmación: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return "", fmt.Errorf("transacción falló en blockchain")
	}

	return tx.Hash().Hex(), nil
}

// registrarConTransaccionSimple registra usando una transacción simple (fallback)
func (s *BlockchainService) registrarConTransaccionSimple(ctx context.Context, hash, cid string) (string, error) {
	publicKey := s.privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("error casting public key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Obtener nonce
	nonce, err := s.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", fmt.Errorf("error obteniendo nonce: %w", err)
	}

	// Obtener gas price
	gasPrice, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		return "", fmt.Errorf("error obteniendo gas price: %w", err)
	}

	// Preparar datos: hash + CID concatenados
	data := []byte(fmt.Sprintf("%s:%s", hash, cid))

	// Crear transacción simple
	tx := types.NewTransaction(
		nonce,
		common.HexToAddress("0x0000000000000000000000000000000000000000"),
		big.NewInt(0),
		uint64(100000),
		gasPrice,
		data,
	)

	// Firmar transacción
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(s.chainID), s.privateKey)
	if err != nil {
		return "", fmt.Errorf("error firmando transacción: %w", err)
	}

	// Enviar transacción
	err = s.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", fmt.Errorf("error enviando transacción: %w", err)
	}

	return signedTx.Hash().Hex(), nil
}

// VerificarEnBlockchain verifica un hash contra la blockchain usando el smart contract
func (s *BlockchainService) VerificarEnBlockchain(ctx context.Context, txHash, hashEsperado string) (bool, error) {
	// Si tenemos un contrato, usar el método del contrato
	if s.contract != nil {
		return s.verificarConContrato(ctx, txHash, hashEsperado)
	}

	// Fallback: verificar en transacción simple
	return s.verificarConTransaccionSimple(ctx, txHash, hashEsperado)
}

// verificarConContrato verifica usando el smart contract
func (s *BlockchainService) verificarConContrato(ctx context.Context, txHash, hashEsperado string) (bool, error) {
	// Convertir hash esperado a bytes32
	hashBytes, err := hex.DecodeString(hashEsperado)
	if err != nil {
		return false, fmt.Errorf("error decodificando hash: %w", err)
	}
	if len(hashBytes) != 32 {
		return false, fmt.Errorf("hash debe tener 32 bytes")
	}
	var hashBytes32 [32]byte
	copy(hashBytes32[:], hashBytes)

	// El txHash es el hashTransaccion que usamos en el registro
	var hashTransaccion [32]byte
	copy(hashTransaccion[:], common.HexToHash(txHash).Bytes())

	// Llamar al contrato
	valido, err := s.contract.VerificarHash(&bind.CallOpts{Context: ctx}, hashTransaccion, hashBytes32)
	if err != nil {
		return false, fmt.Errorf("error verificando en contrato: %w", err)
	}

	return valido, nil
}

// verificarConTransaccionSimple verifica en una transacción simple (fallback)
func (s *BlockchainService) verificarConTransaccionSimple(ctx context.Context, txHash, hashEsperado string) (bool, error) {
	// Obtener la transacción
	tx, isPending, err := s.client.TransactionByHash(ctx, common.HexToHash(txHash))
	if err != nil {
		return false, fmt.Errorf("error obteniendo transacción: %w", err)
	}

	if isPending {
		return false, fmt.Errorf("transacción aún pendiente")
	}

	// Verificar que la transacción fue exitosa
	receipt, err := s.client.TransactionReceipt(ctx, common.HexToHash(txHash))
	if err != nil {
		return false, fmt.Errorf("error obteniendo receipt: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return false, fmt.Errorf("transacción falló en blockchain")
	}

	// Extraer datos de la transacción
	data := tx.Data()
	dataStr := string(data)

	// Verificar si el hash está en los datos
	// Formato: "hash:cid"
	return len(dataStr) > 0 && dataStr[:len(hashEsperado)] == hashEsperado, nil
}

// ObtenerBalance obtiene el balance de la cuenta
func (s *BlockchainService) ObtenerBalance(ctx context.Context) (*big.Int, error) {
	publicKey := s.privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	balance, err := s.client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo balance: %w", err)
	}

	return balance, nil
}

// VerificarConexion verifica la conexión con la blockchain
func (s *BlockchainService) VerificarConexion(ctx context.Context) error {
	_, err := s.client.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("error conectando a blockchain: %w", err)
	}
	return nil
}

// GetTransactionOpts obtiene las opciones de transacción configuradas
func (s *BlockchainService) GetTransactionOpts(ctx context.Context) (*bind.TransactOpts, error) {
	publicKey := s.privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := s.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo nonce: %w", err)
	}

	gasPrice, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo gas price: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(s.privateKey, s.chainID)
	if err != nil {
		return nil, fmt.Errorf("error creando transactor: %w", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	return auth, nil
}

// Close cierra la conexión con el cliente
func (s *BlockchainService) Close() {
	if s.client != nil {
		s.client.Close()
	}
}

// normalizePrivateKey normaliza una clave privada removiendo espacios, prefijos, etc.
// Retorna una clave de 64 caracteres hexadecimales (sin prefijo 0x)
func normalizePrivateKey(key string) string {
	// Remover espacios
	key = strings.TrimSpace(key)
	
	// Remover prefijo 0x si existe
	key = strings.TrimPrefix(key, "0x")
	key = strings.TrimPrefix(key, "0X")
	
	// Remover espacios adicionales
	key = strings.ReplaceAll(key, " ", "")
	key = strings.ReplaceAll(key, "\n", "")
	key = strings.ReplaceAll(key, "\t", "")
	
	// Convertir a minúsculas para consistencia
	key = strings.ToLower(key)
	
	return key
}
