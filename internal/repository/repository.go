package repository

import (
	"github.com/AKuzyashin/checkbox/internal/models"
	"github.com/gin-gonic/gin"
)

type DatabaseRepo interface {
	CreateRequest(request *models.Route) error
	SaveResult(result *models.Route) error
	GetResult(requestID uint64) (*models.Route, error)
	GetUnprocessed() []*models.Route
}

type HandlersRepo interface {
	CreateRequest(c *gin.Context)
	GetResult(c *gin.Context)
	Register(app *gin.Engine)
}

type RoutesRepo interface {
	CreateRequest(request *models.Route) error
	SaveResult(result *models.Route) error
	GetResult(requestID uint64) (*models.Route, error)
	GetUnprocessed() []*models.Route
}
