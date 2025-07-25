// File: logger_interfsce.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Interface generica abstracta para ser utilizada por la implementación de un framework de log (zerolog, log tradicional, http)

// Cualquier cambio a la estructura en este archivo, debe ser previa autorizacion del area de arquitectura

package log

type Logger interface {
	Info(msg string)
	Debug(msg string)
	Error(msg string)
}
