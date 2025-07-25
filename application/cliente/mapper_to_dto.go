// File: mapper_to_dto.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Mapper para convertir entidades del dominio a DTOs

package cliente

import (
	"api-cif/application/cliente/dto"
	"api-cif/domain/cliente"
)

// ToDto convierte una entidad Cliente a ClienteDto
func ToDto(objCliente *cliente.Cliente) *dto.ClienteDto {
	if objCliente == nil {
		return nil
	}

	return &dto.ClienteDto{
		Id:                         objCliente.Id,
		IdDocumento:                objCliente.IdDocumento,
		TipoDocumento:              toTipoDocumentoDto(objCliente.TipoDocumento),
		Estado:                     toEstadoDto(objCliente.Estado),
		TipoEntidad:                toTipoEntidadDto(objCliente.TipoEntidad),
		ActividadEconomica:         toActividadEconomicaDto(objCliente.ActividadEconomica),
		ComposicionInstitucional:   toComposicionInstitucionalDto(objCliente.ComposicionInstitucional),
		CodigoContraparte:          toCodigoContraparteDto(objCliente.CodigoContraparte),
		IndicadorValidez:           toIndicadorValidezDto(objCliente.IndicadorValidez),
		RelacionClienteInstitucion: toRelacionClienteInstitucionDto(objCliente.RelacionClienteInstitucion),
		TipoSociedad:               toTipoSociedadDto(objCliente.TipoSociedad),
		Nombre:                     objCliente.Nombre,
		FechaIngreso:               objCliente.FechaIngreso,
	}
}

// Funciones auxiliares para mapear entidades anidadas
func toTipoDocumentoDto(td *cliente.TipoDocumento) *dto.TipoDocumentoDto {
	if td == nil {
		return nil
	}
	return &dto.TipoDocumentoDto{
		Id:          td.Id,
		Descripcion: td.Descripcion,
	}
}

func toEstadoDto(estado *cliente.Estado) *dto.EstadoDto {
	if estado == nil {
		return nil
	}
	return &dto.EstadoDto{
		Id:          estado.Id,
		Descripcion: estado.Descripcion,
	}
}

func toTipoEntidadDto(te *cliente.TipoEntidad) *dto.TipoEntidadDto {
	if te == nil {
		return nil
	}
	return &dto.TipoEntidadDto{
		Id:          te.Id,
		Descripcion: te.Descripcion,
	}
}

func toActividadEconomicaDto(ae *cliente.ActividadEconomica) *dto.ActividadEconomicaDto {
	if ae == nil {
		return nil
	}
	return &dto.ActividadEconomicaDto{
		Id:          ae.Id,
		Descripcion: ae.Descripcion,
		CodigoCIIU:  ae.CodigoCIIU,
	}
}

func toComposicionInstitucionalDto(ci *cliente.ComposicionInstitucional) *dto.ComposicionInstitucionalDto {
	if ci == nil {
		return nil
	}
	return &dto.ComposicionInstitucionalDto{
		Id:          ci.Id,
		Descripcion: ci.Descripcion,
	}
}

func toCodigoContraparteDto(cc *cliente.CodigoContraparte) *dto.CodigoContraparteDto {
	if cc == nil {
		return nil
	}
	return &dto.CodigoContraparteDto{
		Id:          cc.Id,
		Descripcion: cc.Descripcion,
	}
}

func toIndicadorValidezDto(iv *cliente.IndicadorValidez) *dto.IndicadorValidezDto {
	if iv == nil {
		return nil
	}
	return &dto.IndicadorValidezDto{
		Id:          iv.Id,
		Descripcion: iv.Descripcion,
	}
}

func toRelacionClienteInstitucionDto(rci *cliente.RelacionClienteInstitucion) *dto.RelacionClienteInstitucionDto {
	if rci == nil {
		return nil
	}
	return &dto.RelacionClienteInstitucionDto{
		Id:          rci.Id,
		Descripcion: rci.Descripcion,
	}
}

func toTipoSociedadDto(ts *cliente.TipoSociedad) *dto.TipoSociedadDto {
	if ts == nil {
		return nil
	}
	return &dto.TipoSociedadDto{
		Id:             ts.Id,
		Descripcion:    ts.Descripcion,
		IdKYC:          ts.IdKYC,
		DescripcionKYC: ts.DescripcionKYC,
	}
}
