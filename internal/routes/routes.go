package routes

import (
	"github.com/AKuzyashin/checkbox/internal/models"
	"github.com/AKuzyashin/checkbox/internal/repository"
	"time"
)

type routes struct {
	db repository.DatabaseRepo
}

func NewRoutes(db repository.DatabaseRepo) *routes {
	return &routes{db: db}
}

func (rs *routes) CreateRequest(r *models.Route) error {
	r.CreatedAt = time.Now()
	err := rs.db.CreateRequest(r)
	return err
}

func (rs *routes) SaveResult(r *models.Route) error {
	err := rs.db.SaveResult(r)
	return err
}

func (rs *routes) GetResult(requestID uint64) (*models.Route, error) {
	result, err := rs.db.GetResult(requestID)
	return result, err
}

func (rs *routes) GetUnprocessed() []*models.Route {
	result := rs.db.GetUnprocessed()
	return result
}
