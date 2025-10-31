package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// AESEncryption maneja encriptación AES-256-GCM
type AESEncryption struct {
	key []byte
}

// NewAESEncryption crea una nueva instancia de AESEncryption
func NewAESEncryption(key string) (*AESEncryption, error) {
	if len(key) != 32 {
		return nil, fmt.Errorf("la clave debe tener 32 bytes para AES-256")
	}
	return &AESEncryption{
		key: []byte(key),
	}, nil
}

// Encrypt encripta datos usando AES-256-GCM
func (a *AESEncryption) Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return "", fmt.Errorf("error creando cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("error creando GCM: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("error generando nonce: %w", err)
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt desencripta datos usando AES-256-GCM
func (a *AESEncryption) Decrypt(ciphertext string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("error decodificando base64: %w", err)
	}

	block, err := aes.NewCipher(a.key)
	if err != nil {
		return "", fmt.Errorf("error creando cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("error creando GCM: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("ciphertext demasiado corto")
	}

	nonce, cipherData := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return "", fmt.Errorf("error desencriptando: %w", err)
	}

	return string(plaintext), nil
}
