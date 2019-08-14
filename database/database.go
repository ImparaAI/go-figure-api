package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
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
	connStr := "root:password@tcp(mysql:3306)/" + databaseName + "?parseTime=true"

	connection, err := sqlx.Connect("mysql", connStr)

	if err == nil {
		connection.SetConnMaxLifetime(time.Minute*5);
		connection.SetMaxIdleConns(5);
		connection.SetMaxOpenConns(5);
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
	filename := getSchemaFilename()
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	_, err = connection.Exec(string(file))

	return err
}

func getSchemaFilename() string {
	pwd, _ := os.Getwd()

	if testing {

		//todo: recurse up the pwd until you find correct file
		return "/go/src/app/database/schema.sql"
	}

	return filepath.Join(pwd, "database", "schema.sql")
}
