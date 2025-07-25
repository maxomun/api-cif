// File: service.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Reglas de negocio

package cliente

// Reglas de negocio básicas para el dominio cliente
// Por ahora no hay reglas específicas, solo mapeo de datos
func (objCliente *Cliente) EsValido() bool {
	// Validación básica: el cliente debe tener un ID válido
	return objCliente.Id > 0
}
