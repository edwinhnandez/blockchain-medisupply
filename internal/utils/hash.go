package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/edinfamous/blockchain-medisupply/internal/models"
)

// CalcularHashTransaccion calcula el hash SHA-256 de una transacción
func CalcularHashTransaccion(transaccion *models.Transaccion) string {
	// Usar un formato de fecha explícito y consistente (RFC3339Nano) para evitar discrepancias
	// al guardar y recuperar de la base de datos.
	data := fmt.Sprintf("%s%s%s%s%s",
		transaccion.IDTransaction,
		transaccion.TipoEvento,
		transaccion.IDProducto,
		transaccion.FechaEvento.Format(time.RFC3339Nano),
		transaccion.DatosEvento,
	)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// CalcularHashDatos calcula el hash SHA-256 de datos arbitrarios
func CalcularHashDatos(datos string) string {
	hash := sha256.Sum256([]byte(datos))
	return hex.EncodeToString(hash[:])
}

// VerificarHash verifica si un hash coincide con los datos
func VerificarHash(datos, hashEsperado string) bool {
	hashCalculado := CalcularHashDatos(datos)
	return hashCalculado == hashEsperado
}
