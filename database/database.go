package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"os"
	"path/filepath"
)

var persistentDb *sqlx.DB
var testing = false

func GetDb() *sqlx.DB {
	return persistentDb
}

func Initialize() error {
	var err error
	persistentDb, err = openConnection()

	if err != nil {
		return err
	}

	createDatabase()

	return runMigrations()
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

func openConnection() (*sqlx.DB, error) {
	connStr := "root@tcp(mysql:3306)/?parseTime=true"

	return sqlx.Connect("mysql", connStr)
}

func createDatabase() {
	dbName := getDbName()

	if testing {
		persistentDb.MustExec("DROP DATABASE IF EXISTS " + dbName)
	}

	persistentDb.MustExec("CREATE DATABASE IF NOT EXISTS " + dbName)
	persistentDb.MustExec("USE " + dbName)
}

func getDbName() string {
	if testing {
		return "test"
	} else {
		return "gofigure"
	}
}

func runMigrations() error {
	filename := getSchemaFilename()
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	_, err = persistentDb.Exec(string(file))

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
