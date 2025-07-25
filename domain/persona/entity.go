// File: mapper.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Entidades del dominio del negocio

package persona

import (
	"time"

	"github.com/shopspring/decimal"
)

type Persona struct {
	Id            int
	Nombre        string
	ApPaterno     string
	ApMaterno     string
	Mail          string
	NumCelular    int
	Edad          int
	Renta         int
	Impuesto      decimal.Decimal
	FechaRegistro time.Time
}

func NewPersona(id int,
	nombre string,
	apPaterno string,
	apMaterno string,
	mail string,
	numCelular int,
	edad int,
	renta int,
	impuesto decimal.Decimal) (*Persona, error) {

	return &Persona{
		Id:         id,
		Nombre:     nombre,
		ApPaterno:  apPaterno,
		ApMaterno:  apMaterno,
		Mail:       mail,
		NumCelular: numCelular,
		Edad:       edad,
		Renta:      renta,
		Impuesto:   impuesto,
	}, nil
}
