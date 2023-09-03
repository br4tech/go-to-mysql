package adapter

import (
	"database/sql"
)

type MySQLAdapter struct {
	db *sql.DB
}

func NewMySQLAdapter(db *sql.DB) *MySQLAdapter {
	return &MySQLAdapter{db: db}
}

func (a *MySQLAdapter) GetDB() *sql.DB{
	return a.db
}
