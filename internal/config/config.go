package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config representa la configuración de la aplicación
type Config struct {
	// AWS
	AWSRegion         string
	AWSAccessKeyID    string
	AWSSecretKey      string
	DynamoDBTableName string
	UseAWSSecrets     bool

	// Blockchain
	AlchemyAPIKey            string
	BlockchainRPCURL         string // Full RPC URL (optional, constructed if not provided)
	BlockchainNetwork        string
	BlockchainPrivateKeyName string
	ContractAddress          string

	// IPFS
	IPFSHost        string
	IPFSPort        string
	IPFSGatewayPort string

	// Server
	ServerPort string
	GinMode    string

	// Security
	EncryptionKey string

	// Rate Limiting
	RateLimitRequests int
	RateLimitWindow   int
}

var AppConfig *Config

// LoadConfig carga la configuración desde variables de entorno
func LoadConfig() (*Config, error) {
	// Intentar cargar .env solo en desarrollo
	_ = godotenv.Load()

	config := &Config{
		AWSRegion:                getEnv("AWS_REGION", "us-east-1"),
		AWSAccessKeyID:           getEnv("AWS_ACCESS_KEY_ID", ""),
		AWSSecretKey:             getEnv("AWS_SECRET_ACCESS_KEY", ""),
		DynamoDBTableName:        getEnv("DYNAMODB_TABLE_NAME", "transacciones-blockchain"),
		UseAWSSecrets:            getEnvAsBool("USE_AWS_SECRETS", false),
		AlchemyAPIKey:            getEnv("ALCHEMY_API_KEY", ""),
		BlockchainRPCURL:         getEnv("BLOCKCHAIN_RPC_URL", ""),
		BlockchainNetwork:        getEnv("BLOCKCHAIN_NETWORK", "sepolia"),
		BlockchainPrivateKeyName: getEnv("BLOCKCHAIN_PRIVATE_KEY_SECRET", "blockchain-private-key"),
		ContractAddress:          getEnv("CONTRACT_ADDRESS", ""),
		IPFSHost:                 getEnv("IPFS_HOST", "localhost"),
		IPFSPort:                 getEnv("IPFS_PORT", "5001"),
		IPFSGatewayPort:          getEnv("IPFS_GATEWAY_PORT", "8081"),
		ServerPort:               getEnv("SERVER_PORT", "8080"),
		GinMode:                  getEnv("GIN_MODE", "debug"),
		EncryptionKey:            getEnv("ENCRYPTION_KEY", ""),
		RateLimitRequests:        getEnvAsInt("RATE_LIMIT_REQUESTS", 100),
		RateLimitWindow:          getEnvAsInt("RATE_LIMIT_WINDOW", 60),
	}

	// Validar configuración crítica
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("configuración inválida: %w", err)
	}

	AppConfig = config
	return config, nil
}

// Validate valida la configuración
func (c *Config) Validate() error {
	if c.EncryptionKey == "" {
		return fmt.Errorf("ENCRYPTION_KEY es requerida")
	}

	if len(c.EncryptionKey) != 32 {
		return fmt.Errorf("ENCRYPTION_KEY debe tener exactamente 32 caracteres para AES-256")
	}

	if c.DynamoDBTableName == "" {
		return fmt.Errorf("DYNAMODB_TABLE_NAME es requerida")
	}

	if c.IPFSHost == "" {
		return fmt.Errorf("IPFS_HOST es requerido")
	}

	return nil
}

// getEnv obtiene una variable de entorno o retorna el valor por defecto
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// getEnvAsInt obtiene una variable de entorno como entero
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

// getEnvAsBool obtiene una variable de entorno como booleano
func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}
