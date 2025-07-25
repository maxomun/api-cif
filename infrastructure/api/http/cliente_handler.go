// File: cliente_handler.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Implementación de interfaces externas REST

package http

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	appCliente "api-cif/application/cliente"
	"api-cif/infrastructure/persistence/dbsql"
	shared "api-cif/shared"
	logFactory "api-cif/shared/pattern/log"

	"github.com/gin-gonic/gin"
)

const CLIENTE_PARAM_IDCLIENTE = "id"

const CLIENTE_METH_BUSCAR_POR_ID = "Metodo buscarPorId"
const CLIENTE_CONVERTIR_ID_CLIENTE = "Error al convertir a valor numerico idcliente"

const CLIENTE_ERROR_BUSCAR_POR_ID = "Error al instanciar objeto persistencia"

const CLIENTE_ERROR_ZEROLOG = "Error al crear una instancia de ZeroLog"
const CLIENTE_ERROR_INSTANCIA_BD = "Error al crear instancia de persistencia"
const CLIENTE_ERROR_SERIALIZAR = "Error al serializar "
const CLIENTE_ERROR_SOLICITUD = "Error en la solicitud de ejecución del metodo"
const CLIENTE_ERROR_BUSCAR = "Error al buscar en base de datos"
const CLIENTE_ERROR_INTERNO = "Servicio API con problemas en su funcionalidad"

const CLIENTE_RESPUESTA_OK = "Metodo ejecutada correctamente"
const CLIENTE_RESPUESTA_NO_OK_POR_REGLAS = "Metodo ejecutado sin resultados"

type ClienteHandler struct {
	objSql    *sql.DB
	objLogger logFactory.Logger
	objConfig shared.Configuracion
}

func NewClienteHandler(db *sql.DB, logger logFactory.Logger, config shared.Configuracion) *ClienteHandler {
	return &ClienteHandler{objSql: db, objLogger: logger, objConfig: config}
}

// BuscarPorId es el handler para buscar un cliente por ID
// @Summary Busca cliente por ID
// @Description Busca un cliente basado en el ID proporcionado
// @Tags ClienteHandler
// @Accept json
// @Produce json
// @Param id path int true "ID del cliente"
// @Success 200 {object} shared.MensajeDto[dto.ClienteDto]{}
// @Success 204 {object} shared.MensajeDto[string]{}
// @Failure 400 {object} shared.MensajeDto[string]{}
// @Failure 500 {object} shared.MensajeDto[string]{}
// @Router /clientes/{id} [get]
func (objCntr *ClienteHandler) BuscarPorId(objCtx *gin.Context) {

	var objCadena strings.Builder

	idCliente, objErrorSolicitud := strconv.ParseInt(objCtx.Param(CLIENTE_PARAM_IDCLIENTE), 10, 64)

	if idCliente > 0 && objErrorSolicitud == nil {

		objCadena.WriteString(CLIENTE_PARAM_IDCLIENTE + "-" + strconv.FormatInt(idCliente, 10))

		logFactory.LogMensaje(objCntr.objLogger, CLIENTE_METH_BUSCAR_POR_ID+objCadena.String(), logFactory.Info)

		objClienteReptry, objErrorReptry := dbsql.NewClienteRepository(objCntr.objSql)

		// Aqui solo se evalua que no existan errores, de ser asi, siempre existira una instancia de persistencia
		if objErrorReptry == nil {

			objClienteBusiness := appCliente.NewClienteBusiness(objClienteReptry)

			objCliente, objErrorBuscarPorId := objClienteBusiness.BuscarPorId(idCliente)

			if objErrorBuscarPorId != nil {

				if shared.EsRegla(objErrorBuscarPorId) {
					objCtx.JSON(shared.Respuesta("", http.StatusNotFound, CLIENTE_RESPUESTA_NO_OK_POR_REGLAS))
				} else {
					// Error 500 sin body
					objCtx.Status(http.StatusInternalServerError)
				}

				logFactory.LogMensaje(objCntr.objLogger, objErrorBuscarPorId.Error(), logFactory.Error)
			} else {
				if objCliente != nil {

					// Retornar directamente el objeto JSON plano sin envoltura
					objCtx.JSON(http.StatusOK, objCliente)

					// Log del mensaje de éxito
					logFactory.LogMensaje(objCntr.objLogger, CLIENTE_RESPUESTA_OK, logFactory.Info)
				} else {
					// Cliente no encontrado
					objCtx.JSON(shared.Respuesta("", http.StatusNotFound, CLIENTE_RESPUESTA_NO_OK_POR_REGLAS))
				}
			}

		} else {
			logFactory.LogMensaje(objCntr.objLogger, objErrorReptry.Error(), logFactory.Error)

			// Error 500 sin body
			objCtx.Status(http.StatusInternalServerError)
		}
	} else {
		// Error 400 sin body
		objCtx.Status(http.StatusBadRequest)
	}

	objCtx.Header(shared.API_TIPO_CONTENIDO, shared.API_TIPO_RESPUESTA)
}
