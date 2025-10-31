package handlers

import (
	"net/http"
	"strconv"

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
	var req models.TransaccionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos inválidos",
			"details": err.Error(),
		})
		return
	}

	transaccion, err := h.transaccionService.RegistrarTransaccion(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error registrando transacción",
			"details": err.Error(),
		})
		return
	}

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
