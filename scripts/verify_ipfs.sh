#!/bin/bash

# Script para verificar que IPFS guardó correctamente los datos
# Uso: ./scripts/verify_ipfs.sh [CID]

# Colores para output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuración
IPFS_HOST="${IPFS_HOST:-localhost}"
IPFS_PORT="${IPFS_PORT:-5001}"
API_URL="${API_URL:-http://localhost:8080}"

echo -e "${BLUE}=== Verificación de IPFS ===${NC}"
echo ""

# 1. Verificar que IPFS está corriendo
echo -e "${YELLOW}1. Verificando conexión a IPFS...${NC}"
VERSION=$(curl -s -X POST "http://$IPFS_HOST:$IPFS_PORT/api/v0/version" 2>/dev/null)
if [ $? -eq 0 ] && [ -n "$VERSION" ]; then
    IPFS_VER=$(echo $VERSION | jq -r '.Version' 2>/dev/null || echo "unknown")
    echo -e "${GREEN}✅ IPFS está corriendo (versión: $IPFS_VER)${NC}"
else
    echo -e "${RED}❌ IPFS no está disponible en http://$IPFS_HOST:$IPFS_PORT${NC}"
    echo "   Verifica que IPFS esté corriendo: docker-compose ps ipfs"
    exit 1
fi

# 2. Obtener estadísticas del nodo
echo ""
echo -e "${YELLOW}2. Estadísticas del nodo IPFS:${NC}"
STATS=$(curl -s -X POST "http://$IPFS_HOST:$IPFS_PORT/api/v0/stats/repo" 2>/dev/null)
if [ -n "$STATS" ]; then
    REPO_SIZE=$(echo $STATS | jq -r '.RepoSize' 2>/dev/null || echo "N/A")
    STORAGE_MAX=$(echo $STATS | jq -r '.StorageMax' 2>/dev/null || echo "N/A")
    NUM_OBJECTS=$(echo $STATS | jq -r '.NumObjects' 2>/dev/null || echo "N/A")
    echo "   Repo Size: $REPO_SIZE bytes"
    echo "   Storage Max: $STORAGE_MAX bytes"
    echo "   Num Objects: $NUM_OBJECTS"
else
    echo -e "${RED}❌ No se pudieron obtener estadísticas${NC}"
fi

# 3. Listar archivos pinneados
echo ""
echo -e "${YELLOW}3. Archivos pinneados en el nodo:${NC}"
PINNED=$(curl -s -X POST "http://$IPFS_HOST:$IPFS_PORT/api/v0/pin/ls" 2>/dev/null)
if [ -n "$PINNED" ]; then
    CIDS=$(echo $PINNED | jq -r '.Keys | keys[]' 2>/dev/null)
    if [ -n "$CIDS" ]; then
        COUNT=$(echo "$CIDS" | wc -l | tr -d ' ')
        echo -e "${GREEN}✅ Encontrados $COUNT archivo(s) pinneado(s)${NC}"
        echo "$CIDS" | while read CID; do
            echo "   - $CID"
        done
    else
        echo -e "${YELLOW}⚠️  No hay archivos pinneados${NC}"
    fi
else
    echo -e "${RED}❌ No se pudieron listar archivos pinneados${NC}"
fi

# 4. Si se proporciona un CID, verificar
if [ -n "$1" ]; then
    CID=$1
    echo ""
    echo -e "${YELLOW}4. Verificando CID específico: $CID${NC}"
    
    # Verificar formato
    if ! echo "$CID" | grep -qE '^Qm[a-zA-Z0-9]{44}$'; then
        echo -e "${RED}❌ Formato de CID inválido${NC}"
        echo "   Debe ser un CIDv0 válido (empezar con Qm y tener 46 caracteres)"
        exit 1
    fi
    
    # Verificar que está pinneado
    PIN_CHECK=$(curl -s -X POST "http://$IPFS_HOST:$IPFS_PORT/api/v0/pin/ls?arg=$CID" 2>/dev/null)
    if [ -n "$PIN_CHECK" ]; then
        echo -e "${GREEN}✅ CID está pinneado${NC}"
    else
        echo -e "${YELLOW}⚠️  CID no está pinneado localmente${NC}"
        echo "   Intentando pinnearlo..."
        PIN_RESULT=$(curl -s -X POST "http://$IPFS_HOST:$IPFS_PORT/api/v0/pin/add?arg=$CID" 2>/dev/null)
        if [ -n "$PIN_RESULT" ]; then
            echo -e "${GREEN}✅ CID pinneado exitosamente${NC}"
        fi
    fi
    
    # Recuperar datos
    echo ""
    echo -e "${YELLOW}5. Recuperando datos del CID:${NC}"
    DATA=$(curl -s -X POST "http://$IPFS_HOST:$IPFS_PORT/api/v0/cat?arg=$CID" 2>/dev/null)
    
    if [ -n "$DATA" ]; then
        echo -e "${GREEN}✅ Datos recuperados correctamente${NC}"
        echo ""
        echo "Contenido:"
        echo "$DATA" | jq . 2>/dev/null || echo "$DATA"
        echo ""
        echo "Tamaño: $(echo -n "$DATA" | wc -c) bytes"
    else
        echo -e "${RED}❌ No se pudieron recuperar los datos${NC}"
        echo "   Posibles causas:"
        echo "   - El CID no existe en este nodo"
        echo "   - El nodo no está conectado a la red IPFS"
        echo "   - Los datos fueron eliminados"
    fi
    
    # Verificar vía gateway
    echo ""
    echo -e "${YELLOW}6. Verificando acceso vía gateway:${NC}"
    GATEWAY_DATA=$(curl -s "http://$IPFS_HOST:8081/ipfs/$CID" 2>/dev/null)
    if [ -n "$GATEWAY_DATA" ]; then
        echo -e "${GREEN}✅ Accesible vía gateway local${NC}"
    else
        echo -e "${YELLOW}⚠️  No accesible vía gateway local${NC}"
    fi
fi

# 5. Verificar conexión a la red IPFS
echo ""
echo -e "${YELLOW}7. Verificando conexión a la red IPFS:${NC}"
PEERS=$(curl -s -X POST "http://$IPFS_HOST:$IPFS_PORT/api/v0/swarm/peers" 2>/dev/null)
if [ -n "$PEERS" ]; then
    PEER_COUNT=$(echo "$PEERS" | jq -r '.Strings | length' 2>/dev/null || echo "0")
    if [ "$PEER_COUNT" -gt 0 ]; then
        echo -e "${GREEN}✅ Conectado a $PEER_COUNT peer(s)${NC}"
    else
        echo -e "${YELLOW}⚠️  No conectado a peers (modo offline/local)${NC}"
    fi
else
    echo -e "${YELLOW}⚠️  No se pudo verificar conexión a peers${NC}"
fi

echo ""
echo -e "${BLUE}=== Verificación completada ===${NC}"


