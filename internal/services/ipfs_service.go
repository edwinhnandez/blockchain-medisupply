package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

// IPFSService maneja las operaciones con IPFS
type IPFSService struct {
	host       string
	port       string
	httpClient *http.Client
}

// IPFSAddResponse representa la respuesta de IPFS al agregar un archivo
type IPFSAddResponse struct {
	Name string `json:"Name"`
	Hash string `json:"Hash"`
	Size string `json:"Size"`
}

// NewIPFSService crea una nueva instancia de IPFSService
func NewIPFSService(host, port string) *IPFSService {
	return &IPFSService{
		host: host,
		port: port,
		httpClient: &http.Client{
			Timeout: 60 * time.Second, // Aumentado a 60 segundos para evitar timeouts
		},
	}
}

// GetHost retorna el host de IPFS
func (s *IPFSService) GetHost() string {
	return s.host
}

// GetPort retorna el puerto de IPFS
func (s *IPFSService) GetPort() string {
	return s.port
}

// AlmacenarJSON almacena datos JSON en IPFS y retorna el CID
func (s *IPFSService) AlmacenarJSON(ctx context.Context, data string) (string, error) {
	fmt.Println("Almacenando datos en IPFS...", s.host, s.port)
	fmt.Println("Datos a almacenar:", data)
	return s.Almacenar(ctx, []byte(data))
}

// Almacenar almacena datos en IPFS y retorna el CID
func (s *IPFSService) Almacenar(ctx context.Context, data []byte) (string, error) {
	fmt.Printf("🟡 IPFS: Iniciando almacenamiento en %s:%s\n", s.host, s.port)
	url := fmt.Sprintf("http://%s:%s/api/v0/add", s.host, s.port)
	fmt.Printf("🟡 IPFS: URL de almacenamiento: %s\n", url)
	// Crear multipart form data
	body := &bytes.Buffer{}
	fmt.Println("Body:", body)
	writer := multipart.NewWriter(body)
	fmt.Println("Writer:", writer)
	part, err := writer.CreateFormFile("file", "data.json")
	if err != nil {
		return "", fmt.Errorf("error creando form file: %w", err)
	}

	if _, err := part.Write(data); err != nil {
		return "", fmt.Errorf("error escribiendo datos: %w", err)
	}

	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("error cerrando writer: %w", err)
	}

	// Crear request
	req, err := http.NewRequestWithContext(ctx, "POST", url, body)
	if err != nil {
		return "", fmt.Errorf("error creando request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Ejecutar request
	fmt.Printf("🟡 IPFS: Enviando request (timeout: 60s)...\n")
	startTime := time.Now()
	resp, err := s.httpClient.Do(req)
	elapsed := time.Since(startTime)
	fmt.Printf("🟡 IPFS: Request completado en %v\n", elapsed)

	if err != nil {
		// Verificar si es timeout
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("timeout al conectar con IPFS después de %v: verifique que IPFS esté corriendo en %s:%s", elapsed, s.host, s.port)
		}
		return "", fmt.Errorf("error ejecutando request a IPFS después de %v: %w", elapsed, err)
	}
	fmt.Printf("🟡 IPFS: Response recibido - Status: %d\n", resp.StatusCode)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("IPFS retornó status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	// Parsear respuesta
	var result IPFSAddResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decodificando respuesta IPFS: %w", err)
	}

	if result.Hash == "" {
		return "", fmt.Errorf("no se pudo obtener CID de la respuesta IPFS")
	}

	// Pinear el CID para asegurar su persistencia
	if err := s.pinCID(ctx, result.Hash); err != nil {
		// No retornar error fatal, pero sí loguearlo. El archivo fue añadido.
		fmt.Printf("⚠️ IPFS: ADVERTENCIA: no se pudo pinear el CID %s: %v\n", result.Hash, err)
	}

	return result.Hash, nil
}

// pinCID pinea un CID en el nodo IPFS para asegurar su persistencia
func (s *IPFSService) pinCID(ctx context.Context, cid string) error {
	fmt.Printf("🟡 IPFS: Pineando CID %s...\n", cid)
	url := fmt.Sprintf("http://%s:%s/api/v0/pin/add?arg=%s", s.host, s.port, cid)

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return fmt.Errorf("error creando request de pineo: %w", err)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error ejecutando request de pineo: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("IPFS retornó status %d en pineo: %s", resp.StatusCode, string(bodyBytes))
	}

	fmt.Printf("✅ IPFS: CID %s pineado correctamente\n", cid)
	return nil
}

// RecuperarJSON recupera datos JSON de IPFS usando el CID
func (s *IPFSService) RecuperarJSON(ctx context.Context, cid string) (string, error) {
	data, err := s.Recuperar(ctx, cid)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Recuperar recupera datos de IPFS usando el CID
func (s *IPFSService) Recuperar(ctx context.Context, cid string) ([]byte, error) {
	url := fmt.Sprintf("http://%s:%s/api/v0/cat?arg=%s", s.host, s.port, cid)

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creando request: %w", err)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error recuperando de IPFS: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("IPFS retornó status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error leyendo respuesta IPFS: %w", err)
	}

	return content, nil
}

// VerificarConexion verifica si el nodo IPFS está disponible
func (s *IPFSService) VerificarConexion(ctx context.Context) error {
	url := fmt.Sprintf("http://%s:%s/api/v0/version", s.host, s.port)

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return fmt.Errorf("error creando request: %w", err)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error conectando a IPFS: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("IPFS no está disponible, status: %d", resp.StatusCode)
	}

	return nil
}
