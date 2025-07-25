// File: store.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Apoyo al usar parametros atraves de contextos en toda la aplicación

// Cualquier cambio a la estructura en este archivo, debe ser previa autorizacion del area de arquitectura

package config

import (
	"context"
	"sync"
)

// Es un tipo privado para evitar colisiones
type objIdLlave string

// Es la clave que se usará para almacenar el store en el contexto
const storeKey objIdLlave = "store"

// sync.Map para acceso concurrente seguro
var Store sync.Map

func WithStore(ctx context.Context, sysc *sync.Map) context.Context {
	return context.WithValue(ctx, storeKey, sysc)
}

func FromContext(ctx context.Context) (*sync.Map, bool) {

	sysc, ok := ctx.Value(storeKey).(*sync.Map)

	return sysc, ok
}
