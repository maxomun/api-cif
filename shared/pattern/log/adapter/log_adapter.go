// File: log_adapter.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Fachada unica para utilizar cualquier implementación de log

// Cualquier cambio a la estructura en este archivo, debe ser previa autorizacion del area de arquitectura

package adapter

import (
	"io"
	"log"
)

// StandardLogger es un wrapper para log.Logger
type StandardLogger struct {
	logger *log.Logger
}

// NewStandardLogger crea un StandardLogger
func NewStandardLogger(output io.Writer, prefix string) *StandardLogger {
	return &StandardLogger{
		logger: log.New(output, prefix, log.LstdFlags|log.Lshortfile),
	}
}

// Implementación de la interfaz Logger
func (s *StandardLogger) Info(msg string) {
	s.logger.Println("[INFO] " + msg)
}

func (s *StandardLogger) Debug(msg string) {
	s.logger.Println("[DEBUG] " + msg)
}

func (s *StandardLogger) Error(msg string) {
	s.logger.Println("[ERROR] " + msg)
}
