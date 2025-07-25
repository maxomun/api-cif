// File: excepcion_error.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-05-30
// Description: Estructura para soportar excepciones del tipo regla de negocio

// Cualquier cambio a la estructura en este archivo, debe ser previa autorizacion del area de arquitectura

package shared

import (
	"errors"
)

const SHARED_REGLA_NEGOCIO = ""

type ExcpcnReglaNegocio struct {
	Mensaje string
}

type ListaExcpcnReglaNegocio struct {
	Reglas []*ExcpcnReglaNegocio
}

func (objExpcn *ExcpcnReglaNegocio) Error() string {
	return SHARED_REGLA_NEGOCIO + objExpcn.Mensaje
}

func (listaobjExpcn *ListaExcpcnReglaNegocio) Error() string {

	var mensaje string

	for _, objRegla := range listaobjExpcn.Reglas {
		mensaje += "\n" + objRegla.Error()
	}

	return mensaje
}

func EsExcepcionRegla(objError error) bool {
	_, blnOk := objError.(*ExcpcnReglaNegocio)
	return blnOk
}

func EsExcepcionReglas(objError error) bool {
	var listaExcpcn *ListaExcpcnReglaNegocio
	return errors.As(objError, &listaExcpcn)
}

func NumReglas(objError error) int {
	var numero int = 0
	var listaExcpcn *ListaExcpcnReglaNegocio

	blnOk := errors.As(objError, &listaExcpcn)

	if blnOk {
		numero = len(listaExcpcn.Reglas)
	}

	return numero
}

func EsRegla(objError error) bool {

	var listaExcpcn *ListaExcpcnReglaNegocio

	_, blnOk := objError.(*ExcpcnReglaNegocio)

	if !blnOk {
		blnOk = errors.As(objError, &listaExcpcn)
	}

	return blnOk
}

func (lista *ListaExcpcnReglaNegocio) Agregar(objError *ExcpcnReglaNegocio) {
	lista.Reglas = append(lista.Reglas, objError)
}
