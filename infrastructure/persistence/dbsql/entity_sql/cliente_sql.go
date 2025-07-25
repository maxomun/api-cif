// File: cliente_sql.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Entidades SQL para mapeo de base de datos

package entitysql

import (
	"database/sql"
)

// Entidades SQL para las tablas relacionadas
type TipoDocumentoSql struct {
	Id          sql.NullInt64  `db:"PK_tipo_documento_id"`
	Descripcion sql.NullString `db:"descripcion"`
}

type EstadoClienteSql struct {
	Id          sql.NullInt64  `db:"PK_estado_cliente_id"`
	Descripcion sql.NullString `db:"descripcion"`
}

type TipoPersonaSql struct {
	Id          sql.NullInt64  `db:"PK_tipo_persona_id"`
	Descripcion sql.NullString `db:"descripcion"`
}

type ActividadEconomicaSql struct {
	Id          sql.NullInt64  `db:"PK_actividad_economica_id"`
	Descripcion sql.NullString `db:"descripcion"`
	CodigoCIIU  sql.NullInt64  `db:"codigo_CIIU"`
}

type ComposicionInstitucionalSql struct {
	Id          sql.NullInt64  `db:"PK_composicion_institucional_id"`
	Descripcion sql.NullString `db:"descripcion"`
}

type CodigoContraparteSql struct {
	Id          sql.NullInt64  `db:"PK_codigo_contraparte_id"`
	Descripcion sql.NullString `db:"descripcion"`
}

type IndicadorValidezSql struct {
	Id          sql.NullString `db:"PK_indicador_validez_id"`
	Descripcion sql.NullString `db:"descripcion"`
}

type RelacionClienteInstitucionSql struct {
	Id          sql.NullInt64  `db:"PK_relacion_cliente_institucion_id"`
	Descripcion sql.NullString `db:"descripcion"`
}

type TipoSociedadSql struct {
	Id             sql.NullInt64  `db:"PK_tipo_sociedad_id"`
	Descripcion    sql.NullString `db:"descripcion"`
	IdKYC          sql.NullInt64  `db:"ID_KYC"`
	DescripcionKYC sql.NullString `db:"descripcion KYC"`
}

// ClienteSql - Entidad principal para mapear el resultado de la consulta
type ClienteSql struct {
	Id           sql.NullInt64  `db:"id"`
	IdDocumento  sql.NullString `db:"idDocumento"`
	Nombre       sql.NullString `db:"nombre"`
	FechaIngreso sql.NullTime   `db:"fechaIngreso"`

	// Campos para tipo documento
	TipoDocumentoId          sql.NullInt64  `db:"tipoDocumento.id"`
	TipoDocumentoDescripcion sql.NullString `db:"tipoDocumento.descripcion"`

	// Campos para estado
	EstadoId          sql.NullInt64  `db:"estado.id"`
	EstadoDescripcion sql.NullString `db:"estado.descripcion"`

	// Campos para tipo entidad
	TipoEntidadId          sql.NullInt64  `db:"tipoEntidad.id"`
	TipoEntidadDescripcion sql.NullString `db:"tipoEntidad.descripcion"`

	// Campos para actividad económica
	ActividadEconomicaId          sql.NullInt64  `db:"actividadEconomica.id"`
	ActividadEconomicaDescripcion sql.NullString `db:"actividadEconomica.descripcion"`
	ActividadEconomicaCodigoCIIU  sql.NullInt64  `db:"actividadEconomica.codigoCIIU"`

	// Campos para composición institucional
	ComposicionInstitucionalId          sql.NullInt64  `db:"composicionInstitucional.id"`
	ComposicionInstitucionalDescripcion sql.NullString `db:"composicionInstitucional.descripcion"`

	// Campos para código contraparte
	CodigoContraparteId          sql.NullInt64  `db:"codigoContraparte.id"`
	CodigoContraparteDescripcion sql.NullString `db:"codigoContraparte.descripcion"`

	// Campos para indicador validez
	IndicadorValidezId          sql.NullString `db:"indicadorValidez.id"`
	IndicadorValidezDescripcion sql.NullString `db:"indicadorValidez.descripcion"`

	// Campos para relación cliente institución
	RelacionClienteInstitucionId          sql.NullInt64  `db:"relacionClienteInstitucion.id"`
	RelacionClienteInstitucionDescripcion sql.NullString `db:"relacionClienteInstitucion.descripcion"`

	// Campos para tipo sociedad
	TipoSociedadId             sql.NullInt64  `db:"tipoSociedad.id"`
	TipoSociedadDescripcion    sql.NullString `db:"tipoSociedad.descripcion"`
	TipoSociedadIdKYC          sql.NullInt64  `db:"tipoSociedad.idKYC"`
	TipoSociedadDescripcionKYC sql.NullString `db:"tipoSociedad.descripcionKYC"`
}
