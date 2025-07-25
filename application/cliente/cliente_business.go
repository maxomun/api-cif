// File: cliente_business.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Logica enfocadas en la necesidad del negocio

package cliente

import (
	"api-cif/application/cliente/dto"
	"api-cif/domain/cliente"
)

// # Capa de aplicación (casos de uso, DTOs)

type ClienteBusiness struct {
	objDBaseReptry cliente.DBaseRepository
}

func NewClienteBusiness(dbaseReptry cliente.DBaseRepository) *ClienteBusiness {
	return &ClienteBusiness{
		objDBaseReptry: dbaseReptry,
	}
}

func (objClienteBns *ClienteBusiness) BuscarPorId(idCliente int64) (*dto.ClienteDto, error) {

	objCliente, objErrorBuscar := objClienteBns.objDBaseReptry.BuscarPorId(idCliente)

	if objErrorBuscar != nil {
		return nil, objErrorBuscar
	} else if objCliente != nil {

		// Validación básica del cliente
		if !objCliente.EsValido() {
			return nil, nil // Cliente no válido, retornar nil
		}

		return ToDto(objCliente), nil
	}

	return nil, nil
}
