// File: persona_reptry.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Operaciones con la base de datos relacional

package dbsql

import (
	"api-plantilla/domain/persona"
	entitysql "api-plantilla/infrastructure/persistence/dbsql/entity_sql"
	"api-plantilla/infrastructure/persistence/dbsql/mapper"
	"api-plantilla/shared"

	"strconv"

	"database/sql"
	"errors"

	_ "github.com/denisenkom/go-mssqldb"
)

const PERS_NO_EXISTE = "Sin informacion en la base de datos"
const PERS_GUARDAR_ERROR = "Error al guardar la Persona"
const PERS_GENERAL_ERROR = "Error al ejecutar la sentencia en la base de datos"

const PERS_GUARDAR_OK = "Registro actualizado en forma correcta"

type PersonaReptry struct {
	objSql *sql.DB
}

func NewPersonaRepository(conexion *sql.DB) (*PersonaReptry, error) {
	return &PersonaReptry{objSql: conexion}, nil
}

func (objReptry *PersonaReptry) Actualizar(objeto *persona.Persona) (bool, error) {

	qryPersona := `UPDATE Persona SET Nombre = @Nombre, ApPaterno = @ApPaterno WHERE id = @Id`

	result, objErrorQuery := objReptry.objSql.Exec(qryPersona,
		sql.Named("Nombre", objeto.Nombre),
		sql.Named("ApPaterno", objeto.ApPaterno),
		sql.Named("Id", objeto.Id))

	if objErrorQuery != nil {
		return false, errors.New(PERS_GUARDAR_ERROR + objErrorQuery.Error())
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return false, errors.New(PERS_GUARDAR_ERROR + err.Error())
	}

	if rowsAffected > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (objReptry *PersonaReptry) BuscarPorId(idPersona int) (*persona.Persona, error) {

	qryPersona := `SELECT id,  
	               nombre, apPaterno, apMaterno, mail, numCelular, edad, renta, impuesto, fecharegistro
				   FROM Persona WHERE id = @p1`

	objRegistro := objReptry.objSql.QueryRow(qryPersona, idPersona)

	var objPersona entitysql.PersonaSql

	if objErrorLectura := objRegistro.Scan(&objPersona.Id,
		&objPersona.Nombre,
		&objPersona.ApPaterno,
		&objPersona.ApMaterno,
		&objPersona.Mail,
		&objPersona.NumCelular,
		&objPersona.Edad,
		&objPersona.Renta,
		&objPersona.Impuesto,
		&objPersona.FechaRegistro); objErrorLectura != nil {

		if objErrorLectura == sql.ErrNoRows {
			return nil, &shared.ExcpcnReglaNegocio{Mensaje: PERS_NO_EXISTE + " rut_cliente:" + strconv.Itoa(idPersona)}
		}

		return nil, errors.New(PERS_GENERAL_ERROR + objErrorLectura.Error())
	}

	return mapper.ToEntity(&objPersona), nil
}

func (objReptry *PersonaReptry) Crear(objeto *persona.Persona) (int, error) {

	var intIdCliente int

	qryPersona := `INSERT INTO PERSONA (Nombre, ApPaterno, ApMaterno, Mail, NumCelular, Edad, Renta, Impuesto, FechaRegistro)
	               VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9); SELECT SCOPE_IDENTITY();`

	objErrorQuery := objReptry.objSql.QueryRow(qryPersona,
		objeto.Nombre,
		objeto.ApPaterno,
		objeto.ApMaterno,
		objeto.Mail,
		objeto.NumCelular,
		objeto.Edad,
		objeto.Renta,
		objeto.Impuesto,
		objeto.FechaRegistro).Scan(&intIdCliente)

	if objErrorQuery != nil {
		return -1, errors.New(PERS_GUARDAR_ERROR + objErrorQuery.Error())
	}

	return intIdCliente, nil
}

func (objReptry *PersonaReptry) Listar() ([]*persona.Persona, error) {

	qryPersona := "Select id, nombre, apPaterno, apMaterno, mail, numCelular, edad, renta, impuesto From MinimalAPIBD.dbo.Persona"

	objFilas, objErrorQuery := objReptry.objSql.Query(qryPersona)

	if objErrorQuery != nil {
		return nil, objErrorQuery
	}

	defer objFilas.Close()

	var personas []*persona.Persona

	for objFilas.Next() {

		var objPersona entitysql.PersonaSql

		if objErrorFila := objFilas.Scan(&objPersona.Id,
			&objPersona.Nombre,
			&objPersona.ApPaterno,
			&objPersona.ApMaterno,
			&objPersona.Mail,
			&objPersona.NumCelular,
			&objPersona.Edad,
			&objPersona.Renta,
			&objPersona.Impuesto); objErrorFila != nil {

			return nil, errors.New(objErrorFila.Error())
		}

		personas = append(personas, mapper.ToEntity(&objPersona))
	}

	if len(personas) == 0 {
		return nil, &shared.ExcpcnReglaNegocio{Mensaje: PERS_NO_EXISTE}
	}

	return personas, nil
}
