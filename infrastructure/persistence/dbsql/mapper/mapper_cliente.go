// File: mapper_cliente.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Mapper para convertir entidades SQL a entidades del dominio

package mapper

import (
	"api-cif/domain/cliente"
	entitysql "api-cif/infrastructure/persistence/dbsql/entity_sql"
)

// ToClienteEntity convierte una entidad SQL Cliente a entidad del dominio
func ToClienteEntity(objClienteSql *entitysql.ClienteSql) *cliente.Cliente {
	if objClienteSql == nil {
		return nil
	}

	return &cliente.Cliente{
		Id:                         objClienteSql.Id.Int64,
		IdDocumento:                objClienteSql.IdDocumento.String,
		TipoDocumento:              toTipoDocumentoEntity(objClienteSql),
		Estado:                     toEstadoEntity(objClienteSql),
		TipoEntidad:                toTipoEntidadEntity(objClienteSql),
		ActividadEconomica:         toActividadEconomicaEntity(objClienteSql),
		ComposicionInstitucional:   toComposicionInstitucionalEntity(objClienteSql),
		CodigoContraparte:          toCodigoContraparteEntity(objClienteSql),
		IndicadorValidez:           toIndicadorValidezEntity(objClienteSql),
		RelacionClienteInstitucion: toRelacionClienteInstitucionEntity(objClienteSql),
		TipoSociedad:               toTipoSociedadEntity(objClienteSql),
		Nombre:                     objClienteSql.Nombre.String,
		FechaIngreso:               objClienteSql.FechaIngreso.Time,
	}
}

// Funciones auxiliares para mapear entidades anidadas
func toTipoDocumentoEntity(cs *entitysql.ClienteSql) *cliente.TipoDocumento {
	if !cs.TipoDocumentoId.Valid {
		return nil
	}
	return &cliente.TipoDocumento{
		Id:          int(cs.TipoDocumentoId.Int64),
		Descripcion: cs.TipoDocumentoDescripcion.String,
	}
}

func toEstadoEntity(cs *entitysql.ClienteSql) *cliente.Estado {
	if !cs.EstadoId.Valid {
		return nil
	}
	return &cliente.Estado{
		Id:          int(cs.EstadoId.Int64),
		Descripcion: cs.EstadoDescripcion.String,
	}
}

func toTipoEntidadEntity(cs *entitysql.ClienteSql) *cliente.TipoEntidad {
	if !cs.TipoEntidadId.Valid {
		return nil
	}
	return &cliente.TipoEntidad{
		Id:          int(cs.TipoEntidadId.Int64),
		Descripcion: cs.TipoEntidadDescripcion.String,
	}
}

func toActividadEconomicaEntity(cs *entitysql.ClienteSql) *cliente.ActividadEconomica {
	if !cs.ActividadEconomicaId.Valid {
		return nil
	}
	return &cliente.ActividadEconomica{
		Id:          int(cs.ActividadEconomicaId.Int64),
		Descripcion: cs.ActividadEconomicaDescripcion.String,
		CodigoCIIU:  int(cs.ActividadEconomicaCodigoCIIU.Int64),
	}
}

func toComposicionInstitucionalEntity(cs *entitysql.ClienteSql) *cliente.ComposicionInstitucional {
	if !cs.ComposicionInstitucionalId.Valid {
		return nil
	}
	return &cliente.ComposicionInstitucional{
		Id:          int(cs.ComposicionInstitucionalId.Int64),
		Descripcion: cs.ComposicionInstitucionalDescripcion.String,
	}
}

func toCodigoContraparteEntity(cs *entitysql.ClienteSql) *cliente.CodigoContraparte {
	if !cs.CodigoContraparteId.Valid {
		return nil
	}
	return &cliente.CodigoContraparte{
		Id:          int(cs.CodigoContraparteId.Int64),
		Descripcion: cs.CodigoContraparteDescripcion.String,
	}
}

func toIndicadorValidezEntity(cs *entitysql.ClienteSql) *cliente.IndicadorValidez {
	if !cs.IndicadorValidezId.Valid {
		return nil
	}
	return &cliente.IndicadorValidez{
		Id:          cs.IndicadorValidezId.String,
		Descripcion: cs.IndicadorValidezDescripcion.String,
	}
}

func toRelacionClienteInstitucionEntity(cs *entitysql.ClienteSql) *cliente.RelacionClienteInstitucion {
	if !cs.RelacionClienteInstitucionId.Valid {
		return nil
	}
	return &cliente.RelacionClienteInstitucion{
		Id:          int(cs.RelacionClienteInstitucionId.Int64),
		Descripcion: cs.RelacionClienteInstitucionDescripcion.String,
	}
}

func toTipoSociedadEntity(cs *entitysql.ClienteSql) *cliente.TipoSociedad {
	if !cs.TipoSociedadId.Valid {
		return nil
	}
	return &cliente.TipoSociedad{
		Id:             int(cs.TipoSociedadId.Int64),
		Descripcion:    cs.TipoSociedadDescripcion.String,
		IdKYC:          int(cs.TipoSociedadIdKYC.Int64),
		DescripcionKYC: cs.TipoSociedadDescripcionKYC.String,
	}
}
