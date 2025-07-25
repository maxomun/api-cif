// File: repository.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Repositorio define las operaciones de persistencia

package cliente

type DBaseRepository interface {
	BuscarPorId(id int64) (*Cliente, error)
}
