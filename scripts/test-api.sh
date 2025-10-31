#!/bin/bash

# Script para probar la API

API_URL="http://localhost:8080"

echo "🧪 Probando API de Transacción Blockchain"
echo "========================================"

# Health check
echo ""
echo "1️⃣ Health Check..."
curl -s "$API_URL/health" | jq '.'

# Readiness check
echo ""
echo "2️⃣ Readiness Check..."
curl -s "$API_URL/ready" | jq '.'

# Registrar transacción de fabricación
echo ""
echo "3️⃣ Registrando transacción de fabricación..."
RESPONSE=$(curl -s -X POST "$API_URL/api/v1/transaccion/registrar" \
  -H "Content-Type: application/json" \
  -d '{
    "tipoEvento": "fabricacion",
    "idProducto": "PROD-TEST-001",
    "datosEvento": "{\"lote\": \"LOT-12345\", \"fecha_fabricacion\": \"2024-01-15\", \"cantidad\": 1000, \"planta\": \"Planta A\"}",
    "actorEmisor": "Laboratorio Medisupply SA"
  }')

echo "$RESPONSE" | jq '.'

# Extraer ID de transacción
TX_ID=$(echo "$RESPONSE" | jq -r '.data.idTransaction')

if [ "$TX_ID" != "null" ] && [ -n "$TX_ID" ]; then
  echo ""
  echo "✅ Transacción creada con ID: $TX_ID"
  
  # Esperar un momento
  sleep 2
  
  # Obtener transacción
  echo ""
  echo "4️⃣ Obteniendo transacción $TX_ID..."
  curl -s "$API_URL/api/v1/transaccion/$TX_ID" | jq '.'
  
  # Verificar integridad (puede fallar si blockchain no está configurado)
  echo ""
  echo "5️⃣ Verificando integridad de transacción..."
  curl -s "$API_URL/api/v1/transaccion/verificar/$TX_ID" | jq '.'
  
  # Registrar más transacciones para el mismo producto
  echo ""
  echo "6️⃣ Registrando transacción de distribución..."
  curl -s -X POST "$API_URL/api/v1/transaccion/registrar" \
    -H "Content-Type: application/json" \
    -d '{
      "tipoEvento": "distribucion",
      "idProducto": "PROD-TEST-001",
      "datosEvento": "{\"transportista\": \"LogiMed Express\", \"destino\": \"Farmacia Central\", \"temperatura\": \"2-8°C\"}",
      "actorEmisor": "Distribuidora Nacional"
    }' | jq '.'
  
  sleep 1
  
  # Consultar Oracle
  echo ""
  echo "7️⃣ Consultando Oracle para producto PROD-TEST-001..."
  curl -s "$API_URL/api/v1/oracle/datos/PROD-TEST-001" | jq '.'
  
  # Historial verificado
  echo ""
  echo "8️⃣ Obteniendo historial verificado..."
  curl -s "$API_URL/api/v1/oracle/historial/PROD-TEST-001" | jq '.'
  
  # Validar cadena de suministro
  echo ""
  echo "9️⃣ Validando cadena de suministro..."
  curl -s "$API_URL/api/v1/oracle/validar/PROD-TEST-001" | jq '.'
  
else
  echo "❌ Error: No se pudo crear la transacción"
fi

echo ""
echo "🎉 Pruebas completadas"

