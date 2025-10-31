package services

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// BlockchainService maneja las operaciones con la blockchain
type BlockchainService struct {
	client     *ethclient.Client
	privateKey *ecdsa.PrivateKey
	chainID    *big.Int
	network    string
}

// NewBlockchainService crea una nueva instancia de BlockchainService
// rpcURL puede ser Alchemy, Infura, o cualquier otro proveedor RPC compatible con Ethereum
func NewBlockchainService(rpcURL string, privateKeyHex string) (*BlockchainService, error) {
	// Conectar al cliente Ethereum
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("error conectando a Ethereum: %w", err)
	}

	// Parsear private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error parseando private key: %w", err)
	}

	// Obtener chainID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error obteniendo chainID: %w", err)
	}

	return &BlockchainService{
		client:     client,
		privateKey: privateKey,
		chainID:    chainID,
	}, nil
}

// RegistrarEnBlockchain registra un hash en la blockchain
// En una implementación real, esto interactuaría con un smart contract
func (s *BlockchainService) RegistrarEnBlockchain(ctx context.Context, hash, cid string) (string, error) {
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

	// Crear transacción
	// En producción, esto sería una llamada a un smart contract
	// Por ahora, guardamos los datos en el campo data de la transacción
	tx := types.NewTransaction(
		nonce,
		common.HexToAddress("0x0000000000000000000000000000000000000000"), // Dirección nula o del contrato
		big.NewInt(0),
		uint64(100000), // Gas limit
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

	// Retornar hash de la transacción
	return signedTx.Hash().Hex(), nil
}

// VerificarEnBlockchain verifica un hash contra la blockchain
func (s *BlockchainService) VerificarEnBlockchain(ctx context.Context, txHash, hashEsperado string) (bool, error) {
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
