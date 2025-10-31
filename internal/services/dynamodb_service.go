package services

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/edinfamous/blockchain-medisupply/internal/models"
)

// DynamoDBService maneja las operaciones con DynamoDB
type DynamoDBService struct {
	client    *dynamodb.Client
	tableName string
}

// NewDynamoDBService crea una nueva instancia de DynamoDBService
func NewDynamoDBService(client *dynamodb.Client, tableName string) *DynamoDBService {
	return &DynamoDBService{
		client:    client,
		tableName: tableName,
	}
}

// GuardarTransaccion guarda una transacción en DynamoDB
func (s *DynamoDBService) GuardarTransaccion(ctx context.Context, transaccion *models.Transaccion) error {
	// Establecer timestamps
	now := time.Now()
	transaccion.CreatedAt = now
	transaccion.UpdatedAt = now

	// Convertir a attributevalue
	item, err := attributevalue.MarshalMap(transaccion)
	if err != nil {
		return fmt.Errorf("error marshaling transacción: %w", err)
	}

	// Guardar en DynamoDB
	_, err = s.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(s.tableName),
		Item:      item,
	})
	if err != nil {
		return fmt.Errorf("error guardando en DynamoDB: %w", err)
	}

	return nil
}

// ObtenerTransaccion obtiene una transacción por ID
func (s *DynamoDBService) ObtenerTransaccion(ctx context.Context, idTransaccion string) (*models.Transaccion, error) {
	result, err := s.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(s.tableName),
		Key: map[string]types.AttributeValue{
			"idTransaction": &types.AttributeValueMemberS{Value: idTransaccion},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error obteniendo de DynamoDB: %w", err)
	}

	if result.Item == nil {
		return nil, fmt.Errorf("transacción no encontrada")
	}

	var transaccion models.Transaccion
	err = attributevalue.UnmarshalMap(result.Item, &transaccion)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling transacción: %w", err)
	}

	return &transaccion, nil
}

// ObtenerTransaccionesPorProducto obtiene todas las transacciones de un producto
func (s *DynamoDBService) ObtenerTransaccionesPorProducto(ctx context.Context, idProducto string) ([]*models.Transaccion, error) {
	// Usar GSI (Global Secondary Index) si está configurado, o hacer scan
	input := &dynamodb.ScanInput{
		TableName:        aws.String(s.tableName),
		FilterExpression: aws.String("idProducto = :idProducto"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":idProducto": &types.AttributeValueMemberS{Value: idProducto},
		},
	}

	result, err := s.client.Scan(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("error scanning DynamoDB: %w", err)
	}

	var transacciones []*models.Transaccion
	for _, item := range result.Items {
		var transaccion models.Transaccion
		if err := attributevalue.UnmarshalMap(item, &transaccion); err != nil {
			continue // Skip items que no se pueden unmarshal
		}
		transacciones = append(transacciones, &transaccion)
	}

	return transacciones, nil
}

// ActualizarHashBlockchain actualiza el hash de blockchain de una transacción
func (s *DynamoDBService) ActualizarHashBlockchain(ctx context.Context, idTransaccion, hashBlockchain string) error {
	_, err := s.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(s.tableName),
		Key: map[string]types.AttributeValue{
			"idTransaction": &types.AttributeValueMemberS{Value: idTransaccion},
		},
		UpdateExpression: aws.String("SET directionBlockchain = :hash, updatedAt = :updatedAt, estado = :estado"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":hash":      &types.AttributeValueMemberS{Value: hashBlockchain},
			":updatedAt": &types.AttributeValueMemberS{Value: time.Now().Format(time.RFC3339)},
			":estado":    &types.AttributeValueMemberS{Value: "confirmado"},
		},
	})
	if err != nil {
		return fmt.Errorf("error actualizando hash blockchain: %w", err)
	}

	return nil
}

// ActualizarEstado actualiza el estado de una transacción
func (s *DynamoDBService) ActualizarEstado(ctx context.Context, idTransaccion, estado string) error {
	_, err := s.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(s.tableName),
		Key: map[string]types.AttributeValue{
			"idTransaction": &types.AttributeValueMemberS{Value: idTransaccion},
		},
		UpdateExpression: aws.String("SET estado = :estado, updatedAt = :updatedAt"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":estado":    &types.AttributeValueMemberS{Value: estado},
			":updatedAt": &types.AttributeValueMemberS{Value: time.Now().Format(time.RFC3339)},
		},
	})
	if err != nil {
		return fmt.Errorf("error actualizando estado: %w", err)
	}

	return nil
}

// ListarTransacciones lista todas las transacciones (con paginación opcional)
func (s *DynamoDBService) ListarTransacciones(ctx context.Context, limit int32) ([]*models.Transaccion, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(s.tableName),
	}

	if limit > 0 {
		input.Limit = aws.Int32(limit)
	}

	result, err := s.client.Scan(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("error listing transacciones: %w", err)
	}

	var transacciones []*models.Transaccion
	for _, item := range result.Items {
		var transaccion models.Transaccion
		if err := attributevalue.UnmarshalMap(item, &transaccion); err != nil {
			continue
		}
		transacciones = append(transacciones, &transaccion)
	}

	return transacciones, nil
}
