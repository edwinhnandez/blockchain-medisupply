package services

import (
	"context"
	"fmt"
	"time"

	"github.com/edinfamous/blockchain-medisupply/internal/models"
)

// OracleService implementa el patrón Oracle para exponer datos verificados
type OracleService struct {
	transaccionService *TransaccionService
	dynamoDBService    *DynamoDBService
}

// NewOracleService crea una nueva instancia de OracleService
func NewOracleService(transaccion *TransaccionService, dynamo *DynamoDBService) *OracleService {
	return &OracleService{
		transaccionService: transaccion,
		dynamoDBService:    dynamo,
	}
}

// ObtenerDatosVerificados obtiene y verifica el historial completo de un producto (patrón Oracle)
func (s *OracleService) ObtenerDatosVerificados(ctx context.Context, idProducto string) (*models.OracleDataResponse, error) {
	// 1. Consultar transacciones relacionadas en DynamoDB
	transacciones, err := s.dynamoDBService.ObtenerTransaccionesPorProducto(ctx, idProducto)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo transacciones: %w", err)
	}

	if len(transacciones) == 0 {
		return nil, fmt.Errorf("no se encontraron transacciones para el producto: %s", idProducto)
	}

	// 2. Construir respuesta del Oracle
	response := &models.OracleDataResponse{
		IDProducto:          idProducto,
		UltimaActualizacion: time.Now(),
		Historial:           make([]models.EventoVerificado, 0),
		Metadata:            make(map[string]string),
	}

	// 3. Validar cada transacción contra blockchain
	cadenaVerificada := true
	var ultimaFecha time.Time

	for _, transaccion := range transacciones {
		// Verificar integridad de cada transacción
		verificacion, err := s.transaccionService.VerificarIntegridad(ctx, transaccion.IDTransaction)

		evento := models.EventoVerificado{
			IDEvento:              transaccion.IDTransaction,
			TipoEvento:            transaccion.TipoEvento,
			Fecha:                 transaccion.FechaEvento,
			ReferenciaBlockchain:  transaccion.DirectionBlockchain,
			IPFSCid:               transaccion.IPFSCid,
			ActorEmisor:           transaccion.ActorEmisor,
			ResultadoVerificacion: false,
		}

		if err != nil {
			evento.ErrorVerificacion = err.Error()
			cadenaVerificada = false
		} else {
			evento.ResultadoVerificacion = verificacion.Verificado
			if !verificacion.Verificado {
				evento.ErrorVerificacion = verificacion.Mensaje
				cadenaVerificada = false
			}
		}

		response.Historial = append(response.Historial, evento)

		// Actualizar última fecha
		if transaccion.FechaEvento.After(ultimaFecha) {
			ultimaFecha = transaccion.FechaEvento
		}
	}

	// 4. Establecer estado basado en verificaciones
	response.CadenaVerificada = cadenaVerificada
	if cadenaVerificada {
		response.Estado = "verificado"
	} else {
		response.Estado = "no_verificado"
	}

	// 5. Agregar metadata útil
	response.Metadata["total_eventos"] = fmt.Sprintf("%d", len(transacciones))
	response.Metadata["ultima_actualizacion"] = ultimaFecha.Format(time.RFC3339)
	response.Metadata["tipo_ultimo_evento"] = transacciones[len(transacciones)-1].TipoEvento

	return response, nil
}

// ObtenerHistorialVerificado obtiene el historial verificado de un producto
func (s *OracleService) ObtenerHistorialVerificado(ctx context.Context, idProducto string) (*models.HistorialVerificado, error) {
	// 1. Consultar transacciones relacionadas en DynamoDB
	transacciones, err := s.dynamoDBService.ObtenerTransaccionesPorProducto(ctx, idProducto)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo transacciones: %w", err)
	}

	// 2. Construir historial
	historial := &models.HistorialVerificado{
		IDProducto:    idProducto,
		TotalEventos:  len(transacciones),
		Verificados:   0,
		NoVerificados: 0,
		Eventos:       make([]models.EventoVerificado, 0),
		FechaConsulta: time.Now(),
	}

	// 3. Validar cada una contra blockchain y construir historial
	for _, transaccion := range transacciones {
		verificacion, err := s.transaccionService.VerificarIntegridad(ctx, transaccion.IDTransaction)

		eventoVerificado := models.EventoVerificado{
			IDEvento:             transaccion.IDTransaction,
			TipoEvento:           transaccion.TipoEvento,
			Fecha:                transaccion.FechaEvento,
			ReferenciaBlockchain: transaccion.DirectionBlockchain,
			IPFSCid:              transaccion.IPFSCid,
			ActorEmisor:          transaccion.ActorEmisor,
		}

		if err != nil {
			eventoVerificado.ResultadoVerificacion = false
			eventoVerificado.ErrorVerificacion = err.Error()
			historial.NoVerificados++
		} else {
			eventoVerificado.ResultadoVerificacion = verificacion.Verificado
			if verificacion.Verificado {
				historial.Verificados++
			} else {
				historial.NoVerificados++
				eventoVerificado.ErrorVerificacion = verificacion.Mensaje
			}
		}

		historial.Eventos = append(historial.Eventos, eventoVerificado)
	}

	return historial, nil
}

// ValidarCadenaSupply valida que la cadena de suministro sea coherente
func (s *OracleService) ValidarCadenaSupply(ctx context.Context, idProducto string) (bool, []string, error) {
	transacciones, err := s.dynamoDBService.ObtenerTransaccionesPorProducto(ctx, idProducto)
	if err != nil {
		return false, nil, fmt.Errorf("error obteniendo transacciones: %w", err)
	}

	// Validar orden lógico de eventos
	// TODO: Implementar validación de secuencia de eventos esperados
	// eventosEsperados := []string{"fabricacion", "distribucion", "recepcion", "verificacion"}
	errores := make([]string, 0)

	// Verificar que exista al menos fabricación
	tieneFabricacion := false
	for _, tx := range transacciones {
		if tx.TipoEvento == "fabricacion" {
			tieneFabricacion = true
			break
		}
	}

	if !tieneFabricacion {
		errores = append(errores, "Falta evento de fabricación")
	}

	// Validar orden cronológico
	var ultimaFecha time.Time
	for i, tx := range transacciones {
		if i > 0 && tx.FechaEvento.Before(ultimaFecha) {
			errores = append(errores, fmt.Sprintf("Evento %s tiene fecha anterior al evento previo", tx.IDTransaction))
		}
		ultimaFecha = tx.FechaEvento
	}

	return len(errores) == 0, errores, nil
}
