package app

import (
	"wikivin/internal/config"

	mysql "wikivin/pkg/database/MySQL"

	_ "github.com/go-sql-driver/mysql"
)
func Run(path string) {
	cfg, err:= config.Init(path)
	if err != nil{
		// FATAL
	}
	db, err:= mysql.MySQLConnection(cfg.MySQL)
	if err != nil{
		// FATAL
	}
	defer db.Close()

	
}