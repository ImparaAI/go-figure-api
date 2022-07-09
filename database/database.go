package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"cloud.google.com/go/datastore"
	"os"
	"time"
	"context"
)

var datastoreClient *datastore.Client
var persistentDb *sqlx.DB
var testing = false

func GetDatastore() *datastore.Client {
	return datastoreClient
}

func GetDb() *sqlx.DB {
	return persistentDb
}

func Initialize() error {
	if os.Getenv("DB_STORE") == "google_datastore" {
		var err error
		ctx := context.Background()
		// TODO: set project id in envar
		project := os.Getenv("DATASTORE_PROJECT")
		datastoreClient, err = datastore.NewClient(ctx, project)

		return err
	} else {
		err := createDatabase()

		if err != nil {
			return err
		}

		persistentDb, err = openConnection(getDbName())

		return err
	}
}

func SetTestingEnvironment() {
	testing = true
}

func Close() {
	if os.Getenv("DB_STORE") == "google_datastore" {
		datastoreClient.Close()
	} else {
		persistentDb.Close()
	}
}

func ClearTestingDb() {
	if testing {
		Initialize()
	}
}

func openConnection(databaseName string) (*sqlx.DB, error) {
	password := os.Getenv("DB_PASSWORD")
	connStr := "root:" + password + "@tcp(mysql:3306)/" + databaseName + "?parseTime=true"

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
