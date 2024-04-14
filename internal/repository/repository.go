package repository

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	*sqlx.DB
	*BannersRepository
}

func New(dsn string) (*Repository, error) {
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatal("Cannot connect to Postgres")
	}

	return &Repository{
		DB:                db,
		BannersRepository: &BannersRepository{db},
	}, nil
}
