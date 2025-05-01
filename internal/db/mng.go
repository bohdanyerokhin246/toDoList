package db

import (
	"database/sql"
	"fmt"
)

type Mng struct {
	DB *sql.DB
}

func (m *Mng) Connect() *sql.DB {
	fmt.Println("Mongo connected")
	mng := new(Mng)
	return mng.DB
}
