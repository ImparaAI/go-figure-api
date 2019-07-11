package database

import (
	//"database/sql"
)

type Connection struct {
	Val string
}

func GetConnection() (*Connection) {
	return persistentConnection
}

func runMigrations() {

}

var persistentConnection = &Connection{"a value"}