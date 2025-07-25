// File: mapper_to_dto.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Tranformación de entidades del dominio en una o varias representaciones DTO

package persona

import (
	"api-plantilla/application/persona/dto"
	"api-plantilla/domain/persona"
)

// ToPersonaDto convierte una entidad Persona a un DTO de respuesta.
func ToDto(objPersona *persona.Persona) *dto.PersonaDto {

	if objPersona == nil {
		return nil // Manejo de casos nulos
	}

	return &dto.PersonaDto{
		Id:         objPersona.Id,
		Nombre:     objPersona.Nombre,
		ApPaterno:  objPersona.ApPaterno,
		ApMaterno:  objPersona.ApMaterno,
		Mail:       objPersona.Mail,
		NumCelular: objPersona.NumCelular,
		Edad:       objPersona.Edad,
		Renta:      objPersona.Renta,
		Impuesto:   objPersona.Impuesto,
	}
}

func ToListaDto(lista []*persona.Persona) []*dto.PersonaDto {

	var personas []*dto.PersonaDto

	for _, objPersona := range lista {
		personas = append(personas, ToDto(objPersona))
	}

	return personas
}
