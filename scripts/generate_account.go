package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Generar nueva clave privada
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// Obtener información de la cuenta
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Mostrar información
	fmt.Println("=== NUEVA CUENTA ETHEREUM ===")
	fmt.Println()
	fmt.Printf("Private Key: 0x%x\n", privateKeyBytes)
	fmt.Printf("Public Key:  0x%x\n", publicKeyBytes)
	fmt.Printf("Address:     %s\n", address.Hex())
	fmt.Println()
	fmt.Println("  IMPORTANTE:")
	fmt.Println("   - Guarda esta clave privada en un lugar SEGURO")
	fmt.Println("   - NUNCA compartas tu clave privada")
	fmt.Println("   - Usa esta clave solo para desarrollo/testnet")
	fmt.Println()

	// Opcional: guardar en keystore
	saveKeystore := os.Getenv("SAVE_KEYSTORE")
	if saveKeystore == "true" {
		ksDir := "./keystore"
		if err := os.MkdirAll(ksDir, 0700); err != nil {
			log.Fatal(err)
		}

		ks := keystore.NewKeyStore(ksDir, keystore.StandardScryptN, keystore.StandardScryptP)
		password := os.Getenv("KEYSTORE_PASSWORD")
		if password == "" {
			password = "password123" // Default para desarrollo
		}

		account, err := ks.ImportECDSA(privateKey, password)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("✅ Keystore guardado en: %s/%s\n", ksDir, account.Address.Hex())
		fmt.Println()
	}

	// Mostrar para copiar en .env
	fmt.Println("=== PARA TU .env ===")
	fmt.Println()
	fmt.Printf("BLOCKCHAIN_PRIVATE_KEY=0x%x\n", privateKeyBytes)
	fmt.Println()
}
