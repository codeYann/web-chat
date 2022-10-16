// Package database deals with database connection and its dependencies.
package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/codeYann/web-chat/settings"

	// It Uses blank space to recognize the Postgres driver.
	_ "github.com/lib/pq"
)

var dbSettings *settings.DBConfig = settings.ExportDBConfig()

// OpenConnection returns a database connection.
func OpenConnection() (*sql.DB, error) {
	stringConnection := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbSettings.Host,
		dbSettings.Port,
		dbSettings.User,
		dbSettings.Password,
		dbSettings.Database,
	)

	connection, err := sql.Open("postgres", stringConnection)
	if err != nil {
		log.Fatal("Fail to open database connection", err.Error())
	}

	err = connection.Ping()

	return connection, err
}
