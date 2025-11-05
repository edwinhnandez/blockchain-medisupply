package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/edinfamous/blockchain-medisupply/internal/services"
)

// IPFSHandler maneja las peticiones relacionadas con IPFS
type IPFSHandler struct {
	ipfsService *services.IPFSService
}

// NewIPFSHandler crea una nueva instancia de IPFSHandler
func NewIPFSHandler(ipfsService *services.IPFSService) *IPFSHandler {
	return &IPFSHandler{
		ipfsService: ipfsService,
	}
}

// IPFSFile representa un archivo en IPFS
type IPFSFile struct {
	CID       string    `json:"cid"`
	Size      string    `json:"size"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Pinned    bool      `json:"pinned"`
	GatewayURL string   `json:"gateway_url"`
	LocalURL  string    `json:"local_url"`
}

// IPFSStats representa estadísticas del nodo IPFS
type IPFSStats struct {
	RepoSize   string `json:"repo_size"`
	StorageMax string `json:"storage_max"`
	NumObjects string `json:"num_objects"`
	Version    string `json:"version"`
}

// ListarArchivos lista todos los archivos pinneados en IPFS
func (h *IPFSHandler) ListarArchivos(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	// Verificar conexión
	if err := h.ipfsService.VerificarConexion(ctx); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error":   "IPFS no está disponible",
			"details": err.Error(),
		})
		return
	}

	// Obtener archivos pinneados usando la API HTTP de IPFS
	// Nota: Necesitamos hacer una llamada directa a la API de IPFS
	// porque el servicio actual no tiene un método para listar todos los archivos
	
	// Por ahora, retornamos información sobre cómo acceder
	c.JSON(http.StatusOK, gin.H{
		"message": "Usa IPFS Companion o la API directa para ver archivos",
		"ipfs_api": "http://localhost:5001/api/v0/pin/ls",
		"ipfs_webui": "http://localhost:5001/webui",
		"gateway": "http://localhost:8081/ipfs/",
		"instructions": gin.H{
			"method1": "Instalar IPFS Companion extension",
			"method2": "Usar curl: curl -X POST http://localhost:5001/api/v0/pin/ls",
			"method3": "Usar script: ./scripts/verify_ipfs.sh",
		},
	})
}

// ObtenerArchivo obtiene información de un archivo específico por CID
func (h *IPFSHandler) ObtenerArchivo(c *gin.Context) {
	cid := c.Param("cid")
	if cid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "CID es requerido",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	// Verificar conexión
	if err := h.ipfsService.VerificarConexion(ctx); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error":   "IPFS no está disponible",
			"details": err.Error(),
		})
		return
	}

	// Intentar recuperar los datos
	data, err := h.ipfsService.RecuperarJSON(ctx, cid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "No se pudo recuperar el archivo",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cid":        cid,
		"data":       data,
		"gateway_url": "http://localhost:8081/ipfs/" + cid,
		"api_url":    "http://localhost:5001/api/v0/cat?arg=" + cid,
	})
}

// ObtenerEstadisticas obtiene estadísticas del nodo IPFS
func (h *IPFSHandler) ObtenerEstadisticas(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	// Verificar conexión
	if err := h.ipfsService.VerificarConexion(ctx); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error":   "IPFS no está disponible",
			"details": err.Error(),
		})
		return
	}

	// Retornar información de acceso
	c.JSON(http.StatusOK, gin.H{
		"status": "IPFS está disponible",
		"endpoints": gin.H{
			"api":     "http://localhost:5001/api/v0",
			"gateway": "http://localhost:8081/ipfs/",
			"webui":   "http://localhost:5001/webui",
		},
		"instructions": "Para ver archivos, usa IPFS Companion o consulta la API directamente",
	})
}


