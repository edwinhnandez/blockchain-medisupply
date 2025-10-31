package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/edinfamous/blockchain-medisupply/internal/services"
)

// OracleHandler maneja las peticiones HTTP del Oracle (patrón Oracle)
type OracleHandler struct {
	oracleService *services.OracleService
}

// NewOracleHandler crea una nueva instancia de OracleHandler
func NewOracleHandler(oracleService *services.OracleService) *OracleHandler {
	return &OracleHandler{
		oracleService: oracleService,
	}
}

// ObtenerDatosVerificados maneja GET /oracle/datos/:id
// Expone datos verificados de un producto (patrón Oracle)
func (h *OracleHandler) ObtenerDatosVerificados(c *gin.Context) {
	idProducto := c.Param("id")

	if idProducto == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de producto requerido",
		})
		return
	}

	datos, err := h.oracleService.ObtenerDatosVerificados(c.Request.Context(), idProducto)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "No se pudieron obtener datos verificados",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Datos verificados del Oracle",
		"data":    datos,
	})
}

// ObtenerHistorialVerificado maneja GET /oracle/historial/:id
func (h *OracleHandler) ObtenerHistorialVerificado(c *gin.Context) {
	idProducto := c.Param("id")

	if idProducto == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de producto requerido",
		})
		return
	}

	historial, err := h.oracleService.ObtenerHistorialVerificado(c.Request.Context(), idProducto)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "No se pudo obtener historial",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Historial verificado",
		"data":    historial,
	})
}

// ValidarCadenaSupply maneja GET /oracle/validar/:id
func (h *OracleHandler) ValidarCadenaSupply(c *gin.Context) {
	idProducto := c.Param("id")

	if idProducto == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de producto requerido",
		})
		return
	}

	valido, errores, err := h.oracleService.ValidarCadenaSupply(c.Request.Context(), idProducto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error validando cadena de suministro",
			"details": err.Error(),
		})
		return
	}

	status := http.StatusOK
	if !valido {
		status = http.StatusUnprocessableEntity
	}

	c.JSON(status, gin.H{
		"idProducto":        idProducto,
		"cadenaValida":      valido,
		"erroresDetectados": errores,
	})
}
