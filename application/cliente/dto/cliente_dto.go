// File: cliente_dto.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Estructuras de transferencias en respuesta a los consumidores en los handler

package dto

import "time"

// DTOs para las entidades anidadas
type TipoDocumentoDto struct {
	Id          int    `json:"id"`
	Descripcion string `json:"descripcion"`
}

type EstadoDto struct {
	Id          int    `json:"id"`
	Descripcion string `json:"descripcion"`
}

type TipoEntidadDto struct {
	Id          int    `json:"id"`
	Descripcion string `json:"descripcion"`
}

type ActividadEconomicaDto struct {
	Id          int    `json:"id"`
	Descripcion string `json:"descripcion"`
	CodigoCIIU  int    `json:"codigoCIIU"`
}

type ComposicionInstitucionalDto struct {
	Id          int    `json:"id"`
	Descripcion string `json:"descripcion"`
}

type CodigoContraparteDto struct {
	Id          int    `json:"id"`
	Descripcion string `json:"descripcion"`
}

type IndicadorValidezDto struct {
	Id          string `json:"id"`
	Descripcion string `json:"descripcion"`
}

type RelacionClienteInstitucionDto struct {
	Id          int    `json:"id"`
	Descripcion string `json:"descripcion"`
}

type TipoSociedadDto struct {
	Id             int    `json:"id"`
	Descripcion    string `json:"descripcion"`
	IdKYC          int    `json:"idKYC"`
	DescripcionKYC string `json:"descripcionKYC"`
}

// ClienteDto - DTO principal de respuesta
type ClienteDto struct {
	Id                         int64                          `json:"id"`
	IdDocumento                string                         `json:"idDocumento"`
	TipoDocumento              *TipoDocumentoDto              `json:"tipoDocumento"`
	Estado                     *EstadoDto                     `json:"estado"`
	TipoEntidad                *TipoEntidadDto                `json:"tipoEntidad"`
	ActividadEconomica         *ActividadEconomicaDto         `json:"actividadEconomica"`
	ComposicionInstitucional   *ComposicionInstitucionalDto   `json:"composicionInstitucional"`
	CodigoContraparte          *CodigoContraparteDto          `json:"codigoContraparte"`
	IndicadorValidez           *IndicadorValidezDto           `json:"indicadorValidez"`
	RelacionClienteInstitucion *RelacionClienteInstitucionDto `json:"relacionClienteInstitucion"`
	TipoSociedad               *TipoSociedadDto               `json:"tipoSociedad"`
	Nombre                     string                         `json:"nombre"`
	FechaIngreso               time.Time                      `json:"fechaIngreso"`
}
