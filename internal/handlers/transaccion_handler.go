package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/edinfamous/blockchain-medisupply/internal/models"
	"github.com/edinfamous/blockchain-medisupply/internal/services"
)

// TransaccionHandler maneja las peticiones HTTP relacionadas con transacciones
type TransaccionHandler struct {
	transaccionService *services.TransaccionService
}

// NewTransaccionHandler crea una nueva instancia de TransaccionHandler
func NewTransaccionHandler(transaccionService *services.TransaccionService) *TransaccionHandler {
	return &TransaccionHandler{
		transaccionService: transaccionService,
	}
}

// RegistrarTransaccion maneja POST /transaccion/registrar
func (h *TransaccionHandler) RegistrarTransaccion(c *gin.Context) {
	fmt.Println("🔵 Handler: RegistrarTransaccion - INICIADO")
	fmt.Printf("🔵 Handler: Método HTTP: %s\n", c.Request.Method)
	fmt.Printf("🔵 Handler: Path: %s\n", c.Request.URL.Path)
	
	var req models.TransaccionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("🔴 Handler: Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos inválidos",
			"details": err.Error(),
		})
		return
	}

	fmt.Printf("🔵 Handler: JSON binded correctamente. TipoEvento: %s, IDProducto: %s\n", req.TipoEvento, req.IDProducto)
	fmt.Println("🔵 Handler: Llamando a transaccionService.RegistrarTransaccion...")
	
	// Crear contexto con timeout más largo (90 segundos para IPFS + DynamoDB)
	ctx, cancel := context.WithTimeout(c.Request.Context(), 90*time.Second)
	defer cancel()
	
	transaccion, err := h.transaccionService.RegistrarTransaccion(ctx, &req)
	if err != nil {
		fmt.Printf("🔴 Handler: Error del servicio: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error registrando transacción",
			"details": err.Error(),
		})
		return
	}
	fmt.Printf("🟢 Handler: Transacción registrada exitosamente. ID: %s\n", transaccion.IDTransaction)

	// Convertir a response
	response := &models.TransaccionResponse{
		IDTransaction:       transaccion.IDTransaction,
		TipoEvento:          transaccion.TipoEvento,
		IDProducto:          transaccion.IDProducto,
		FechaEvento:         transaccion.FechaEvento,
		HashEvento:          transaccion.HashEvento,
		DirectionBlockchain: transaccion.DirectionBlockchain,
		IPFSCid:             transaccion.IPFSCid,
		ActorEmisor:         transaccion.ActorEmisor,
		Estado:              transaccion.Estado,
		CreatedAt:           transaccion.CreatedAt,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Transacción registrada exitosamente",
		"data":    response,
	})
}

// ObtenerTransaccion maneja GET /transaccion/:id
func (h *TransaccionHandler) ObtenerTransaccion(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de transacción requerido",
		})
		return
	}

	transaccion, err := h.transaccionService.ObtenerTransaccion(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Transacción no encontrada",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transaccion,
	})
}

// VerificarTransaccion maneja GET /transaccion/verificar/:id
func (h *TransaccionHandler) VerificarTransaccion(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de transacción requerido",
		})
		return
	}

	verificacion, err := h.transaccionService.VerificarIntegridad(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error verificando transacción",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": verificacion,
	})
}

// ListarTransacciones maneja GET /transaccion
func (h *TransaccionHandler) ListarTransacciones(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "50")
	limit, err := strconv.ParseInt(limitStr, 10, 32)
	if err != nil {
		limit = 50
	}

	transacciones, err := h.transaccionService.ListarTransacciones(c.Request.Context(), int32(limit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error listando transacciones",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": len(transacciones),
		"data":  transacciones,
	})
}

// ObtenerTransaccionesPorProducto maneja GET /transaccion/producto/:id
func (h *TransaccionHandler) ObtenerTransaccionesPorProducto(c *gin.Context) {
	idProducto := c.Param("id")

	if idProducto == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de producto requerido",
		})
		return
	}

	transacciones, err := h.transaccionService.ObtenerTransaccionesPorProducto(c.Request.Context(), idProducto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error obteniendo transacciones del producto",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"idProducto": idProducto,
		"total":      len(transacciones),
		"data":       transacciones,
	})
}

// ObtenerEstadoBlockchain maneja GET /transaccion/estado-blockchain/:id
// Permite verificar si una transacción fue registrada exitosamente en blockchain
func (h *TransaccionHandler) ObtenerEstadoBlockchain(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de transacción requerido",
		})
		return
	}

	fmt.Printf("🔵 Handler: Consultando estado de blockchain para transacción %s\n", id)
	
	estado, err := h.transaccionService.ObtenerEstadoBlockchain(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Transacción no encontrada",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": estado,
	})
}
