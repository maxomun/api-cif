// File: util.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Funcionalidades de apoyo transversales

// Cualquier cambio a la estructura en este archivo, debe ser previa autorizacion del area de arquitectura

package shared

import (
	"encoding/json"
	"errors"
	"os"
)

const PLTL_AMBIENTE = "APP_ENV"
const PLTL_NOMBRE_AMBIENTE = "local"

const PLTL_PARAM_CONFIGURACION = "configuracion.json"
const PLTL_ERROR_SERIALIZARCION = "Error al deserializar el objeto"

const API_TIPO_CONTENIDO = "Content-Type"
const API_TIPO_RESPUESTA = "application/json"

// Serializar convierte MensajeDto a JSON.
func (objMensaje *MensajeDto[T]) Serializar() ([]byte, error) {
	return json.Marshal(objMensaje)
}

// Deserializar convierte JSON a MensajeDto.
func (objMensaje *MensajeDto[T]) Deserializar(objeto []byte) error {
	return json.Unmarshal(objeto, objMensaje)
}

func Respuesta[T any](objeto T, tipo int, mensaje string) (int, MensajeDto[T]) {

	objMensajeDto := MensajeDto[T]{
		Objeto: objeto,
		Estado: EstadoHttpDto{
			Codigo:  tipo,
			Mensaje: mensaje,
		},
	}

	return tipo, objMensajeDto
}

func LeerConfiguracion(nombre string) (Configuracion, error) {

	var objKeyConfig []byte
	var objConfig Configuracion
	var objErrorLectura error
	var objErrorConfig error

	strAppAmbiente := os.Getenv(PLTL_AMBIENTE)

	if strAppAmbiente == PLTL_NOMBRE_AMBIENTE {

		objKeyConfig, objErrorLectura = os.ReadFile(PLTL_PARAM_CONFIGURACION)

		if objErrorLectura != nil {
			return objConfig, objErrorLectura
		}
	} else {

		objKeyConfig = []byte(os.Getenv(nombre))
	}

	objErrorConfig = json.Unmarshal(objKeyConfig, &objConfig)

	if objErrorConfig != nil {

		return objConfig, errors.New(PLTL_ERROR_SERIALIZARCION + objErrorConfig.Error())
	}

	return objConfig, nil
}
