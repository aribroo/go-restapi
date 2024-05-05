package main

import (
	"github.com/aribroo/go-ecommerce/cmd/api"
	"github.com/aribroo/go-ecommerce/config"
	"github.com/aribroo/go-ecommerce/database"
	"github.com/go-sql-driver/mysql"
)

func main() {
	// database config
	cfg := mysql.NewConfig()

	cfg.User = config.Envs.DBUser
	cfg.DBName = config.Envs.DBName
	cfg.Passwd = config.Envs.DBPass
	cfg.Addr = config.Envs.DBAddr
	cfg.ParseTime = true
	cfg.Net = "tcp"

	// setup database
	db := database.GetDBConnection(cfg)

	// setup server
	appPort := config.Envs.AppPort
	server := api.NewApiServer(":"+appPort, db)

	// running the server
	err := server.Run()
	if err != nil {
		panic(err)
	}

}
