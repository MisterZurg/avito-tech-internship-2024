package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"

	"github.com/MisterZurg/avito-tech-internship-2024/api"
	"github.com/MisterZurg/avito-tech-internship-2024/config"
	"github.com/MisterZurg/avito-tech-internship-2024/internal/repository"
	"github.com/MisterZurg/avito-tech-internship-2024/internal/service"
)

const (
	DBType = "postgres"
)

func setupEchoServer(si api.ServerInterface) *echo.Echo {
	e := echo.New()
	// This implements the pet store interface.
	api.RegisterHandlers(e, si)

	return e
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Cannot get Config")
	}

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
