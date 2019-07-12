package database

import (
	"os"
	"fmt"
	"io/ioutil"
	"database/sql"
	"path/filepath"
	_ "github.com/mattn/go-sqlite3"
)

var persistentDb *sql.DB
var testing = false

func GetDb() (*sql.DB) {
	return persistentDb
}

func Initialize() error {
	filename := getFilename()

	fmt.Println(filename)

	if testing {
		os.Remove(filename)
	}

	db, err := sql.Open("sqlite3", filename)

	if err != nil {
		return err
	}

	persistentDb = db

	runMigrations()

	return nil
}

func SetTestingEnvironment() {
	testing = true
}

func Close() {
	persistentDb.Close()
}

func runMigrations() {
	pwd, _ := os.Getwd()
	filename := filepath.Join(pwd, "database", "schema.sql")
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Print(err)
	}

	_, err = persistentDb.Exec(string(file))

	fmt.Print(err)
}

func getFilename() string {
	pwd, _ := os.Getwd()

	if testing {
		return filepath.Join(pwd, "database", "sqlite", "test.db")
	} else {
		return filepath.Join(pwd, "database", "sqlite", "gofigure.db")
	}
}