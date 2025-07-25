// File: dto.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Estructuras de transferencias en respuesta a los consumidores en los handler

package dto

import "github.com/shopspring/decimal"

// PersonaDto
type PersonaDto struct {
	Id         int             `json:"id"`
	Nombre     string          `json:"nombre"`
	ApPaterno  string          `json:"apPaterno"`
	ApMaterno  string          `json:"apMaterno"`
	Mail       string          `json:"mail"`
	NumCelular int             `json:"numCelular"`
	Edad       int             `json:"edad"`
	Renta      int             `json:"renta"`
	Impuesto   decimal.Decimal `json:"impuesto"`
}
