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
	fmt.Println("🟢 Service: RegistrarTransaccion - INICIADO")
	fmt.Printf("🟢 Service: Request recibido - TipoEvento: %s, IDProducto: %s, ActorEmisor: %s\n", req.TipoEvento, req.IDProducto, req.ActorEmisor)

	// 1. Validar datos de entrada
	fmt.Println("🟢 Service: Validando datos de entrada...")
	if err := validation.ValidateStruct(req); err != nil {
		fmt.Printf("🔴 Service: Error en validación: %v\n", err)
		return nil, fmt.Errorf("validación fallida: %w", err)
	}
	fmt.Println("🟢 Service: Validación exitosa")

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
	fmt.Println("Transacción creada con ID:", transaccion.IDTransaction)

	// 3. Almacenar datos detallados en IPFS (off-chain storage)
	// Crear contexto con timeout más largo para IPFS (60 segundos)
	ipfsCtx, ipfsCancel := context.WithTimeout(ctx, 60*time.Second)
	defer ipfsCancel()

	// Verificar conectividad con IPFS antes de intentar almacenar (con timeout corto de 5 segundos)
	fmt.Println("🟢 Service: Verificando conectividad con IPFS...")
	checkCtx, checkCancel := context.WithTimeout(ctx, 5*time.Second)
	defer checkCancel()
	if err := s.ipfsService.VerificarConexion(checkCtx); err != nil {
		fmt.Printf("🔴 Service: IPFS no está disponible: %v\n", err)
		if checkCtx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("timeout al verificar IPFS (5s): verifique que IPFS esté corriendo y accesible en %s:%s", s.ipfsService.GetHost(), s.ipfsService.GetPort())
		}
		return nil, fmt.Errorf("IPFS no está disponible: verifique que IPFS esté corriendo en %s:%s. Error: %w", s.ipfsService.GetHost(), s.ipfsService.GetPort(), err)
	}
	fmt.Println("🟢 Service: IPFS está disponible, intentando almacenar...")

	cid, err := s.ipfsService.AlmacenarJSON(ipfsCtx, transaccion.DatosEvento)
	if err != nil {
		// Verificar si es un timeout
		if ipfsCtx.Err() == context.DeadlineExceeded {
			fmt.Printf("🔴 Service: Timeout al almacenar en IPFS: %v\n", err)
			return nil, fmt.Errorf("timeout al almacenar en IPFS (60s): verifique que IPFS esté corriendo y accesible en %s:%s", s.ipfsService.GetHost(), s.ipfsService.GetPort())
		}
		fmt.Printf("🔴 Service: Error almacenando en IPFS: %v\n", err)
		return nil, fmt.Errorf("error almacenando en IPFS: %w", err)
	}
	fmt.Printf("🟢 Service: Datos almacenados en IPFS con CID: %s\n", cid)
	transaccion.IPFSCid = cid
	fmt.Println("Transacción con IPFSCid:", transaccion.IPFSCid)
	fmt.Println("Transacción con DatosEvento:", transaccion.DatosEvento)
	fmt.Println("Transacción con ActorEmisor:", transaccion.ActorEmisor)
	fmt.Println("Transacción con Estado:", transaccion.Estado)
	fmt.Println("Transacción con CreatedAt:", transaccion.CreatedAt)
	fmt.Println("Transacción con UpdatedAt:", transaccion.UpdatedAt)
	fmt.Println("Transacción con TipoEvento:", transaccion.TipoEvento)
	fmt.Println("Transacción con IDProducto:", transaccion.IDProducto)
	fmt.Println("Transacción con FechaEvento:", transaccion.FechaEvento)
	fmt.Println("Transacción con HashEvento:", transaccion.HashEvento)
	fmt.Println("Transacción con DirectionBlockchain:", transaccion.DirectionBlockchain)
	fmt.Println("Transacción con FirmaDigital:", transaccion.FirmaDigital)
	fmt.Println("Transacción con IDTransaction:", transaccion.IDTransaction)

	// 4. Calcular hash de integridad
	hash := utils.CalcularHashTransaccion(transaccion)
	fmt.Println("Hash de integridad calculado:", hash)
	transaccion.HashEvento = hash

	// 5. Registrar en DynamoDB primero
	// Crear contexto con timeout para DynamoDB (30 segundos)
	dynamoCtx, dynamoCancel := context.WithTimeout(ctx, 30*time.Second)
	defer dynamoCancel()

	fmt.Println("🟢 Service: Intentando guardar en DynamoDB...")
	if err := s.dynamoDBService.GuardarTransaccion(dynamoCtx, transaccion); err != nil {
		// Verificar si es un timeout
		if dynamoCtx.Err() == context.DeadlineExceeded {
			fmt.Printf("🔴 Service: Timeout al guardar en DynamoDB: %v\n", err)
			return nil, fmt.Errorf("timeout al guardar en DynamoDB (30s): verifique la conexión con AWS DynamoDB")
		}
		fmt.Printf("🔴 Service: Error guardando en DynamoDB: %v\n", err)
		return nil, fmt.Errorf("error guardando en DynamoDB: %w", err)
	}
	fmt.Println("🟢 Service: Transacción guardada exitosamente en DynamoDB")

	// 6. Enviar transacción a blockchain (solo hash + CID) - asíncrono
	go s.registrarEnBlockchainAsync(transaccion.IDTransaction, hash, cid)

	return transaccion, nil
}

// registrarEnBlockchainAsync registra la transacción en blockchain de forma asíncrona
// Esta función se ejecuta en un goroutine separado para no bloquear la respuesta HTTP
func (s *TransaccionService) registrarEnBlockchainAsync(idTransaccion, hash, cid string) {
	ctxBg := context.Background()

	fmt.Printf("🟡 Blockchain: Iniciando registro asíncrono para transacción %s\n", idTransaccion)
	fmt.Printf("🟡 Blockchain: Hash: %s, CID: %s\n", hash, cid)

	// Solo proceder si el servicio de blockchain está disponible
	if s.blockchainService == nil {
		fmt.Println("⚠️  Blockchain: Servicio de blockchain no disponible, saltando registro")
		return
	}

	txHash, err := s.blockchainService.RegistrarEnBlockchain(ctxBg, hash, cid)
	if err != nil {
		// Log error y actualizar estado
		fmt.Printf("🔴 Blockchain: Error registrando transacción %s en blockchain: %v\n", idTransaccion, err)
		if updateErr := s.dynamoDBService.ActualizarEstado(ctxBg, idTransaccion, "fallido"); updateErr != nil {
			fmt.Printf("🔴 Blockchain: Error actualizando estado en DynamoDB: %v\n", updateErr)
		}
		return
	}

	fmt.Printf("🟢 Blockchain: Transacción %s registrada en blockchain con txHash: %s\n", idTransaccion, txHash)

	// Actualizar con hash de blockchain (esto también actualiza el estado a "confirmado")
	if err := s.dynamoDBService.ActualizarHashBlockchain(ctxBg, idTransaccion, txHash); err != nil {
		fmt.Printf("🔴 Blockchain: Error actualizando hash de blockchain en DynamoDB: %v\n", err)
		return
	}

	fmt.Printf("🟢 Blockchain: Hash de blockchain actualizado en DynamoDB para transacción %s (estado: confirmado)\n", idTransaccion)
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

// ObtenerEstadoBlockchain obtiene el estado del registro en blockchain de una transacción
// Retorna información sobre si la transacción fue registrada exitosamente en blockchain
func (s *TransaccionService) ObtenerEstadoBlockchain(ctx context.Context, idTransaccion string) (*models.EstadoBlockchainResponse, error) {
	// Obtener la transacción desde DynamoDB
	transaccion, err := s.dynamoDBService.ObtenerTransaccion(ctx, idTransaccion)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo transacción: %w", err)
	}

	response := &models.EstadoBlockchainResponse{
		IDTransaction:          idTransaccion,
		Estado:                 transaccion.Estado,
		RegistradoEnBlockchain: transaccion.DirectionBlockchain != "",
		DirectionBlockchain:    transaccion.DirectionBlockchain,
		Timestamp:              transaccion.UpdatedAt.Format(time.RFC3339),
	}

	// Determinar mensaje según el estado
	switch transaccion.Estado {
	case "confirmado":
		if transaccion.DirectionBlockchain != "" {
			response.Mensaje = fmt.Sprintf("Transacción registrada exitosamente en blockchain. TxHash: %s", transaccion.DirectionBlockchain)
		} else {
			response.Mensaje = "Transacción confirmada pero sin hash de blockchain (puede estar en proceso)"
		}
	case "fallido":
		response.Mensaje = "Error al registrar la transacción en blockchain"
	case "pendiente":
		if transaccion.DirectionBlockchain != "" {
			response.Mensaje = "Transacción registrada en blockchain pero aún en estado pendiente"
		} else {
			response.Mensaje = "Transacción pendiente de registro en blockchain"
		}
	default:
		response.Mensaje = fmt.Sprintf("Estado desconocido: %s", transaccion.Estado)
	}

	return response, nil
}
