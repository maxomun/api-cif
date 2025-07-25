// File: mapper_persona.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-05-01
// Description: Funcionalidad de conversión entre estructuras SQL de persistencia a estructuras o entidades de dominio

package mapper

import (
	"api-plantilla/domain/persona"
	entitysql "api-plantilla/infrastructure/persistence/dbsql/entity_sql"
	"api-plantilla/infrastructure/persistence/dbsql/util"

	"github.com/shopspring/decimal"
)

func ToEntity(objPersona *entitysql.PersonaSql) *persona.Persona {

	if objPersona == nil {
		return nil // Manejo de casos nulos
	}

	return &persona.Persona{
		Id:         util.NullToInt32(objPersona.Id, 0),
		Nombre:     util.NullToString(objPersona.Nombre, ""),
		ApPaterno:  util.NullToString(objPersona.ApPaterno, ""),
		ApMaterno:  util.NullToString(objPersona.ApMaterno, ""),
		Mail:       util.NullToString(objPersona.Mail, ""),
		NumCelular: util.NullToInt32(objPersona.NumCelular, 0),
		Edad:       util.NullToInt16(objPersona.Edad, 0),
		Renta:      util.NullToInt32(objPersona.Renta, 0),
		Impuesto:   util.NullToFloat64ToDecimal(objPersona.Impuesto, decimal.Zero),
	}
}

func ToListaToEntity(lista []*entitysql.PersonaSql) []*persona.Persona {

	var personas []*persona.Persona

	for _, objPersona := range lista {
		personas = append(personas, ToEntity(objPersona))
	}

	return personas
}
