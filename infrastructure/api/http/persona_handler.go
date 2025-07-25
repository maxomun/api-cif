// File: handler.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Implementación de interfaces externas REST

package http

import (
	appPersona "api-plantilla/application/persona"
	"api-plantilla/application/persona/dto"
	"api-plantilla/application/persona/dto/request"
	"api-plantilla/infrastructure/persistence/dbsql"
	shared "api-plantilla/shared"
	logFactory "api-plantilla/shared/pattern/log"
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const PERS_PARAM_IDPERSONA = "id_persona"

const PERS_METH_BUSCAR_POR_ID = "Metodo buscarPorId"
const PERS_METH_LISTAR = "Metodo listar personas"
const PERS_CONVERTIR_ID_PERSONA = "Error al convertir a valor numerico idpersona"

const PERS_ERROR_BUSCAR_POR_ID = "Error al instanciar objeto persistencia"

const PERS_ERROR_ZEROLOG = "Error al crear una instancia de ZeroLog"
const PERS_ERROR_INSTANCIA_BD = "Error al crear instancia de persistencia"
const PERS_ERROR_SERIALIZAR = "Error al serializar "
const PERS_ERROR_SOLICITUD = "Error en la solicitud de ejecución del metodo"
const PERS_ERROR_BUSCAR = "Error al buscar en base de datos"
const PERS_ERROR_CREAR = "Error al crear el registro en base de datos"
const PERS_ERROR_ACTUALIZAR = "Error al actualizar el registro en base de datos"
const PERS_ERROR_INTERNO = "Servicio API con problemas en su funcionalidad"

const PERS_RESPUESTA_OK = "Metodo ejecutada correctamente"
const PERS_RESPUESTA_NO_OK_POR_REGLAS = "Metodo ejecutado sin resultados"

const API_CONFIGURACION = "CONFIGURACION"

var ErrRegistroNoExiste = errors.New("regla")

//var objConfig shared.Configuracion

type PersonaHandler struct {
	objSql    *sql.DB
	objLogger logFactory.Logger
	objConfig shared.Configuracion
}

func NewPersonaHandler(db *sql.DB, logger logFactory.Logger, config shared.Configuracion) *PersonaHandler {
	return &PersonaHandler{objSql: db, objLogger: logger, objConfig: config}
}

// Actualizar es el handler para actualizar una persona
// @Summary Actualiza persona
// @Description Actualiza una persona basado en los datos proporcionados
// @Tags PersonaHandler
// @Accept json
// @Produce json
// @Param user body appPersona.PersonaDto true "Información de la persona"
// @Success 200 {object} shared.MensajeDto[bool]{}
// @Success 204 {object} shared.MensajeDto[string]{}
// @Failure 400 {object} shared.MensajeDto[string]{}
// @Failure 500 {object} shared.MensajeDto[string]{}
// @Router /personas [put]
func (objCntr *PersonaHandler) Actualizar(objCtx *gin.Context) {

	var objErrorSolicitud error
	var objSolicitud dto.PersonaDto
	var objMensajeDto shared.MensajeDto[bool]

	objErrorSolicitud = objCtx.ShouldBindJSON(&objSolicitud)

	if objErrorSolicitud == nil {

		objPersonaReptry, objErrorReptry := dbsql.NewPersonaRepository(objCntr.objSql)

		if objErrorReptry == nil {

			objPersonaBusiness := appPersona.NewPersonanBusiness(objPersonaReptry)

			blnRespuesta, objErrorActualizar := objPersonaBusiness.Actualizar(&objSolicitud)

			if objErrorActualizar != nil {

				if shared.EsRegla(objErrorActualizar) {
					objCtx.JSON(shared.Respuesta("", http.StatusNotFound, PERS_RESPUESTA_NO_OK_POR_REGLAS))
				} else {
					objCtx.JSON(shared.Respuesta("", http.StatusInternalServerError, PERS_ERROR_INTERNO))
				}

				logFactory.LogMensaje(objCntr.objLogger, objErrorActualizar.Error(), logFactory.Error)
			} else {

				_, objMensajeDto = shared.Respuesta(blnRespuesta, http.StatusOK, PERS_RESPUESTA_OK)

				objCtx.JSON(http.StatusOK, objMensajeDto)

				strMensaje, objPersonaSerializar := objMensajeDto.Serializar()

				if objPersonaSerializar == nil {

					// Solo en casos de tener configuracion log como servicio externo, en ningun caso enviarlo por consonla
					logFactory.LogMensaje(objCntr.objLogger, string(strMensaje), logFactory.Info)
				}
			}
		} else {

			logFactory.LogMensaje(objCntr.objLogger, objErrorReptry.Error(), logFactory.Error)

			objCtx.JSON(shared.Respuesta("", http.StatusNotFound, PERS_ERROR_INSTANCIA_BD))
		}
	} else {

		logFactory.LogMensaje(objCntr.objLogger, objErrorSolicitud.Error(), logFactory.Error)

		objCtx.JSON(shared.Respuesta("", http.StatusBadRequest, PERS_ERROR_SOLICITUD))
	}

	objCtx.Header(shared.API_TIPO_CONTENIDO, shared.API_TIPO_RESPUESTA)
}

// BuscarPorId es un endpoint que devuelve una persona
// @Summary      buscar persona por identificador
// @Description  Devuelve una persona
// @Tags         PersonaHandler
// @Accept       json
// @Produce      json
// @Param        id_persona path int true "id_persona"
// @Success 200 {object} shared.MensajeDto[appPersona.PersonaDto]{}
// @Success 204 {object} shared.MensajeDto[string]{}
// @Failure 400 {object} shared.MensajeDto[string]{}
// @Failure 500 {object} shared.MensajeDto[string]{}
// @Router       /personas/{id_persona}  [get]
func (objCntr *PersonaHandler) BuscarPorId(objCtx *gin.Context) {

	var objCadena strings.Builder
	var objMensajeDto shared.MensajeDto[dto.PersonaDto]

	idPersona, objErrorSolicitud := strconv.Atoi(objCtx.Param(PERS_PARAM_IDPERSONA))

	if idPersona > 0 && objErrorSolicitud == nil {

		objCadena.WriteString(PERS_PARAM_IDPERSONA + "-" + strconv.Itoa(idPersona))

		logFactory.LogMensaje(objCntr.objLogger, PERS_METH_BUSCAR_POR_ID+objCadena.String(), logFactory.Info)

		objPersonaReptry, objErrorReptry := dbsql.NewPersonaRepository(objCntr.objSql)

		// Aqui solo se evalua que no existan errores, de ser asi, siempre existira una instancia de persistencia
		if objErrorReptry == nil {

			objPersonaBusiness := appPersona.NewPersonanBusiness(objPersonaReptry)

			objPersona, objErrorBuscarPorId := objPersonaBusiness.BuscarPorId(idPersona)

			if objErrorBuscarPorId != nil {

				if shared.EsRegla(objErrorBuscarPorId) {
					objCtx.JSON(shared.Respuesta("", http.StatusNotFound, PERS_RESPUESTA_NO_OK_POR_REGLAS))
				} else {
					objCtx.JSON(shared.Respuesta("", http.StatusInternalServerError, PERS_ERROR_INTERNO))
				}

				logFactory.LogMensaje(objCntr.objLogger, objErrorBuscarPorId.Error(), logFactory.Error)
			} else {
				if objPersona != nil {

					_, objMensajeDto = shared.Respuesta(*objPersona, http.StatusOK, PERS_RESPUESTA_OK)

					objCtx.JSON(http.StatusOK, objMensajeDto)

					strMensaje, objErrorSilzr := objMensajeDto.Serializar()

					if objErrorSilzr == nil {

						// Solo en casos de tener configuracion log como servicio externo, en ningun caso enviarlo por consonla
						logFactory.LogMensaje(objCntr.objLogger, string(strMensaje), logFactory.Info)
					}
				}
			}

		} else {
			logFactory.LogMensaje(objCntr.objLogger, objErrorReptry.Error(), logFactory.Error)

			objCtx.JSON(shared.Respuesta("", http.StatusInternalServerError, PERS_ERROR_INSTANCIA_BD))
		}
	} else {
		objCtx.JSON(shared.Respuesta("", http.StatusBadRequest, PERS_ERROR_SOLICITUD))
	}

	objCtx.Header(shared.API_TIPO_CONTENIDO, shared.API_TIPO_RESPUESTA)
}

// Crear es el handler para crear nueva persona
// @Summary Crea nueva persona
// @Description Crea una persona basado en los datos proporcionados
// @Tags PersonaHandler
// @Accept json
// @Produce json
// @Param persona body appPersona.PersonaRequest true "Información de la persona"
// @Success 200 {object} shared.MensajeDto[appPersona.PersonaDto]{}
// @Success 204 {object} shared.MensajeDto[string]{}
// @Failure 400 {object} shared.MensajeDto[string]{}
// @Failure 500 {object} shared.MensajeDto[string]{}
// @Router /personas [post]
func (objCntr *PersonaHandler) Crear(objCtx *gin.Context) {

	var objErrorSolicitud error
	var objSolicitud request.PersonaRequest
	var objMensajeDto shared.MensajeDto[dto.PersonaDto]

	objErrorSolicitud = objCtx.ShouldBindJSON(&objSolicitud)

	if objErrorSolicitud == nil {

		objPersonaReptry, objErrorReptry := dbsql.NewPersonaRepository(objCntr.objSql)

		if objErrorReptry == nil {

			objPersonaBusiness := appPersona.NewPersonanBusiness(objPersonaReptry)

			objPersonaDto, objErrorCrear := objPersonaBusiness.Crear(objSolicitud)

			if objErrorCrear == nil && objPersonaDto != nil {

				_, objMensajeDto = shared.Respuesta(*objPersonaDto, http.StatusOK, PERS_RESPUESTA_OK)

				objCtx.JSON(http.StatusOK, objMensajeDto)

				strMensaje, objPersonaSerializar := objMensajeDto.Serializar()

				if objPersonaSerializar == nil {

					// Solo en casos de tener configuracion log como servicio externo, en ningun caso enviarlo por consonla
					logFactory.LogMensaje(objCntr.objLogger, string(strMensaje), logFactory.Info)
				}
			} else {
				logFactory.LogMensaje(objCntr.objLogger, objErrorCrear.Error(), logFactory.Error)

				objCtx.JSON(shared.Respuesta("", http.StatusInternalServerError, PERS_ERROR_CREAR))
			}

		} else {

			logFactory.LogMensaje(objCntr.objLogger, objErrorReptry.Error(), logFactory.Error)

			objCtx.JSON(shared.Respuesta("", http.StatusNotFound, PERS_ERROR_INSTANCIA_BD))
		}

	} else {

		logFactory.LogMensaje(objCntr.objLogger, objErrorSolicitud.Error(), logFactory.Error)

		objCtx.JSON(shared.Respuesta("", http.StatusBadRequest, PERS_ERROR_SOLICITUD))
	}

	objCtx.Header(shared.API_TIPO_CONTENIDO, shared.API_TIPO_RESPUESTA)
}

// Listar es el handler para listar personas
// @Summary Listado de personas
// @Description Listado de personas vigentes
// @Tags PersonaHandler
// @Accept json
// @Produce json
// @Success 200 {object} shared.MensajeDto[appPersona.PersonaDto]{}
// @Success 204 {object} shared.MensajeDto[string]{}
// @Failure 400 {object} shared.MensajeDto[string]{}
// @Failure 500 {object} shared.MensajeDto[string]{}
// @Router /personas [get]
func (objCntr *PersonaHandler) Listar(objCtx *gin.Context) {

	var objMensajeDto shared.MensajeDto[[]*dto.PersonaDto]

	objPersonaReptry, objErrorReptry := dbsql.NewPersonaRepository(objCntr.objSql)

	if objErrorReptry == nil {

		objPersonaBusiness := appPersona.NewPersonanBusiness(objPersonaReptry)

		personas, objErrorListar := objPersonaBusiness.Listar()

		if objErrorListar != nil {

			if shared.EsRegla(objErrorListar) {
				objCtx.JSON(shared.Respuesta("", http.StatusNotFound, PERS_RESPUESTA_NO_OK_POR_REGLAS))
			} else {
				objCtx.JSON(shared.Respuesta("", http.StatusInternalServerError, PERS_ERROR_INTERNO))
			}

			logFactory.LogMensaje(objCntr.objLogger, objErrorListar.Error(), logFactory.Error)
		} else {
			_, objMensajeDto = shared.Respuesta(personas, http.StatusOK, PERS_RESPUESTA_OK)

			objCtx.JSON(http.StatusOK, objMensajeDto)
		}
	}

	objCtx.Header(shared.API_TIPO_CONTENIDO, shared.API_TIPO_RESPUESTA)
}
