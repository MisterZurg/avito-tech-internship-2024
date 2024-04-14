package main

import (
	"fmt"
	"github.com/MisterZurg/avito-tech-internship-2024/config"
	"github.com/MisterZurg/avito-tech-internship-2024/internal/repository"
	"github.com/MisterZurg/avito-tech-internship-2024/internal/service"
	"github.com/labstack/echo/v4"
	"log"

	"github.com/MisterZurg/avito-tech-internship-2024/api"
)

const (
	DBType = "postgres"
)

func setupEchoServer(si api.ServerInterface) *echo.Echo {
	e := echo.New()
	// This implements the pet store interface
	api.RegisterHandlers(e, si)

	return e
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Cannot get Config")
	}
	// TODO CONNECT POSTGRE
	pgDSN := config.GetPostgresConnectionString(DBType, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	pgDB, err := repository.New(pgDSN)
	if err != nil {
		log.Fatalf("error parsing config")
	}

	fmt.Println("Connected tho")

	srv := service.New(pgDB)

	e := setupEchoServer(srv)
	e.Logger.Fatal(e.Start(":8080"))
}
