package tests

import (
	"time"

	"github.com/edinfamous/blockchain-medisupply/internal/models"
)

// GetMockTransaccionRequest retorna una solicitud de transacción de prueba
func GetMockTransaccionRequest() *models.TransaccionRequest {
	return &models.TransaccionRequest{
		TipoEvento:  "fabricacion",
		IDProducto:  "PROD-TEST-001",
		DatosEvento: `{"lote": "12345", "fecha_fabricacion": "2024-01-15", "cantidad": 1000}`,
		ActorEmisor: "Laboratorio Test SA",
	}
}

// GetMockTransaccion retorna una transacción de prueba completa
func GetMockTransaccion() *models.Transaccion {
	return &models.Transaccion{
		IDTransaction:       "TX-TEST-001",
		TipoEvento:          "fabricacion",
		IDProducto:          "PROD-TEST-001",
		FechaEvento:         time.Now(),
		DatosEvento:         `{"lote": "12345", "fecha_fabricacion": "2024-01-15"}`,
		HashEvento:          "abc123def456",
		DirectionBlockchain: "0x1234567890abcdef",
		IPFSCid:             "QmTest123",
		ActorEmisor:         "Laboratorio Test SA",
		Estado:              "confirmado",
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
}

// GetMockTransaccionFabricacion retorna transacción de fabricación
func GetMockTransaccionFabricacion(idProducto string) *models.Transaccion {
	tx := GetMockTransaccion()
	tx.TipoEvento = "fabricacion"
	tx.IDProducto = idProducto
	tx.DatosEvento = `{"lote": "LOT-001", "planta": "Planta A", "calidad": "premium"}`
	return tx
}

// GetMockTransaccionDistribucion retorna transacción de distribución
func GetMockTransaccionDistribucion(idProducto string) *models.Transaccion {
	tx := GetMockTransaccion()
	tx.TipoEvento = "distribucion"
	tx.IDProducto = idProducto
	tx.DatosEvento = `{"transportista": "LogiMed", "destino": "Farmacia Central", "temperatura": "2-8°C"}`
	return tx
}

// GetMockTransaccionRecepcion retorna transacción de recepción
func GetMockTransaccionRecepcion(idProducto string) *models.Transaccion {
	tx := GetMockTransaccion()
	tx.TipoEvento = "recepcion"
	tx.IDProducto = idProducto
	tx.DatosEvento = `{"receptor": "Farmacia Central", "condicion": "optima", "inspeccionado_por": "Inspector A"}`
	return tx
}
