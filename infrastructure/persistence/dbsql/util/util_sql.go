// File: util.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-05-01
// Description: Funcionalidades de apoyo con bases de datos relacional

package util

import (
	"database/sql"

	"github.com/shopspring/decimal"
)

func NullToInt32(valor sql.NullInt32, defaultVal int) int {
	if valor.Valid {
		return int(valor.Int32)
	}
	return defaultVal
}

func NullToInt16(valor sql.NullInt16, defaultVal int) int {
	if valor.Valid {
		return int(valor.Int16)
	}
	return defaultVal
}

func NullToString(valor sql.NullString, defaultVal string) string {
	if valor.Valid {
		return valor.String
	}
	return defaultVal
}

func NullToFloat64ToDecimal(valor sql.NullFloat64, defaultVal decimal.Decimal) decimal.Decimal {
	if valor.Valid {
		return decimal.NewFromFloat(valor.Float64)
	}
	return defaultVal
}
