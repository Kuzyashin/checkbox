package storage

import (
	"fmt"
	"github.com/AKuzyashin/checkbox/internal/config"
	"github.com/AKuzyashin/checkbox/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type postgresDb struct {
	db *gorm.DB
}

func (p *postgresDb) CreateRequest(r *models.Route) error {
	r.CreatedAt = time.Now()
	err := p.db.Model(r).Create(r).Error
	return err
}

func (p *postgresDb) SaveResult(r *models.Route) error {
	err := p.db.Model(r).Where("id", r.ID).
		Updates(
			map[string]interface{}{
				"length_in_meters": r.LengthInMeters,
				"travel_time_in_seconds": r.TravelTimeInSeconds,
				"processed_at": r.ProcessedAt},
			).Error
	return err
}

func (p *postgresDb) GetResult(requestID uint64) (*models.Route, error) {
	var result models.Route
	err := p.db.Model(&result).First(&result, requestID).Error
	return &result, err
}

func (p *postgresDb) GetUnprocessed() []*models.Route {
	var routes []*models.Route
	p.db.Model(&models.Route{}).Where("processed_at ISNULL").Find(&routes)
	return routes
}

func newGORMDb(cfg *config.AppConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Database.Host, cfg.Database.Username, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewPostgres(cfg *config.AppConfig) (*postgresDb, error) {
	db, err := newGORMDb(cfg)
	if err != nil {
		return nil, err
	}
	_ = db.AutoMigrate(&models.Route{})
	return &postgresDb{db: db}, nil
}