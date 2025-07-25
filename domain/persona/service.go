// File: service.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Reglas de negocio

package persona

import (
	"github.com/shopspring/decimal"
)

const REGLA_EDAD_SUPERIOR_O_IGUAL = 18
const REGLA_IMPUESTO_SUPERIOR_O_IGUAL = 10

// Este valor deja en duro, solo para efectos de prueba, pero deben salir de una servicio de parametros
const FACTOR_MONTO = 20

func (objPersona *Persona) EdadMaxima() bool {

	// La regla esta permitda para personas mayor a 18 años
	return objPersona.Edad >= REGLA_EDAD_SUPERIOR_O_IGUAL // true = cumple la regla, False = no la cumple
}

func (objPersona *Persona) MontoImpuestoMinimo() bool {

	// El impuesto debe ser mayor o igual a 18
	return objPersona.Impuesto.GreaterThanOrEqual(decimal.NewFromInt(REGLA_IMPUESTO_SUPERIOR_O_IGUAL)) // true = cumple la regla, False = no la cumple
}

func (objPersona *Persona) CalcularMontoFuturo() decimal.Decimal {

	factorImpuesto := decimal.NewFromInt(FACTOR_MONTO)

	montoImpFinal := objPersona.Impuesto.Mul(factorImpuesto)

	return montoImpFinal
}
