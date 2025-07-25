// File: log.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Implementacion del log para ser utilizado en todo el servicio API REST

// Cualquier cambio a la estructura en este archivo, debe ser previa autorizacion del area de arquitectura

package log

type NivelLog int

const (
	Info NivelLog = iota
	Debug
	Error
)

func LogMensaje(objLogger Logger, mensaje string, level NivelLog) {

	switch level {
	case Info:
		objLogger.Info(mensaje)
	case Debug:
		objLogger.Debug(mensaje)
	case Error:
		objLogger.Error(mensaje)
	default:
		objLogger.Info("Nivel no especificado, usando INFO: " + mensaje)
	}
}
