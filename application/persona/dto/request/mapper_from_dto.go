package request

import (
	domPersona "api-plantilla/domain/persona"
)

type IOperacionDto interface {
	transformar() *domPersona.Persona
}

func (objPersDto *PersonaRequest) transformar() *domPersona.Persona {

	return &domPersona.Persona{
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
