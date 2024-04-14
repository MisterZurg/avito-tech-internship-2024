package service

import (
	"github.com/MisterZurg/avito-tech-internship-2024/api"
	"github.com/MisterZurg/avito-tech-internship-2024/internal/repository"
)

type Repository interface {
	api.ServerInterface
}

type Service struct {
	*repository.Repository
}

func New(repos *repository.Repository) *Service {
	return &Service{repos}
}
