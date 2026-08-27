// Database repository pattern for better reusability
// This package helps to abstract database interfaces from their implementation
// for looser coupling between components
package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "embed"
	"log/slog"
	_ "path/filepath"

	_ "github.com/lib/pq"
)

/*
*   Database repository pattern for better reusability
*
*
 */


type Repository struct {
	pg          PgConnection
	resourceSql string
}

type Resource struct {
	Url          string `json:"url"`
	ResourceId   string `json:"resourceId"`
	ResourceName string `json:"resourceName"`
}


//global variables to initialize module level state
//go:embed getresource.sql
var resourcesql string
//go:embed postresource.sql
var postresourcesql string 
var repository *Repository = &Repository{}

func New() (*Repository, error) {
	err := repository.initConnect()

	if err != nil {
		slog.Error("Database Initialization failed", "err", err)
		return nil, fmt.Errorf("Database Initialization failed: %w", err)
	}

	repository.resourceSql = resourcesql
	return repository, nil
}

func (repo *Repository) initConnect() error {
	err := repo.pg.getConnection()

	if err != nil {
		slog.Error("Database Init failed:", "err", err)
		return err
	}
	return nil
}

func (repo *Repository) CreateResource(resource Resource) (sql.Result, error) {
    //TODO   
	result, err := repo.pg.connect.Exec(postresourcesql, resource.ResourceName, resource.Url)
    
	if err != nil {
		return nil, fmt.Errorf("error on insert statement for createResource(): ", "err", err)
	}

	return result, nil
}

func (repo *Repository) GetResources() ([]Resource, error) {
	db := repo.pg.connect
	sql := repo.resourceSql
	rows, error := db.Query(sql)

	if error != nil {
		errStr := "Database Query failed: " + sql + " %w"
	    return nil , fmt.Errorf(errStr, error)
		
	}

	defer rows.Close()

	resources := make([]Resource, 10)
	// https://pkg.go.dev/database/sql#DB.Query
	for rows.Next() {
		var resource_id, resource_name, url string
		if err := rows.Scan(&resource_id, &resource_name, &url); err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}
		resource := Resource{ResourceId: resource_id, ResourceName: resource_name, Url: url}

		resources = append(resources, resource)

	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating through rows in database call: %w", err)
	}

	return resources, nil

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
