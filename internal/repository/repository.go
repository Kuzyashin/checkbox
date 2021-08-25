package repository

import (
	"github.com/AKuzyashin/checkbox/internal/models"
)

type RoutesDBRepo interface {
	CreateRequest(request *models.Route) error
	SaveResult(result *models.Route) error
	GetResult(requestID uint64) (*models.Route, error)
	GetUnprocessed() []*models.Route
}
