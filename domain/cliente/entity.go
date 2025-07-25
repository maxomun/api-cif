// File: entity.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Entidades del dominio del negocio

package cliente

import (
	"time"
)

type TipoDocumento struct {
	Id          int    `json:"id"`
	Descripcion string `json:"descripcion"`
}

type Estado struct {
	Id          int    `json:"id"`
	Descripcion string `json:"descripcion"`
}

type TipoEntidad struct {
	Id          int    `json:"id"`
	Descripcion string `json:"descripcion"`
}

type ActividadEconomica struct {
	Id          int    `json:"id"`
	Descripcion string `json:"descripcion"`
	CodigoCIIU  int    `json:"codigoCIIU"`
}

type ComposicionInstitucional struct {
	Id          int    `json:"id"`
	Descripcion string `json:"descripcion"`
}

type CodigoContraparte struct {
	Id          int    `json:"id"`
	Descripcion string `json:"descripcion"`
}

type IndicadorValidez struct {
	Id          string `json:"id"`
	Descripcion string `json:"descripcion"`
}

type RelacionClienteInstitucion struct {
	Id          int    `json:"id"`
	Descripcion string `json:"descripcion"`
}

type TipoSociedad struct {
	Id             int    `json:"id"`
	Descripcion    string `json:"descripcion"`
	IdKYC          int    `json:"idKYC"`
	DescripcionKYC string `json:"descripcionKYC"`
}

type Cliente struct {
	Id                         int64
	IdDocumento                string
	TipoDocumento              *TipoDocumento
	Estado                     *Estado
	TipoEntidad                *TipoEntidad
	ActividadEconomica         *ActividadEconomica
	ComposicionInstitucional   *ComposicionInstitucional
	CodigoContraparte          *CodigoContraparte
	IndicadorValidez           *IndicadorValidez
	RelacionClienteInstitucion *RelacionClienteInstitucion
	TipoSociedad               *TipoSociedad
	Nombre                     string
	FechaIngreso               time.Time
}

func NewCliente(id int64,
	idDocumento string,
	tipoDocumento *TipoDocumento,
	estado *Estado,
	tipoEntidad *TipoEntidad,
	actividadEconomica *ActividadEconomica,
	composicionInstitucional *ComposicionInstitucional,
	codigoContraparte *CodigoContraparte,
	indicadorValidez *IndicadorValidez,
	relacionClienteInstitucion *RelacionClienteInstitucion,
	tipoSociedad *TipoSociedad,
	nombre string,
	fechaIngreso time.Time) (*Cliente, error) {

	return &Cliente{
		Id:                         id,
		IdDocumento:                idDocumento,
		TipoDocumento:              tipoDocumento,
		Estado:                     estado,
		TipoEntidad:                tipoEntidad,
		ActividadEconomica:         actividadEconomica,
		ComposicionInstitucional:   composicionInstitucional,
		CodigoContraparte:          codigoContraparte,
		IndicadorValidez:           indicadorValidez,
		RelacionClienteInstitucion: relacionClienteInstitucion,
		TipoSociedad:               tipoSociedad,
		Nombre:                     nombre,
		FechaIngreso:               fechaIngreso,
	}, nil
}
