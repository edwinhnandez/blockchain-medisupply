package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/edinfamous/blockchain-medisupply/internal/models"
	"github.com/edinfamous/blockchain-medisupply/internal/utils"
)

func TestCalcularHashTransaccion(t *testing.T) {
	t.Run("Hash consistente", func(t *testing.T) {
		fecha := time.Now()
		transaccion := &models.Transaccion{
			IDTransaction: "TX-001",
			TipoEvento:    "fabricacion",
			IDProducto:    "PROD-001",
			FechaEvento:   fecha,
			DatosEvento:   `{"lote": "12345"}`,
		}

		hash1 := utils.CalcularHashTransaccion(transaccion)
		hash2 := utils.CalcularHashTransaccion(transaccion)

		assert.Equal(t, hash1, hash2, "El hash debe ser consistente")
		assert.Len(t, hash1, 64, "Hash SHA-256 debe tener 64 caracteres hex")
	})

	t.Run("Hash diferente para datos diferentes", func(t *testing.T) {
		fecha := time.Now()
		transaccion1 := &models.Transaccion{
			IDTransaction: "TX-001",
			TipoEvento:    "fabricacion",
			IDProducto:    "PROD-001",
			FechaEvento:   fecha,
			DatosEvento:   `{"lote": "12345"}`,
		}

		transaccion2 := &models.Transaccion{
			IDTransaction: "TX-002",
			TipoEvento:    "fabricacion",
			IDProducto:    "PROD-001",
			FechaEvento:   fecha,
			DatosEvento:   `{"lote": "12345"}`,
		}

		hash1 := utils.CalcularHashTransaccion(transaccion1)
		hash2 := utils.CalcularHashTransaccion(transaccion2)

		assert.NotEqual(t, hash1, hash2, "Hashes deben ser diferentes para transacciones diferentes")
	})
}

func TestCalcularHashDatos(t *testing.T) {
	t.Run("Hash de datos", func(t *testing.T) {
		datos := "test data"
		hash := utils.CalcularHashDatos(datos)

		assert.Len(t, hash, 64)
		assert.NotEmpty(t, hash)
	})
}

func TestVerificarHash(t *testing.T) {
	t.Run("Verificación exitosa", func(t *testing.T) {
		datos := "test data"
		hash := utils.CalcularHashDatos(datos)

		verificado := utils.VerificarHash(datos, hash)
		assert.True(t, verificado)
	})

	t.Run("Verificación fallida", func(t *testing.T) {
		datos := "test data"
		hashIncorrecto := "0000000000000000000000000000000000000000000000000000000000000000"

		verificado := utils.VerificarHash(datos, hashIncorrecto)
		assert.False(t, verificado)
	})
}
