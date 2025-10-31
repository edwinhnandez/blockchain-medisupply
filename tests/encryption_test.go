package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/edinfamous/blockchain-medisupply/pkg/encryption"
)

func TestAESEncryption(t *testing.T) {
	key := "12345678901234567890123456789012" // 32 bytes

	aes, err := encryption.NewAESEncryption(key)
	require.NoError(t, err)

	t.Run("Encriptar y desencriptar texto", func(t *testing.T) {
		plaintext := "Datos sensibles de paciente"

		// Encriptar
		encrypted, err := aes.Encrypt(plaintext)
		require.NoError(t, err)
		require.NotEmpty(t, encrypted)
		assert.NotEqual(t, plaintext, encrypted)

		// Desencriptar
		decrypted, err := aes.Decrypt(encrypted)
		require.NoError(t, err)
		assert.Equal(t, plaintext, decrypted)
	})

	t.Run("Encriptar JSON", func(t *testing.T) {
		jsonData := `{"paciente": "Juan Perez", "diagnostico": "confidencial"}`

		encrypted, err := aes.Encrypt(jsonData)
		require.NoError(t, err)

		decrypted, err := aes.Decrypt(encrypted)
		require.NoError(t, err)
		assert.Equal(t, jsonData, decrypted)
	})

	t.Run("Clave inválida", func(t *testing.T) {
		_, err := encryption.NewAESEncryption("clave_corta")
		assert.Error(t, err)
	})

	t.Run("Desencriptar datos corruptos", func(t *testing.T) {
		_, err := aes.Decrypt("datos_corruptos_no_base64")
		assert.Error(t, err)
	})
}
