package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/edinfamous/blockchain-medisupply/internal/models"
	"github.com/edinfamous/blockchain-medisupply/internal/utils"
	"github.com/edinfamous/blockchain-medisupply/pkg/validation"
)

// TransaccionService orquesta las operaciones de transacciones
type TransaccionService struct {
	blockchainService *BlockchainService
	ipfsService       *IPFSService
	dynamoDBService   *DynamoDBService
}

// NewTransaccionService crea una nueva instancia de TransaccionService
func NewTransaccionService(blockchain *BlockchainService, ipfs *IPFSService, dynamo *DynamoDBService) *TransaccionService {
	return &TransaccionService{
		blockchainService: blockchain,
		ipfsService:       ipfs,
		dynamoDBService:   dynamo,
	}
}

// RegistrarTransaccion registra una nueva transacción aplicando el patrón off-chain storage
func (s *TransaccionService) RegistrarTransaccion(ctx context.Context, req *models.TransaccionRequest) (*models.Transaccion, error) {
	// 1. Validar datos de entrada
	if err := validation.ValidateStruct(req); err != nil {
		return nil, fmt.Errorf("validación fallida: %w", err)
	}

	// 2. Crear transacción
	transaccion := &models.Transaccion{
		IDTransaction: uuid.New().String(),
		TipoEvento:    req.TipoEvento,
		IDProducto:    req.IDProducto,
		FechaEvento:   time.Now(),
		DatosEvento:   req.DatosEvento,
		ActorEmisor:   req.ActorEmisor,
		Estado:        "pendiente",
	}

	// 3. Almacenar datos detallados en IPFS (off-chain storage)
	cid, err := s.ipfsService.AlmacenarJSON(ctx, transaccion.DatosEvento)
	if err != nil {
		return nil, fmt.Errorf("error almacenando en IPFS: %w", err)
	}
	transaccion.IPFSCid = cid

	// 4. Calcular hash de integridad
	hash := utils.CalcularHashTransaccion(transaccion)
	transaccion.HashEvento = hash

	// 5. Registrar en DynamoDB primero
	if err := s.dynamoDBService.GuardarTransaccion(ctx, transaccion); err != nil {
		return nil, fmt.Errorf("error guardando en DynamoDB: %w", err)
	}

	// 6. Enviar transacción a blockchain (solo hash + CID) - asíncrono
	// En producción, esto podría ser una cola para no bloquear la respuesta
	go func() {
		ctxBg := context.Background()
		txHash, err := s.blockchainService.RegistrarEnBlockchain(ctxBg, hash, cid)
		if err != nil {
			// Log error y actualizar estado
			_ = s.dynamoDBService.ActualizarEstado(ctxBg, transaccion.IDTransaction, "fallido")
			return
		}

		// Actualizar con hash de blockchain
		_ = s.dynamoDBService.ActualizarHashBlockchain(ctxBg, transaccion.IDTransaction, txHash)
	}()

	return transaccion, nil
}

// ObtenerTransaccion obtiene una transacción por ID
func (s *TransaccionService) ObtenerTransaccion(ctx context.Context, idTransaccion string) (*models.Transaccion, error) {
	transaccion, err := s.dynamoDBService.ObtenerTransaccion(ctx, idTransaccion)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo transacción: %w", err)
	}
	return transaccion, nil
}

// VerificarIntegridad verifica la integridad de una transacción contra blockchain e IPFS
func (s *TransaccionService) VerificarIntegridad(ctx context.Context, idTransaccion string) (*models.VerificacionResponse, error) {
	// 1. Obtener datos de DynamoDB
	transaccion, err := s.dynamoDBService.ObtenerTransaccion(ctx, idTransaccion)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo transacción: %w", err)
	}

	response := &models.VerificacionResponse{
		IDTransaction: idTransaccion,
		Verificado:    false,
	}

	// 2. Verificar que tenga hash de blockchain
	if transaccion.DirectionBlockchain == "" {
		response.Mensaje = "Transacción aún no confirmada en blockchain"
		return response, nil
	}

	// 3. Calcular hash local
	hashLocal := utils.CalcularHashTransaccion(transaccion)
	response.HashLocal = hashLocal

	// 4. Verificar hash contra registro blockchain
	verificadoBlockchain, err := s.blockchainService.VerificarEnBlockchain(ctx, transaccion.DirectionBlockchain, hashLocal)
	if err != nil {
		response.Mensaje = fmt.Sprintf("Error verificando blockchain: %v", err)
		return response, nil
	}

	// 5. Recuperar datos de IPFS usando CID
	datosIPFS, err := s.ipfsService.RecuperarJSON(ctx, transaccion.IPFSCid)
	if err != nil {
		response.Mensaje = fmt.Sprintf("Error recuperando de IPFS: %v", err)
		return response, nil
	}

	// 6. Verificar que los datos de IPFS coincidan
	datosIPFSVerificados := (datosIPFS == transaccion.DatosEvento)
	response.DatosIPFSVerificados = datosIPFSVerificados
	response.HashBlockchain = transaccion.HashEvento

	// 7. Resultado final
	response.Verificado = verificadoBlockchain && datosIPFSVerificados

	if response.Verificado {
		response.Mensaje = "Transacción verificada exitosamente"
	} else {
		response.Mensaje = "Transacción NO verificada: discrepancia detectada"
	}

	return response, nil
}

// ListarTransacciones lista todas las transacciones
func (s *TransaccionService) ListarTransacciones(ctx context.Context, limit int32) ([]*models.Transaccion, error) {
	return s.dynamoDBService.ListarTransacciones(ctx, limit)
}

// ObtenerTransaccionesPorProducto obtiene todas las transacciones de un producto
func (s *TransaccionService) ObtenerTransaccionesPorProducto(ctx context.Context, idProducto string) ([]*models.Transaccion, error) {
	return s.dynamoDBService.ObtenerTransaccionesPorProducto(ctx, idProducto)
}
