package app

import (
	"wikivin/internal/config"
	repo "wikivin/internal/repository/mysql"
	"wikivin/internal/server"
	"wikivin/internal/service"
	rest "wikivin/internal/transports/http"

	mysql "wikivin/pkg/database/MySQL"
	"wikivin/pkg/logger"

	_ "github.com/go-sql-driver/mysql"
)
func Run(path string) {
	cfg, err:= config.Init(path)
	if err != nil{
		logger.Error(err)
	}
	db, err:= mysql.MySQLConnection(cfg.MySQL)
	if err != nil{
		logger.Error(err)
	}
	defer db.Close()
	repositories := repo.NewRepositories(db)
	services:= service.NewServices(repositories)
	handler:= rest.NewHandler(services)
	server:= server.NewServer(cfg, handler)
	// go func(){
	// 	if err:= server.Run(); err!= nil{
	// 		logger.Errorf("the server didn't start: %s\n", err)
	// 	}
	// }()

	if err:= server.Run(); err!= nil{
	    logger.Errorf("the server didn't start: %s\n", err)
	}
	
	logger.Info("server started")
}