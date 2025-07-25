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
	domPersona "api-plantilla/domain/persona"
	shared "api-plantilla/shared"
	logFactory "api-plantilla/shared/pattern/log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

const PERS_PARAM_IDPERSONA_PRUEBA = "id_persona"

type PruebaHandler struct {
	objLogger logFactory.Logger
	objConfig shared.Configuracion
}

func NewPruebaHandler(logger logFactory.Logger, config shared.Configuracion) *PruebaHandler {
	return &PruebaHandler{objLogger: logger, objConfig: config}
}

// Actualizar es el handler para actualizar una persona
// @Summary Actualiza persona
// @Description Actualiza una persona basado en los datos proporcionados
// @Tags PruebaHandler
// @Accept json
// @Produce json
// @Param user body appPersona.PersonaDto true "Información de la persona"
// @Success 200 {object} shared.MensajeDto[appPersona.PersonaDto]{}
// @Success 204 {object} shared.MensajeDto[string]{}
// @Failure 400 {object} shared.MensajeDto[string]{}
// @Failure 500 {object} shared.MensajeDto[string]{}
// @Router /pruebas [put]
func (objCntr *PruebaHandler) Actualizar(objCtx *gin.Context) {

	var objPersonaDto dto.PersonaDto
	var objMensajeDto shared.MensajeDto[dto.PersonaDto]

	if objErrorSolicitud := objCtx.ShouldBindJSON(&objPersonaDto); objErrorSolicitud != nil {
		objCtx.JSON(http.StatusBadRequest, gin.H{"error": objErrorSolicitud.Error()})
		return
	}

	objPersonaDto.Nombre = "Prueba Servicio"

	_, objMensajeDto = shared.Respuesta(objPersonaDto, http.StatusOK, PERS_RESPUESTA_OK)

	strMensaje, objPersonaSerializar := objMensajeDto.Serializar()

	if objPersonaSerializar != nil {
		logFactory.LogMensaje(objCntr.objLogger, PERS_ERROR_SERIALIZAR+string(strMensaje), logFactory.Error)
	}

	objCtx.JSON(http.StatusOK, objMensajeDto)

	objCtx.Header(shared.API_TIPO_CONTENIDO, shared.API_TIPO_RESPUESTA)
}

// BuscarPorId es un endpoint que devuelve un objeto tipo Persona
// @Summary      buscar persona por identificador
// @Description  Devuelve una persona
// @Tags         PruebaHandler
// @Accept       json
// @Produce      json
// @Param        id_persona path int true "id_persona"
// @Success      200 {object}  shared.MensajeDto[appPersona.PersonaDto]{}
// @Success 	 204 {object} shared.MensajeDto[string]{}
// @Failure 	 400 {object} shared.MensajeDto[string]{}
// @Failure 	 500 {object} shared.MensajeDto[string]{}
// @Router       /pruebas/buscar-por-criterio/{id_persona}  [get]
func (objCntr *PruebaHandler) BuscarPorId(objCtx *gin.Context) {

	var objCadena strings.Builder
	var objMensajeDto shared.MensajeDto[dto.PersonaDto]

	idPersona, objErrorParEntrada := strconv.Atoi(objCtx.Param(PERS_PARAM_IDPERSONA_PRUEBA))

	if objErrorParEntrada != nil {
		logFactory.LogMensaje(objCntr.objLogger, objErrorParEntrada.Error(), logFactory.Error)
	}

	objCadena.WriteString("id Persona:" + strconv.Itoa(idPersona))

	// Estructura de ejemplo
	objPersona, objErrorConstructor := domPersona.NewPersona(idPersona, objCntr.objConfig.ApiNombre,
		objCntr.objConfig.ApiDescripcion,
		objCntr.objConfig.ApiSwagger,
		objCntr.objConfig.ApiVersion,
		111111111,
		1,
		500000,
		decimal.NewFromFloat(123.45))

	if objErrorConstructor != nil {
	}

	_, objMensajeDto = shared.Respuesta(*appPersona.ToDto(objPersona), http.StatusOK, PERS_RESPUESTA_OK)

	strMensaje, objPersonaSerializar := objMensajeDto.Serializar()

	if objPersonaSerializar != nil {
		logFactory.LogMensaje(objCntr.objLogger, PERS_ERROR_SERIALIZAR+string(strMensaje), logFactory.Error)
	}

	objCtx.JSON(shared.Respuesta(*appPersona.ToDto(objPersona), http.StatusOK, PERS_RESPUESTA_OK))

	objCtx.Header(shared.API_TIPO_CONTENIDO, shared.API_TIPO_RESPUESTA)
}
