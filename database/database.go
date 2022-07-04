package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
	"fmt"
)

var persistentDb *sqlx.DB
var testing = false

func GetDb() *sqlx.DB {
	return persistentDb
}

func Initialize() error {
	err := createDatabase()

	if err != nil {
		return err
	}

	persistentDb, err = openConnection(getDbName())

	return err
}

func SetTestingEnvironment() {
	testing = true
}

func Close() {
	persistentDb.Close()
}

func ClearTestingDb() {
	if testing {
		Initialize()
	}
}

func openConnection(databaseName string) (*sqlx.DB, error) {
	connection := os.Getenv("DB_CONNECTION")
	password := os.Getenv("DB_PASSWORD")
	connStr := fmt.Sprintf("root:%s@%s/%s?parseTime=true", password, connection, databaseName)

	connection, err := sqlx.Connect("mysql", connStr)

	if err == nil {
		connection.SetConnMaxLifetime(time.Minute * 5)
		connection.SetMaxIdleConns(5)
		connection.SetMaxOpenConns(5)
	}

	return connection, err
}

func createDatabase() error {
	connection, err := openConnection("")

	if err != nil {
		return err
	}

	dbName := getDbName()

	if testing {
		connection.MustExec("DROP DATABASE IF EXISTS " + dbName)
	}

	connection.MustExec("CREATE DATABASE IF NOT EXISTS " + dbName)
	connection.MustExec("USE " + dbName)

	err = runMigrations(connection)

	if err != nil {
		return err
	}

	connection.Close()

	return nil
}

func getDbName() string {
	if testing {
		return "test"
	} else {
		return "gofigure"
	}
}

func runMigrations(connection *sqlx.DB) error {
	_, err := connection.Exec(Schema)

	return err
}
