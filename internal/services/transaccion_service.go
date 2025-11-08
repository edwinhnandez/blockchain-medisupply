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

	logicalHash, ethereumTxHash, err := s.blockchainService.RegistrarEnBlockchain(ctxBg, hash, cid)
	if err != nil {
		// Log error y actualizar estado
		fmt.Printf("🔴 Blockchain: Error registrando transacción %s en blockchain: %v\n", idTransaccion, err)
		if updateErr := s.dynamoDBService.ActualizarEstado(ctxBg, idTransaccion, "fallido"); updateErr != nil {
			fmt.Printf("🔴 Blockchain: Error actualizando estado en DynamoDB: %v\n", updateErr)
		}
		return
	}

	fmt.Printf("🟢 Blockchain: Transacción %s registrada en blockchain con hash lógico: %s, TxHash Ethereum: %s\n", idTransaccion, logicalHash, ethereumTxHash)

	// Actualizar con hash lógico y hash de transacción de Ethereum (esto también actualiza el estado a "confirmado")
	if err := s.dynamoDBService.ActualizarHashesBlockchain(ctxBg, idTransaccion, logicalHash, ethereumTxHash); err != nil {
		fmt.Printf("🔴 Blockchain: Error actualizando hashes de blockchain en DynamoDB: %v\n", err)
		return
	}

	fmt.Printf("🟢 Blockchain: Hashes de blockchain actualizados en DynamoDB para transacción %s (estado: confirmado)\n", idTransaccion)
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
	fmt.Printf("🔍 VERIFICAR: Iniciando verificación de integridad para ID: %s\n", idTransaccion)

	// 1. Obtener datos de DynamoDB
	transaccion, err := s.dynamoDBService.ObtenerTransaccion(ctx, idTransaccion)
	if err != nil {
		fmt.Printf("🔴 VERIFICAR: Error obteniendo transacción %s de DynamoDB: %v\n", idTransaccion, err)
		return nil, fmt.Errorf("error obteniendo transacción: %w", err)
	}
	fmt.Printf("🔍 VERIFICAR: Transacción obtenida de DynamoDB: ID=%s, CID=%s, HashEvento=%s, DatosEvento=%s, DirectionBlockchain=%s\n",
		transaccion.IDTransaction, transaccion.IPFSCid, transaccion.HashEvento, transaccion.DatosEvento, transaccion.DirectionBlockchain)

	response := &models.VerificacionResponse{
		IDTransaction: idTransaccion,
		Verificado:    false,
	}

	// 2. Verificar que tenga hash de blockchain
	if transaccion.DirectionBlockchain == "" {
		response.Mensaje = "Transacción aún no confirmada en blockchain"
		fmt.Printf("🔍 VERIFICAR: Transacción %s aún no confirmada en blockchain. DirectionBlockchain está vacío.\n", idTransaccion)
		return response, nil
	}

	// 3. Calcular hash local
	hashLocal := utils.CalcularHashTransaccion(transaccion)
	response.HashLocal = hashLocal
	fmt.Printf("🔍 VERIFICAR: Hash local calculado: %s\n", hashLocal)

	// 4. Verificar hash contra registro blockchain
	fmt.Printf("🔍 VERIFICAR: Verificando hash %s contra blockchain con DirectionBlockchain: %s\n", hashLocal, transaccion.DirectionBlockchain)
	verificadoBlockchain, err := s.blockchainService.VerificarEnBlockchain(ctx, transaccion.DirectionBlockchain, hashLocal)
	if err != nil {
		fmt.Printf("🔴 VERIFICAR: Error verificando en blockchain para %s: %v\n", idTransaccion, err)
		response.Mensaje = fmt.Sprintf("Error verificando blockchain: %v", err)
		return response, nil
	}
	fmt.Printf("🔍 VERIFICAR: Resultado verificación blockchain: %t\n", verificadoBlockchain)

	// 5. Recuperar datos de IPFS usando CID
	fmt.Printf("🔍 VERIFICAR: Recuperando datos de IPFS con CID: %s\n", transaccion.IPFSCid)
	datosIPFS, err := s.ipfsService.RecuperarJSON(ctx, transaccion.IPFSCid)
	if err != nil {
		fmt.Printf("🔴 VERIFICAR: Error recuperando de IPFS para %s (CID: %s): %v\n", idTransaccion, transaccion.IPFSCid, err)
		response.Mensaje = fmt.Sprintf("Error recuperando de IPFS: %v", err)
		return response, nil
	}
	fmt.Printf("🔍 VERIFICAR: Datos recuperados de IPFS (primeros 100 chars): %s...\n", datosIPFS[:min(100, len(datosIPFS))])

	// 6. Verificar que los datos de IPFS coincidan
	fmt.Printf("🔍 VERIFICAR: Comparando datos de IPFS con DatosEvento de DynamoDB.\n")
	fmt.Printf("🔍 VERIFICAR: Datos IPFS: %s\n", datosIPFS)
	fmt.Printf("🔍 VERIFICAR: Datos DynamoDB: %s\n", transaccion.DatosEvento)
	datosIPFSVerificados := (datosIPFS == transaccion.DatosEvento)
	response.DatosIPFSVerificados = datosIPFSVerificados
	response.HashBlockchain = transaccion.HashEvento
	fmt.Printf("🔍 VERIFICAR: Coincidencia de datos IPFS y DynamoDB: %t\n", datosIPFSVerificados)

	// 7. Resultado final
	response.Verificado = verificadoBlockchain && datosIPFSVerificados
	fmt.Printf("🔍 VERIFICAR: Resultado final de verificación (Blockchain && IPFS): %t\n", response.Verificado)

	if response.Verificado {
		response.Mensaje = "Transacción verificada exitosamente"
	} else {
		response.Mensaje = "Transacción NO verificada: discrepancia detectada"
		fmt.Printf("🔴 VERIFICAR: Discrepancia detectada para transacción %s. Blockchain: %t, IPFS: %t\n", idTransaccion, verificadoBlockchain, datosIPFSVerificados)
	}

	return response, nil
}

// min es una función auxiliar para obtener el mínimo de dos enteros
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
		EthereumTxHash:         transaccion.EthereumTxHash, // Incluir el hash de la transacción de Ethereum
		Timestamp:              transaccion.UpdatedAt.Format(time.RFC3339),
	}

	// Determinar mensaje según el estado
	switch transaccion.Estado {
	case "confirmado":
		if transaccion.DirectionBlockchain != "" {
			response.Mensaje = fmt.Sprintf("Transacción registrada exitosamente en blockchain. Hash Lógico: %s, TxHash Ethereum: %s", transaccion.DirectionBlockchain, transaccion.EthereumTxHash)
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
