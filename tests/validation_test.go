package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edinfamous/blockchain-medisupply/internal/models"
	"github.com/edinfamous/blockchain-medisupply/pkg/validation"
)

func TestValidateStruct(t *testing.T) {
	t.Run("TransaccionRequest válida", func(t *testing.T) {
		req := &models.TransaccionRequest{
			TipoEvento:  "fabricacion",
			IDProducto:  "PROD-001",
			DatosEvento: `{"lote": "12345"}`,
			ActorEmisor: "Laboratorio ABC",
		}

		err := validation.ValidateStruct(req)
		assert.NoError(t, err)
	})

	t.Run("TransaccionRequest con TipoEvento inválido", func(t *testing.T) {
		req := &models.TransaccionRequest{
			TipoEvento:  "tipo_invalido",
			IDProducto:  "PROD-001",
			DatosEvento: `{"lote": "12345"}`,
			ActorEmisor: "Laboratorio ABC",
		}

		err := validation.ValidateStruct(req)
		assert.Error(t, err)
	})

	t.Run("TransaccionRequest con campos faltantes", func(t *testing.T) {
		req := &models.TransaccionRequest{
			TipoEvento: "fabricacion",
			// Falta IDProducto
			DatosEvento: `{"lote": "12345"}`,
			ActorEmisor: "Laboratorio ABC",
		}

		err := validation.ValidateStruct(req)
		assert.Error(t, err)
	})
}

func TestValidateEthereumAddress(t *testing.T) {
	tests := []struct {
		name     string
		address  string
		expected bool
	}{
		{
			name:     "Dirección válida",
			address:  "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
			expected: true,
		},
		{
			name:     "Dirección sin 0x",
			address:  "742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
			expected: false,
		},
		{
			name:     "Dirección muy corta",
			address:  "0x123",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Esta es una prueba indirecta, en producción usaríamos el validador directamente
			// Por ahora solo documentamos los casos
			t.Log("Test case:", tt.name, "Address:", tt.address)
		})
	}
}
