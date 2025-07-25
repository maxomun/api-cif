// File: cliente_reptry.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Operaciones con la base de datos relacional

package dbsql

import (
	"api-cif/domain/cliente"
	entitysql "api-cif/infrastructure/persistence/dbsql/entity_sql"
	"api-cif/infrastructure/persistence/dbsql/mapper"
	"api-cif/shared"
	"database/sql"
	"errors"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
)

const CLIENTE_NO_EXISTE = "Sin informacion en la base de datos"
const CLIENTE_GENERAL_ERROR = "Error al ejecutar la sentencia en la base de datos"

type ClienteReptry struct {
	objSql *sql.DB
}

func NewClienteRepository(conexion *sql.DB) (*ClienteReptry, error) {
	return &ClienteReptry{objSql: conexion}, nil
}

func (objReptry *ClienteReptry) BuscarPorId(idCliente int64) (*cliente.Cliente, error) {

	qryCliente := `SELECT
		c.PK_cliente_id AS id,
		c.documento_id AS idDocumento,
		c.nombre,
		c.fecha_primer_ingreso_bco AS fechaIngreso,
		td.PK_tipo_documento_id AS [tipoDocumento.id],
		td.descripcion AS [tipoDocumento.descripcion],
		ec.PK_estado_cliente_id AS [estado.id],
		ec.descripcion AS [estado.descripcion],
		tp.PK_tipo_persona_id AS [tipoEntidad.id],
		tp.descripcion AS [tipoEntidad.descripcion],
		ae.PK_actividad_economica_id AS [actividadEconomica.id],
		ae.descripcion AS [actividadEconomica.descripcion],
		ae.codigo_CIIU AS [actividadEconomica.codigoCIIU],
		ci.PK_composicion_institucional_id AS [composicionInstitucional.id],
		ci.descripcion AS [composicionInstitucional.descripcion],
		cc.PK_codigo_contraparte_id AS [codigoContraparte.id],
		cc.descripcion AS [codigoContraparte.descripcion],
		iv.PK_indicador_validez_id AS [indicadorValidez.id],
		iv.descripcion AS [indicadorValidez.descripcion],
		rci.PK_relacion_cliente_institucion_id AS [relacionClienteInstitucion.id],
		rci.descripcion AS [relacionClienteInstitucion.descripcion],
		ts.PK_tipo_sociedad_id AS [tipoSociedad.id],
		ts.descripcion AS [tipoSociedad.descripcion],
		ts.ID_KYC AS [tipoSociedad.idKYC],
		ts.[descripcion KYC] AS [tipoSociedad.descripcionKYC]
	FROM cif.tb_cliente c
	LEFT JOIN catalogo.tb_tipo_documento td ON td.PK_tipo_documento_id = c.FK_tipo_documento_id
	LEFT JOIN catalogo.tb_estado_cliente ec ON ec.PK_estado_cliente_id = c.FK_estado_cliente_id
	LEFT JOIN catalogo.tb_tipo_persona tp ON tp.PK_tipo_persona_id = c.FK_tipo_persona_id
	LEFT JOIN catalogo.tb_actividad_economica ae ON ae.PK_actividad_economica_id = c.FK_actividad_economica_id
	LEFT JOIN catalogo.tb_composicion_institucional ci ON ci.PK_composicion_institucional_id = c.FK_composicion_institucional_id
	LEFT JOIN catalogo.tb_codigo_contraparte cc ON cc.PK_codigo_contraparte_id = c.FK_codigo_contraparte_id
	LEFT JOIN catalogo.tb_indicador_validez iv ON iv.PK_indicador_validez_id = c.FK_indicador_validez_id
	LEFT JOIN catalogo.tb_relacion_cliente_institucion rci ON rci.PK_relacion_cliente_institucion_id = c.FK_relacion_cliente_institucion_id
	LEFT JOIN catalogo.tb_tipo_sociedad ts ON ts.PK_tipo_sociedad_id = c.FK_tipo_sociedad_id
	WHERE c.PK_cliente_id = @id`

	objRegistro := objReptry.objSql.QueryRow(qryCliente, sql.Named("id", idCliente))

	var objCliente entitysql.ClienteSql

	if objErrorLectura := objRegistro.Scan(
		&objCliente.Id,
		&objCliente.IdDocumento,
		&objCliente.Nombre,
		&objCliente.FechaIngreso,
		&objCliente.TipoDocumentoId,
		&objCliente.TipoDocumentoDescripcion,
		&objCliente.EstadoId,
		&objCliente.EstadoDescripcion,
		&objCliente.TipoEntidadId,
		&objCliente.TipoEntidadDescripcion,
		&objCliente.ActividadEconomicaId,
		&objCliente.ActividadEconomicaDescripcion,
		&objCliente.ActividadEconomicaCodigoCIIU,
		&objCliente.ComposicionInstitucionalId,
		&objCliente.ComposicionInstitucionalDescripcion,
		&objCliente.CodigoContraparteId,
		&objCliente.CodigoContraparteDescripcion,
		&objCliente.IndicadorValidezId,
		&objCliente.IndicadorValidezDescripcion,
		&objCliente.RelacionClienteInstitucionId,
		&objCliente.RelacionClienteInstitucionDescripcion,
		&objCliente.TipoSociedadId,
		&objCliente.TipoSociedadDescripcion,
		&objCliente.TipoSociedadIdKYC,
		&objCliente.TipoSociedadDescripcionKYC); objErrorLectura != nil {

		if objErrorLectura == sql.ErrNoRows {
			return nil, &shared.ExcpcnReglaNegocio{Mensaje: CLIENTE_NO_EXISTE + " id_cliente:" + strconv.FormatInt(idCliente, 10)}
		}

		return nil, errors.New(CLIENTE_GENERAL_ERROR + objErrorLectura.Error())
	}

	return mapper.ToClienteEntity(&objCliente), nil
}
