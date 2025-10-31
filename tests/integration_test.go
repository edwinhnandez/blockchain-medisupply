package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/edinfamous/blockchain-medisupply/internal/handlers"
	"github.com/edinfamous/blockchain-medisupply/internal/models"
	"github.com/edinfamous/blockchain-medisupply/internal/services"
)

// setupTestRouter configura un router para testing con servicios mock
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Crear servicios mock
	_ = services.NewIPFSService("localhost", "5001") // TODO: Use actual service when implementing full integration tests

	// Para tests, estos servicios deberían ser mocks
	// En una implementación real, usaríamos interfaces y mocks completos

	transaccionHandler := handlers.NewTransaccionHandler(nil)

	router.POST("/transaccion/registrar", transaccionHandler.RegistrarTransaccion)

	return router
}

func TestHealthEndpoint(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	router := gin.New()

	ipfsService := services.NewIPFSService("localhost", "5001")
	healthHandler := handlers.NewHealthHandler(ipfsService, nil)

	router.GET("/health", healthHandler.HealthCheck)

	t.Run("Health check devuelve 200", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/health", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Equal(t, "ok", response["status"])
	})
}

func TestTransaccionEndpoints(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	t.Run("Registrar transacción con datos inválidos", func(t *testing.T) {
		router := setupTestRouter()

		reqBody := models.TransaccionRequest{
			TipoEvento: "tipo_invalido", // Tipo inválido
			IDProducto: "PROD-001",
		}

		jsonData, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/transaccion/registrar", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Debería fallar la validación
		assert.NotEqual(t, http.StatusCreated, w.Code)
	})
}

func TestOracleEndpoints(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Mock del oracle handler
	oracleHandler := handlers.NewOracleHandler(nil)
	router.GET("/oracle/datos/:id", oracleHandler.ObtenerDatosVerificados)

	t.Run("Obtener datos con ID vacío", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/oracle/datos/", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		// Debería retornar not found
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
