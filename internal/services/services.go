package services

import (
	"github.com/AKuzyashin/checkbox/internal/config"
	"github.com/AKuzyashin/checkbox/internal/models"
	"github.com/AKuzyashin/checkbox/internal/repository"
	"github.com/AKuzyashin/checkbox/pkg/tomtom"
	"github.com/rs/zerolog"
)

type RoutesRepo interface {
	CreateRequest(request *models.Route) error
	SaveResult(result *models.Route) error
	GetResult(requestID uint64) (*models.Route, error)
	GetUnprocessed() []*models.Route
}

type Services struct {
	logger zerolog.Logger
	TomTom tomtom.Client
	Routes RoutesRepo
}

func NewServices(cfg *config.AppConfig, db repository.RoutesDBRepo) *Services {
	tomTom := tomtom.NewClient(cfg.TomTom.ApiKey)
	rs := NewRoutes(db)
	return &Services{TomTom: tomTom, Routes: rs}
}
