package mySQL

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Name     string
}

func (db MySQLConfig) getDatabaseConnectionString() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s", db.Username, db.Password, db.Host, db.Port, db.Name)
}

func MySQLConnection(cfg MySQLConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", cfg.getDatabaseConnectionString())
	if err != nil{
		return nil, err
	}
	return db, err
}