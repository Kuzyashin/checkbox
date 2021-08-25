package services

import (
	"github.com/AKuzyashin/checkbox/internal/config"
	"github.com/AKuzyashin/checkbox/internal/repository"
	"github.com/AKuzyashin/checkbox/internal/routes"
	"github.com/AKuzyashin/checkbox/pkg/tomtom"
	"github.com/rs/zerolog"
)

type Services struct {
	logger zerolog.Logger
	TomTom tomtom.Client
	Routes repository.RoutesRepo
}

func NewServices(cfg *config.AppConfig, db repository.DatabaseRepo) *Services {
	tomTom := tomtom.NewClient(cfg.TomTom.ApiKey)
	rs := routes.NewRoutes(db)
	return &Services{TomTom: tomTom, Routes: rs}
}
