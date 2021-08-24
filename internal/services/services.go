package services

import (
	"github.com/AKuzyashin/checkbox/internal/config"
	"github.com/AKuzyashin/checkbox/internal/repository"
	"github.com/AKuzyashin/checkbox/pkg/tomtom"
	"github.com/rs/zerolog"
)

type Services struct {
	logger zerolog.Logger
	TomTom tomtom.Client
	Database repository.DatabaseRepo
}

func NewServices(cfg *config.AppConfig, db repository.DatabaseRepo) *Services {
	tomTom := tomtom.NewClient(cfg.TomTom.ApiKey)
	return &Services{TomTom: tomTom, Database: db}
}
