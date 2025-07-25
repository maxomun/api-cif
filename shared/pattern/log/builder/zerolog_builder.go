// File: zerolog_builder.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Logica de manejo y de logs utilizando el framework de zerolog

// Cualquier cambio a la estructura en este archivo, debe ser previa autorizacion del area de arquitectura

package builder

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// LoggerBuilder define la estructura del builder
type LoggerBuilder struct {
	output   io.Writer
	level    zerolog.Level
	timeFmt  string
	withTime bool
}

// NewLoggerBuilder crea un nuevo builder con valores por defecto
func NewLoggerBuilder() *LoggerBuilder {
	return &LoggerBuilder{
		output:   os.Stdout,
		level:    zerolog.InfoLevel,
		timeFmt:  time.RFC3339,
		withTime: true,
	}
}

// SetOutput cambia la salida del logger
func (objeto *LoggerBuilder) SetOutput(output io.Writer) *LoggerBuilder {
	objeto.output = output
	return objeto
}

// SetLevel establece el nivel de logging
func (objeto *LoggerBuilder) SetLevel(level zerolog.Level) *LoggerBuilder {
	objeto.level = level
	return objeto
}

// SetTimeFormat define el formato de tiempo
func (objeto *LoggerBuilder) SetTimeFormat(format string) *LoggerBuilder {
	objeto.timeFmt = format
	return objeto
}

// DisableTimestamp desactiva el timestamp
func (objeto *LoggerBuilder) DisableTimestamp() *LoggerBuilder {
	objeto.withTime = false
	return objeto
}

// Build crea el logger configurado
func (objeto *LoggerBuilder) Build() zerolog.Logger {
	logger := zerolog.New(objeto.output).Level(objeto.level)
	if objeto.withTime {
		logger = logger.With().Timestamp().Logger()
	}
	return logger
}
