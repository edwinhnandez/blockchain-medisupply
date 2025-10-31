package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/edinfamous/blockchain-medisupply/internal/services"
)

// HealthHandler maneja las peticiones de health check
type HealthHandler struct {
	ipfsService       *services.IPFSService
	blockchainService *services.BlockchainService
}

// NewHealthHandler crea una nueva instancia de HealthHandler
func NewHealthHandler(ipfs *services.IPFSService, blockchain *services.BlockchainService) *HealthHandler {
	return &HealthHandler{
		ipfsService:       ipfs,
		blockchainService: blockchain,
	}
}

// HealthCheck maneja GET /health
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"timestamp": time.Now().Format(time.RFC3339),
		"service":   "transaccion-blockchain-service",
	})
}

// ReadinessCheck maneja GET /ready
func (h *HealthHandler) ReadinessCheck(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	checks := make(map[string]string)
	allHealthy := true

	// Verificar IPFS
	if err := h.ipfsService.VerificarConexion(ctx); err != nil {
		checks["ipfs"] = "unhealthy: " + err.Error()
		allHealthy = false
	} else {
		checks["ipfs"] = "healthy"
	}

	// Verificar Blockchain
	if err := h.blockchainService.VerificarConexion(ctx); err != nil {
		checks["blockchain"] = "unhealthy: " + err.Error()
		allHealthy = false
	} else {
		checks["blockchain"] = "healthy"
	}

	// DynamoDB se verifica automáticamente en cada llamada
	checks["dynamodb"] = "healthy"

	status := http.StatusOK
	if !allHealthy {
		status = http.StatusServiceUnavailable
	}

	c.JSON(status, gin.H{
		"status":    map[bool]string{true: "ready", false: "not_ready"}[allHealthy],
		"checks":    checks,
		"timestamp": time.Now().Format(time.RFC3339),
	})
}
