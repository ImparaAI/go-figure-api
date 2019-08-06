package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
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
	persistentDb, err = openDb()

	if err != nil {
		return err
	}

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

func openDb() (*sqlx.DB, error) {
	filename := getDatabaseFilename()

	if testing {
		os.Remove(filename)
	}

	return sqlx.Connect("sqlite3", filename)
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

func getDatabaseFilename() string {
	pwd, _ := os.Getwd()

	if testing {
		return filepath.Join(os.TempDir(), "test.db")
	}

	return filepath.Join(pwd, "database", "sqlite", "gofigure.db")
}
