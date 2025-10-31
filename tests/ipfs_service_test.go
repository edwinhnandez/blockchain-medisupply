package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/edinfamous/blockchain-medisupply/internal/services"
)

func TestIPFSService_AlmacenarYRecuperar(t *testing.T) {
	// Skip si IPFS no está disponible
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	ipfsService := services.NewIPFSService("localhost", "5001")

	// Verificar conexión
	err := ipfsService.VerificarConexion(context.Background())
	if err != nil {
		t.Skip("IPFS no disponible:", err)
	}

	t.Run("Almacenar y recuperar JSON", func(t *testing.T) {
		ctx := context.Background()
		testData := `{"producto": "Medicamento A", "lote": "12345", "fabricante": "Laboratorio XYZ"}`

		// Almacenar
		cid, err := ipfsService.AlmacenarJSON(ctx, testData)
		require.NoError(t, err)
		require.NotEmpty(t, cid)

		// Recuperar
		recuperado, err := ipfsService.RecuperarJSON(ctx, cid)
		require.NoError(t, err)
		assert.Equal(t, testData, recuperado)
	})

	t.Run("Recuperar CID inexistente", func(t *testing.T) {
		ctx := context.Background()
		cidInvalido := "QmInvalidCIDxxxxxxxxxxxxxxxxxxxxxxxxxxx"

		_, err := ipfsService.RecuperarJSON(ctx, cidInvalido)
		assert.Error(t, err)
	})
}

func TestIPFSService_VerificarConexion(t *testing.T) {
	t.Run("Conexión exitosa", func(t *testing.T) {
		if testing.Short() {
			t.Skip("Skipping integration test")
		}

		ipfsService := services.NewIPFSService("localhost", "5001")
		err := ipfsService.VerificarConexion(context.Background())

		// Solo verificamos si IPFS está corriendo
		if err != nil {
			t.Log("IPFS no disponible:", err)
		}
	})

	t.Run("Conexión fallida", func(t *testing.T) {
		ipfsService := services.NewIPFSService("invalid-host", "9999")
		err := ipfsService.VerificarConexion(context.Background())
		assert.Error(t, err)
	})
}
