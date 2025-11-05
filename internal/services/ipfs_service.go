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
			Timeout: 30 * time.Second,
		},
	}
}

// AlmacenarJSON almacena datos JSON en IPFS y retorna el CID
func (s *IPFSService) AlmacenarJSON(ctx context.Context, data string) (string, error) {
	fmt.Println("Almacenando datos en IPFS...", s.host, s.port)
	fmt.Println("Datos a almacenar:", data)
	return s.Almacenar(ctx, []byte(data))
}

// Almacenar almacena datos en IPFS y retorna el CID
func (s *IPFSService) Almacenar(ctx context.Context, data []byte) (string, error) {
	fmt.Println("Almacenando datos en IPFS...", s.host, s.port)
	url := fmt.Sprintf("http://%s:%s/api/v0/add", s.host, s.port)
	fmt.Println("URL de almacenamiento en IPFS:", url)
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
	resp, err := s.httpClient.Do(req)
	fmt.Println("Response:", resp)
	if err != nil {
		return "", fmt.Errorf("error ejecutando request a IPFS: %w", err)
	}
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

	return result.Hash, nil
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
