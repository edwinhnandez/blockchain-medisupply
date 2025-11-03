package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"

	appConfig "github.com/edinfamous/blockchain-medisupply/internal/config"
	"github.com/edinfamous/blockchain-medisupply/internal/handlers"
	"github.com/edinfamous/blockchain-medisupply/internal/middleware"
	"github.com/edinfamous/blockchain-medisupply/internal/services"
)

func main() {
	// Cargar configuración
	cfg, err := appConfig.LoadConfig()
	if err != nil {
		log.Fatalf("Error cargando configuración: %v", err)
	}

	log.Println("✅ Configuración cargada correctamente")

	// Configurar Gin
	if cfg.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Inicializar servicios
	log.Println("🔧 Inicializando servicios...")

	// 1. Inicializar IPFS Service
	ipfsService := services.NewIPFSService(cfg.IPFSHost, cfg.IPFSPort)
	if err := ipfsService.VerificarConexion(context.Background()); err != nil {
		log.Printf("⚠️  ADVERTENCIA: No se pudo conectar a IPFS: %v", err)
		log.Println("   Asegúrese de que el nodo IPFS esté corriendo")
	} else {
		log.Println("✅ Conectado a IPFS")
	}

	// 2. Inicializar DynamoDB Service
	dynamoDBClient, err := initializeDynamoDB(cfg)
	if err != nil {
		log.Fatalf("Error inicializando DynamoDB: %v", err)
	}
	dynamoDBService := services.NewDynamoDBService(dynamoDBClient, cfg.DynamoDBTableName)
	log.Println("✅ Conectado a DynamoDB")

	// 3. Inicializar Blockchain Service
	var blockchainService *services.BlockchainService

	// Construir RPC URL usando Alchemy, o usar URL personalizada si está configurada
	var rpcURL string
	if cfg.BlockchainRPCURL != "" {
		// Usar URL personalizada si está configurada
		rpcURL = cfg.BlockchainRPCURL
		log.Printf("🔗 Usando RPC URL personalizada")
	} else if cfg.AlchemyAPIKey != "" {
		// Construir URL de Alchemy basada en la red
		rpcURL = fmt.Sprintf("https://eth-%s.g.alchemy.com/v2/%s", cfg.BlockchainNetwork, cfg.AlchemyAPIKey)
		log.Printf("🔗 Usando Alchemy RPC para red %s", cfg.BlockchainNetwork)
	}

	if rpcURL != "" {
		// En producción, obtener la clave privada de AWS Secrets Manager
		// Por ahora, usar una clave de prueba o desde variable de entorno
		privateKey := os.Getenv("BLOCKCHAIN_PRIVATE_KEY")
		if privateKey == "" {
			log.Println("ADVERTENCIA: BLOCKCHAIN_PRIVATE_KEY no configurada")
			log.Println("   El servicio funcionará pero no podrá escribir en blockchain")
		} else {
			// Obtener dirección del contrato (puede estar vacía para modo sin contrato)
			contractAddress := cfg.ContractAddress
			blockchainService, err = services.NewBlockchainService(rpcURL, privateKey, contractAddress)
			if err != nil {
				log.Printf("ADVERTENCIA: Error inicializando blockchain: %v", err)
			} else {
				if contractAddress != "" {
					log.Printf("✅ Conectado a Blockchain con Smart Contract: %s", contractAddress)
				} else {
					log.Println("✅ Conectado a Blockchain (modo sin contrato)")
				}
			}
		}
	}

	// Si blockchain no está disponible, crear un servicio mock para no romper la aplicación
	if blockchainService == nil {
		log.Println("Usando modo sin blockchain (solo almacenamiento off-chain)")
	}

	// 4. Inicializar servicios de negocio
	transaccionService := services.NewTransaccionService(blockchainService, ipfsService, dynamoDBService)
	oracleService := services.NewOracleService(transaccionService, dynamoDBService)

	// 5. Inicializar handlers
	transaccionHandler := handlers.NewTransaccionHandler(transaccionService)
	oracleHandler := handlers.NewOracleHandler(oracleService)
	healthHandler := handlers.NewHealthHandler(ipfsService, blockchainService)

	// Configurar router
	router := setupRouter(cfg, transaccionHandler, oracleHandler, healthHandler)

	// Iniciar servidor con graceful shutdown
	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: router,
	}

	// Canal para señales del sistema
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Iniciar servidor en goroutine
	go func() {
		log.Printf("🚀 Servidor iniciado en puerto %s", cfg.ServerPort)
		log.Printf("📝 Documentación: http://localhost:%s/health", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error iniciando servidor: %v", err)
		}
	}()

	// Esperar señal de interrupción
	<-quit
	log.Println("🛑 Apagando servidor...")

	// Graceful shutdown con timeout de 5 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Error en shutdown: %v", err)
	}

	// Cerrar conexión blockchain si existe
	if blockchainService != nil {
		blockchainService.Close()
	}

	log.Println("✅ Servidor detenido correctamente")
}

// setupRouter configura todas las rutas de la API
func setupRouter(cfg *appConfig.Config, transaccionHandler *handlers.TransaccionHandler, oracleHandler *handlers.OracleHandler, healthHandler *handlers.HealthHandler) *gin.Engine {
	router := gin.New()

	// Middleware globales
	router.Use(gin.Recovery())
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.RateLimitMiddleware(cfg.RateLimitRequests, cfg.RateLimitWindow))

	// Health checks
	router.GET("/health", healthHandler.HealthCheck)
	router.GET("/ready", healthHandler.ReadinessCheck)

	// API v1
	v1 := router.Group("/api/v1")
	{
		// Rutas de transacciones
		transacciones := v1.Group("/transaccion")
		{
			transacciones.POST("/registrar", transaccionHandler.RegistrarTransaccion)
			transacciones.GET("/:id", transaccionHandler.ObtenerTransaccion)
			transacciones.GET("/verificar/:id", transaccionHandler.VerificarTransaccion)
			transacciones.GET("", transaccionHandler.ListarTransacciones)
			transacciones.GET("/producto/:id", transaccionHandler.ObtenerTransaccionesPorProducto)
		}

		// Rutas del Oracle (patrón Oracle)
		oracle := v1.Group("/oracle")
		{
			oracle.GET("/datos/:id", oracleHandler.ObtenerDatosVerificados)
			oracle.GET("/historial/:id", oracleHandler.ObtenerHistorialVerificado)
			oracle.GET("/validar/:id", oracleHandler.ValidarCadenaSupply)
		}
	}

	// Ruta raíz con información de la API
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"service":     "Transacción Blockchain MediSupply",
			"version":     "1.0.0",
			"description": "Microservicio con patrones Oracle, Off-chain Storage e IPFS",
			"endpoints": gin.H{
				"health":        "/health",
				"ready":         "/ready",
				"api":           "/api/v1",
				"transacciones": "/api/v1/transaccion",
				"oracle":        "/api/v1/oracle",
			},
		})
	})

	return router
}

// initializeDynamoDB inicializa el cliente de DynamoDB
func initializeDynamoDB(cfg *appConfig.Config) (*dynamodb.Client, error) {
	ctx := context.Background()

	var awsCfg aws.Config
	var err error

	if cfg.AWSAccessKeyID != "" && cfg.AWSSecretKey != "" {
		// Usar credenciales explícitas
		awsCfg, err = config.LoadDefaultConfig(ctx,
			config.WithRegion(cfg.AWSRegion),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
				cfg.AWSAccessKeyID,
				cfg.AWSSecretKey,
				"",
			)),
		)
	} else {
		// Usar credenciales por defecto (IAM role, etc.)
		awsCfg, err = config.LoadDefaultConfig(ctx,
			config.WithRegion(cfg.AWSRegion),
		)
	}

	if err != nil {
		return nil, fmt.Errorf("error cargando config de AWS: %w", err)
	}

	return dynamodb.NewFromConfig(awsCfg), nil
}
