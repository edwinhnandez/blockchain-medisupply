#!/bin/bash

# Script para listar y mostrar archivos IPFS de forma visual
# Uso: ./scripts/list_ipfs_files.sh

# Colores
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m'

IPFS_HOST="${IPFS_HOST:-localhost}"
IPFS_PORT="${IPFS_PORT:-5001}"

echo -e "${BLUE}=== Archivos en IPFS ===${NC}"
echo ""

# Verificar conexión
if ! curl -s -X POST "http://$IPFS_HOST:$IPFS_PORT/api/v0/version" > /dev/null; then
    echo -e "${RED}❌ IPFS no está disponible${NC}"
    exit 1
fi

# Obtener lista de archivos pinneados
echo -e "${YELLOW}Obteniendo lista de archivos...${NC}"
PINNED=$(curl -s -X POST "http://$IPFS_HOST:$IPFS_PORT/api/v0/pin/ls" 2>/dev/null)

if [ -z "$PINNED" ]; then
    echo -e "${YELLOW}⚠️  No hay archivos pinneados${NC}"
    exit 0
fi

# Extraer CIDs
CIDS=$(echo "$PINNED" | jq -r '.Keys | keys[]' 2>/dev/null)

if [ -z "$CIDS" ]; then
    echo -e "${YELLOW}⚠️  No se encontraron archivos${NC}"
    exit 0
fi

COUNT=$(echo "$CIDS" | wc -l | tr -d ' ')
echo -e "${GREEN}✅ Encontrados $COUNT archivo(s)${NC}"
echo ""

# Mostrar cada archivo
INDEX=1
echo "$CIDS" | while read CID; do
    echo -e "${BLUE}[$INDEX] CID: $CID${NC}"
    echo "   Gateway: http://localhost:8081/ipfs/$CID"
    echo "   API: http://localhost:5001/api/v0/cat?arg=$CID"
    echo ""
    INDEX=$((INDEX + 1))
done

echo -e "${GREEN}✅ Lista completa${NC}"
echo ""
echo "Para ver contenido de un archivo específico:"
echo "  curl -X POST \"http://localhost:5001/api/v0/cat?arg=<CID>\" | jq ."
echo ""
echo "O usar el endpoint de la API:"
echo "  curl http://localhost:8080/api/v1/ipfs/archivo/<CID>"
