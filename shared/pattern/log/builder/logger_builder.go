// File: logger_builder.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Logica de manejo y de logs utilizando log tradicional por consola

// Cualquier cambio a la estructura en este archivo, debe ser previa autorizacion del area de arquitectura

package builder

import (
	"io"
	"log"
	"os"
)

// LogBuilder estructura del builder para log tradicional
type LogBuilder struct {
	output io.Writer
	prefix string
	flags  int
}

// NewLogBuilder crea un nuevo builder con valores por defecto
func NewLogBuilder() *LogBuilder {
	return &LogBuilder{
		output: os.Stdout,
		prefix: "",
		flags:  log.LstdFlags,
	}
}

// SetOutput cambia la salida del logger
func (objeto *LogBuilder) SetOutput(output io.Writer) *LogBuilder {
	objeto.output = output
	return objeto
}

// SetPrefix define un prefijo para los logs
func (objeto *LogBuilder) SetPrefix(prefix string) *LogBuilder {
	objeto.prefix = prefix
	return objeto
}

// SetFlags establece los flags del logger
func (objeto *LogBuilder) SetFlags(flags int) *LogBuilder {
	objeto.flags = flags
	return objeto
}

// Build crea el logger configurado
func (objeto *LogBuilder) Build() *log.Logger {
	return log.New(objeto.output, objeto.prefix, objeto.flags)
}
