package models

import "time"

// Transaccion representa una transacción en el sistema blockchain
type Transaccion struct {
	IDTransaction       string    `json:"idTransaction" dynamodbav:"idTransaction" validate:"required"`
	TipoEvento          string    `json:"tipoEvento" dynamodbav:"tipoEvento" validate:"required,oneof=fabricacion distribucion recepcion verificacion"`
	IDProducto          string    `json:"idProducto" dynamodbav:"idProducto" validate:"required"`
	FechaEvento         time.Time `json:"fechaEvento" dynamodbav:"fechaEvento" validate:"required"`
	DatosEvento         string    `json:"datosEvento" dynamodbav:"datosEvento" validate:"required"` // JSON string con datos completos
	HashEvento          string    `json:"hashEvento" dynamodbav:"hashEvento"`
	DirectionBlockchain string    `json:"directionBlockchain" dynamodbav:"directionBlockchain"`
	IPFSCid             string    `json:"ipfsCid" dynamodbav:"ipfsCid"` // CID de IPFS para off-chain storage
	ActorEmisor         string    `json:"actorEmisor" dynamodbav:"actorEmisor" validate:"required"`
	Estado              string    `json:"estado" dynamodbav:"estado" validate:"required,oneof=pendiente confirmado fallido"`
	FirmaDigital        string    `json:"firmaDigital" dynamodbav:"firmaDigital"`
	CreatedAt           time.Time `json:"createdAt" dynamodbav:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt" dynamodbav:"updatedAt"`
}

// TransaccionRequest representa el payload de creación de transacción
type TransaccionRequest struct {
	TipoEvento  string `json:"tipoEvento" validate:"required,oneof=fabricacion distribucion recepcion verificacion"`
	IDProducto  string `json:"idProducto" validate:"required"`
	DatosEvento string `json:"datosEvento" validate:"required"`
	ActorEmisor string `json:"actorEmisor" validate:"required"`
}

// TransaccionResponse representa la respuesta de una transacción
type TransaccionResponse struct {
	IDTransaction       string    `json:"idTransaction"`
	TipoEvento          string    `json:"tipoEvento"`
	IDProducto          string    `json:"idProducto"`
	FechaEvento         time.Time `json:"fechaEvento"`
	HashEvento          string    `json:"hashEvento"`
	DirectionBlockchain string    `json:"directionBlockchain"`
	IPFSCid             string    `json:"ipfsCid"`
	ActorEmisor         string    `json:"actorEmisor"`
	Estado              string    `json:"estado"`
	CreatedAt           time.Time `json:"createdAt"`
}

// VerificacionResponse representa el resultado de verificación de integridad
type VerificacionResponse struct {
	IDTransaction        string `json:"idTransaction"`
	Verificado           bool   `json:"verificado"`
	HashLocal            string `json:"hashLocal"`
	HashBlockchain       string `json:"hashBlockchain"`
	DatosIPFSVerificados bool   `json:"datosIPFSVerificados"`
	Mensaje              string `json:"mensaje"`
}
