package db

import "database/sql"

type Database interface {
	Connect() *sql.DB
}
