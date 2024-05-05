package main

import (
	"os"

	mysqlCfg "github.com/go-sql-driver/mysql"

	"github.com/aribroo/go-ecommerce/config"
	"github.com/aribroo/go-ecommerce/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := mysqlCfg.NewConfig()

	cfg.User = config.Envs.DBUser
	cfg.DBName = config.Envs.DBName
	cfg.Passwd = config.Envs.DBPass
	cfg.Addr = config.Envs.DBAddr
	cfg.ParseTime = true
	cfg.Net = "tcp"

	db := database.GetDBConnection(cfg)

	driver, _ := mysql.WithInstance(db, &mysql.Config{})

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		panic(err)
	}

	cmd := os.Args[len(os.Args)-1]

	if cmd == "up" {
		err := m.Up()

		if err != nil && err != migrate.ErrNoChange {

		}
	}

	if cmd == "down" {
		err := m.Down()

		if err != nil && err != migrate.ErrNoChange {

		}
	}
}
