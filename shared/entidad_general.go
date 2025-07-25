// File: entidad_general.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Declaración de estructuras generales del servicio

// Cualquier cambio a la estructura en este archivo, debe ser previa autorizacion del area de arquitectura

package shared

type MensajeDto[T any] struct {
	Objeto T             `json:"objeto,omitempty"`
	Estado EstadoHttpDto `json:"estado,omitempty"`
}

type EstadoHttpDto struct {
	Codigo  int    `json:"codigo,omitempty"`
	Mensaje string `json:"mensaje,omitempty"`
	Detalle string `json:"detalle,omitempty"`
}

type Configuracion struct {
	UrlServicioBus      string `json:"UrlServicioBus"`
	SbTopico            string `json:"SbTopico"`
	ApiNombre           string `json:"ApiNombre"`
	ApiVersion          string `json:"ApiVersion"`
	ApiSwagger          string `json:"ApiSwagger"`
	ApiDescripcion      string `json:"ApiDescripcion"`
	BdSql               string `json:"BdSql"`
	BdSqlMaxCnxAbierta  int    `json:"BdSqlMaxCnxAbierta"`
	BdSqlMaxCnxInactiva int    `json:"BdSqlMaxCnxInactiva"`
	BdSqlMaxTiempoVida  int    `json:"BdSqlMaxTiempoVida"`
	BdNoSql             string `json:"BdNoSql"`
	PuertoApi           string `json:"PuertoApi"`
	UrlApiLogger        string `json:"UrlApiLogger"`
}
