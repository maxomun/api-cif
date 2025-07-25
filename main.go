// File: main.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Plantilla de servicios API, archivo de inicio del servicio

package main

import (
	apiHttp "api-cif/infrastructure/api/http"
	"api-cif/infrastructure/persistence/db"
	"api-cif/infrastructure/router"
	"api-cif/shared"

	"os"

	logFactory "api-cif/shared/pattern/log"
	logAdapter "api-cif/shared/pattern/log/adapter"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

const PLTL_CONFIGURACION_NOMBRE = "CONFIGURACION"

const PLTL_MSM_INICIO = "Iniciando Servicio...!"
const PLTL_MSM_INICIADO = "Servicio iniciado..."
const PLTL_MSM_ERROR_INICIO = "Deteniendo el servicio por no existir la configuración requerida..."
const PLTL_MSM_EXITO_CONEXION_BD = "Conectado a la base de datos"

func main() {

	// Implementar log utilizando modo ZeroLog
	objLogger := logAdapter.NewZeroLogger(os.Stdout, zerolog.InfoLevel)

	// Implementar log utilizando modo tradicional
	// objLogger := logAdapter.NewStandardLogger(os.Stdout, "[api-persona] ")

	/*
		// Implementar log utilizando modo HTTP (envio de log aun servicio API)
		   objHttpLog := logBuilder.NewHTTPLogWriterBuilder().
						SetEndpoint("https://example.com/api/logs").
						SetTimeout(10*time.Second).
						SetBufferSize(200).
						SetServiceName("api-persona").
						SetHostName("AKS").
						AddHeader("Authorization", "Bearer your_token").Build()
	*/

	// Carga en memoria el parametro en .env, el donde si es local, carga los parametros desde configuracion.json en coso contrario
	// desde el configMap
	_ = godotenv.Load()

	objConfig, objErrorVarConfigParam := shared.LeerConfiguracion(PLTL_CONFIGURACION_NOMBRE)

	if objErrorVarConfigParam != nil {
		logFactory.LogMensaje(objLogger, objErrorVarConfigParam.Error(), logFactory.Info)
		logFactory.LogMensaje(objLogger, PLTL_MSM_ERROR_INICIO, logFactory.Info)
		return
	}

	logFactory.LogMensaje(objLogger, PLTL_MSM_INICIO, logFactory.Info)

	// Inicializar la conexión a la base de datos
	objDbSql, objErrorCnxSql := db.IniciarConexionDbSql(objConfig)

	if objErrorCnxSql != nil {
		logFactory.LogMensaje(objLogger, objErrorCnxSql.Error(), logFactory.Error)
	} else {
		logFactory.LogMensaje(objLogger, PLTL_MSM_EXITO_CONEXION_BD, logFactory.Info)
	}

	// Cierra la conexión al finalizar la ejecución
	defer db.CerrarConexionDbSql()

	// Establece las rutas a exponer por el API
	objAppRouter := router.NewAppRouter()
	// Exponer las funcionalidades segun indicadas en las rutas
	objAppRouter.AgregarRutasCliente(apiHttp.NewClienteHandler(objDbSql, objLogger, objConfig))

	// Agrega Swagger
	objAppRouter.AgregarSwagger(objConfig)

	objAppRouter.Iniciar(objConfig.PuertoApi)

	logFactory.LogMensaje(objLogger, PLTL_MSM_INICIADO, logFactory.Info)
}
