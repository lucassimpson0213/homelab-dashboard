//   Database repository pattern for better reusability
//   This package helps to abstract database interfaces from their implementation 
//   for looser coupling between components
package database

import (
	"database/sql"
	// "encoding/json"
	// "html/template"
	"log"
	// "net/http"
	_ "path/filepath"

	_ "embed"
	_ "github.com/lib/pq"
)



/*
*   Database repository pattern for better reusability
*
*
*/

//go:embed getresource.sql
var resourcesql string

type Repository struct {
	pg          PgConnection
	resourceSql string
}

func (repository *Repository) New() {
	repository.initConnect()
	repository.resourceSql = resourcesql

}

func (repo *Repository) initConnect() {
	error := repo.pg.getConnection()

	if error != nil {
		log.Fatal(error)
	}

}



/*
*  ################################################
*  Connection specific struct and connection handler
*  ################################################
*
*/



type PgConnection struct {
	connect *sql.DB
}

// TODO make repoistory for database to decouple the database
func (pg *PgConnection) getConnection() error {
	if pg.connect != nil {
		return nil
	}
	db, err := sql.Open("postgres", "host=localhost dbname=postgres connect_timeout=5 sslmode=disable user=lucas password=lucas")
	if err != nil {
		log.Fatal("postgres failed to init")
		return err
	}

	pg.connect = db
	return nil

}
