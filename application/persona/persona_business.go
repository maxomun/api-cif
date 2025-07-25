// File: documento_bussiness.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Logica enfocadas en la necesidad del negocio

package persona

import (
	"api-plantilla/application/persona/dto"
	"api-plantilla/application/persona/dto/request"
	"api-plantilla/domain/persona"
	"api-plantilla/shared"

	"time"
)

// # Capa de aplicación (casos de uso, DTOs)

const ZONA_HORARIO = "America/Santiago"
const REGLA_MAXIMA_EDAD = "regla: no cumple con la regla de edad maxima"
const REGLA_IMPUESTO_PERMITIDO = "regla: monto impuesto por debajo a lo permitido"

type PersonaBusiness struct {
	objDBaseReptry persona.DBaseRepository
}

func NewPersonanBusiness(dbaseReptry persona.DBaseRepository) *PersonaBusiness {
	return &PersonaBusiness{
		objDBaseReptry: dbaseReptry,
	}
}

func (objPerBns *PersonaBusiness) Actualizar(objeto *dto.PersonaDto) (bool, error) {

	objZonaHoraria, objErrorLocation := time.LoadLocation(ZONA_HORARIO)

	if objErrorLocation != nil {
		return false, objErrorLocation
	}

	objPersona := dto.FromDto(objeto)
	objPersona.FechaRegistro = time.Now().In(objZonaHoraria)

	blnRespuesta, objErrorActualizar := objPerBns.objDBaseReptry.Actualizar(objPersona)

	if objErrorActualizar != nil {
		return false, objErrorActualizar
	}

	return blnRespuesta, nil
}

func (objPerBns *PersonaBusiness) BuscarPorId(idPersona int) (*dto.PersonaDto, error) {

	objPersona, objErrorBuscar := objPerBns.objDBaseReptry.BuscarPorId(idPersona)

	if objErrorBuscar != nil {
		return nil, objErrorBuscar
	} else if objPersona != nil {

		// Sección para agregar reglas de negocios

		listaExcpcnRegla := &shared.ListaExcpcnReglaNegocio{}

		// Primera regla
		if !objPersona.EdadMaxima() {
			listaExcpcnRegla.Agregar(&shared.ExcpcnReglaNegocio{Mensaje: REGLA_MAXIMA_EDAD})
		}

		// Segunda regla
		if objPersona.MontoImpuestoMinimo() {
			// El valor impuesto es valido, entonces realizamos un calculo para ajustar el monto hipotetico
			objPersona.Impuesto = objPersona.CalcularMontoFuturo()
		} else {
			listaExcpcnRegla.Agregar(&shared.ExcpcnReglaNegocio{Mensaje: REGLA_IMPUESTO_PERMITIDO})
		}

		if shared.NumReglas(listaExcpcnRegla) > 0 {
			return ToDto(objPersona), listaExcpcnRegla
		}

		return ToDto(objPersona), nil
	}

	return nil, nil
}

func (objPerBns *PersonaBusiness) Crear(objeto request.PersonaRequest) (*dto.PersonaDto, error) {

	var objErrorCrear error

	objZonaHoraria, objErrorLocation := time.LoadLocation(ZONA_HORARIO)

	if objErrorLocation != nil {
		return nil, objErrorLocation
	}

	objPersona := persona.Persona{
		Nombre:        objeto.Nombre,
		ApPaterno:     objeto.ApPaterno,
		ApMaterno:     objeto.ApMaterno,
		Mail:          objeto.Mail,
		NumCelular:    objeto.NumCelular,
		Edad:          objeto.Edad,
		Renta:         objeto.Renta,
		Impuesto:      objeto.Impuesto,
		FechaRegistro: time.Now().In(objZonaHoraria),
	}

	objPersona.Id, objErrorCrear = objPerBns.objDBaseReptry.Crear(&objPersona)

	if objErrorCrear != nil {
		return nil, objErrorCrear
	}

	return ToDto(&objPersona), nil
}

func (objPerBns *PersonaBusiness) Listar() ([]*dto.PersonaDto, error) {

	personas, objErrorListar := objPerBns.objDBaseReptry.Listar()

	if objErrorListar != nil {
		return nil, objErrorListar
	}

	return ToListaDto(personas), nil
}
