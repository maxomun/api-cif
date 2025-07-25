// File: router.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Declaración de las rutas expuestas en el API Rest

package router

import (
	apiHttp "api-plantilla/infrastructure/api/http"
	"api-plantilla/shared"

	"net/http"

	docInfo "api-plantilla/docs"

	"github.com/gin-gonic/gin"
	swaggerFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type AppRouter struct {
	objRouter *gin.Engine
}

func NewAppRouter() *AppRouter {
	gin.SetMode(gin.ReleaseMode)
	return &AppRouter{objRouter: gin.Default()}
}

func (objRuta *AppRouter) AgregarRutasPersona(objCntr *apiHttp.PersonaHandler) {

	objRuta.objRouter.GET("/personas/:id_persona", objCntr.BuscarPorId)
	objRuta.objRouter.POST("/personas", objCntr.Crear)
	objRuta.objRouter.PUT("/personas", objCntr.Actualizar)
	objRuta.objRouter.GET("/personas", objCntr.Listar)
}

func (objRuta *AppRouter) AgregarRutasPrueba(objCntr *apiHttp.PruebaHandler) {

	objRuta.objRouter.GET("/pruebas/:id_persona", objCntr.BuscarPorId)
	objRuta.objRouter.PUT("/pruebas", objCntr.Actualizar)
}

func (objRuta *AppRouter) AgregarSwagger(config shared.Configuracion) {

	docInfo.SwaggerInfo.Title = config.ApiNombre
	docInfo.SwaggerInfo.Description = config.ApiDescripcion
	docInfo.SwaggerInfo.Version = config.ApiVersion
	docInfo.SwaggerInfo.Schemes = []string{"http", "https"}

	objRuta.objRouter.GET(config.ApiSwagger, ginSwagger.WrapHandler(swaggerFile.Handler))
}

func (objRuta *AppRouter) Iniciar(puerto string) {

	objServidor := &http.Server{Addr: ":" + puerto, Handler: objRuta.objRouter}

	objError := objServidor.ListenAndServe()

	if objError != nil {
		panic(objError)
	}
}
