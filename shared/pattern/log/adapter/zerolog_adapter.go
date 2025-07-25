// File: zerolog_adapter.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Fachada unica para utilizar cualquier implementación de log

// Cualquier cambio a la estructura en este archivo, debe ser previa autorizacion del area de arquitectura

package adapter

import (
	"io"

	"github.com/rs/zerolog"
)

// ZeroLogger es un wrapper para zerolog
type ZeroLogger struct {
	logger zerolog.Logger
}

// NewZeroLogger crea un ZeroLogger
func NewZeroLogger(output io.Writer, level zerolog.Level) *ZeroLogger {
	return &ZeroLogger{
		logger: zerolog.New(output).Level(level).With().Timestamp().Logger(),
	}
}

// Implementación de la interfaz Logger
func (z *ZeroLogger) Info(msg string) {
	z.logger.Info().Msg(msg)
}

func (z *ZeroLogger) Debug(msg string) {
	z.logger.Debug().Msg(msg)
}

func (z *ZeroLogger) Error(msg string) {
	z.logger.Error().Msg(msg)
}
