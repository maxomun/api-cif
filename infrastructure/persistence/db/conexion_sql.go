// File: conexion_sql.go
// Author: Oscar Pinto Salazar
// Role: Arquitecto TI Corporativo
// Mail: oscar.pinto@tanner.cl
// Date: 2025-01-09
// Description: Implementación para el uso de base de datos SQL

// Cualquier cambio a la estructura en este archivo, debe ser previa autorizacion del area de arquitectura

package db

import (
	"database/sql"
	"errors"
	"sync"
	"time"

	"api-cif/shared"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const DB_SQL_OK = "conexión a la base de datos cerrada"

var (
	objDb     *sql.DB
	objExcpcn error
	objSync   sync.Once
)

func IniciarConexionDbSql(config shared.Configuracion) (*sql.DB, error) {

	objSync.Do(func() {

		var objError error

		objSql, objErrorCnx := sql.Open("sqlserver", config.BdSql)

		if objErrorCnx != nil {
			objError = objErrorCnx
		}

		// Configurar el pool de conexiones
		objSql.SetMaxOpenConns(config.BdSqlMaxCnxAbierta)                                 // Máximo de conexiones abiertas
		objSql.SetMaxIdleConns(config.BdSqlMaxCnxInactiva)                                // Máximo de conexiones inactivas
		objSql.SetConnMaxLifetime(time.Duration(config.BdSqlMaxTiempoVida) * time.Minute) // Tiempo de vida máximo de las conexiones

		// Asignar la conexión a SQLBoiler
		boil.SetDB(objSql)

		objPingDb := objSql.Ping()

		if objPingDb != nil {
			objError = objPingDb
		}

		objDb = objSql
		objExcpcn = objError
	})

	return objDb, objExcpcn
}

func CerrarConexionDbSql() error {

	if objDb != nil {

		objDb.Close()

		return errors.New(DB_SQL_OK)
	}

	return nil
}
