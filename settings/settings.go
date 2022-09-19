// Package settings define all needed information, such as:
// Which HTTP port server is running and database configs
package settings

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// APIConfig defines a struct with a single port to use when the HTTP server got executed.
type APIConfig struct {
	Port string
}

// DBConfig defines a struct with needed fields to connect the database.
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type config struct {
	API APIConfig
	DB  DBConfig
}

var cfg *config

// ExportAPIConfig returns api config
func ExportAPIConfig() *APIConfig {
	return &cfg.API
}

// ExportDBConfig returns database config
func ExportDBConfig() *DBConfig {
	return &cfg.DB
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg = new(config)

	cfg.API.Port = os.Getenv("API_PORT")
	cfg.DB.Host = os.Getenv("DB_HOST")
	cfg.DB.Port = os.Getenv("DB_PORT")
	cfg.DB.User = os.Getenv("DB_USER")
	cfg.DB.Password = os.Getenv("DB_PASSWORD")
	cfg.DB.Database = os.Getenv("DB_DATABASE")
}
