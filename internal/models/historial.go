package models

import "time"

// HistorialVerificado representa el historial completo verificado de un producto
type HistorialVerificado struct {
	IDProducto    string             `json:"idProducto"`
	TotalEventos  int                `json:"totalEventos"`
	Verificados   int                `json:"verificados"`
	NoVerificados int                `json:"noVerificados"`
	Eventos       []EventoVerificado `json:"eventos"`
	FechaConsulta time.Time          `json:"fechaConsulta"`
}

// EventoVerificado representa un evento individual verificado
type EventoVerificado struct {
	IDEvento              string    `json:"idEvento"`
	TipoEvento            string    `json:"tipoEvento"`
	Fecha                 time.Time `json:"fecha"`
	ResultadoVerificacion bool      `json:"resultadoVerificacion"`
	ReferenciaBlockchain  string    `json:"referenciaBlockchain"`
	IPFSCid               string    `json:"ipfsCid"`
	ActorEmisor           string    `json:"actorEmisor"`
	ErrorVerificacion     string    `json:"errorVerificacion,omitempty"`
}

// OracleDataResponse representa los datos expuestos por el Oracle
type OracleDataResponse struct {
	IDProducto          string             `json:"idProducto"`
	Estado              string             `json:"estado"`
	UltimaActualizacion time.Time          `json:"ultimaActualizacion"`
	CadenaVerificada    bool               `json:"cadenaVerificada"`
	Historial           []EventoVerificado `json:"historial"`
	Metadata            map[string]string  `json:"metadata"`
}
