// File: repository.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Repositorio define las operaciones de persistencia

package persona

type DBaseRepository interface {
	Actualizar(objeto *Persona) (bool, error)
	BuscarPorId(id int) (*Persona, error)
	Crear(objeto *Persona) (int, error)
	Listar() ([]*Persona, error)
}
