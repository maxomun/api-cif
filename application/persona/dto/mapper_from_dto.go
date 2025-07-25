// File: mapper_from_dto.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Tranformación de representaciones DTO a una o varias entidades de dominio

package dto

import (
	domPersona "api-plantilla/domain/persona"
)

type IOperacionDto interface {
	transformar() *domPersona.Persona
}

func (objPersDto *PersonaDto) transformar() *domPersona.Persona {

	return &domPersona.Persona{
		Id:         objPersDto.Id,
		Nombre:     objPersDto.Nombre,
		ApPaterno:  objPersDto.ApPaterno,
		ApMaterno:  objPersDto.ApMaterno,
		Mail:       objPersDto.Mail,
		NumCelular: objPersDto.NumCelular,
		Edad:       objPersDto.Edad,
		Renta:      objPersDto.Renta,
		Impuesto:   objPersDto.Impuesto,
	}

}

func FromDto(objeto IOperacionDto) *domPersona.Persona {

	if objeto == nil {
		return nil
	}

	return objeto.transformar()
}
