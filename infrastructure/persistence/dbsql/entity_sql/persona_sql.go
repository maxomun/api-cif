// File: persona_sql.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-05-01
// Description: Estructura de persistencia, refleja los campos provenientes de una tabla o procedimiento almacenado

package entitysql

import "database/sql"

type PersonaSql struct {
	Id            sql.NullInt32
	Nombre        sql.NullString
	ApPaterno     sql.NullString
	ApMaterno     sql.NullString
	Mail          sql.NullString
	NumCelular    sql.NullInt32
	Edad          sql.NullInt16
	Renta         sql.NullInt32
	Impuesto      sql.NullFloat64
	FechaRegistro sql.NullTime
}
