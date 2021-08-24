package models

import (
	"database/sql"
	"time"
)

type Route struct {
	ID                  uint64       `json:"-" gorm:"primarykey;unique;index"`
	CreatedAt           time.Time    `json:"-"`
	ProcessedAt         sql.NullTime `json:"-"`
	FromLat             float64      `json:"from_lat" binding:"required"`
	FromLng             float64      `json:"from_lng" binding:"required"`
	ToLat               float64      `json:"to_lat" binding:"required"`
	ToLng               float64      `json:"to_lng" binding:"required"`
	LengthInMeters      uint64       `json:"length_in_meters"`
	TravelTimeInSeconds uint64       `json:"travel_time_in_seconds"`
}

func (r *Route) ToWorkerMsg() WorkerMsg {
	return WorkerMsg{
		RequestID: r.ID,
		FromLat:   r.FromLat,
		FromLng:   r.FromLng,
		ToLat:     r.ToLat,
		ToLng:     r.ToLng,
	}
}

func (r *Route) IsValid() bool {
	return r.ID != 0
}

func (r *Route) HasResult() bool {
	return r.ProcessedAt.Valid
}

type WorkerMsg struct {
	RequestID uint64  `json:"request_id"`
	FromLat   float64 `json:"from_lat"`
	FromLng   float64 `json:"from_lng"`
	ToLat     float64 `json:"to_lat"`
	ToLng     float64 `json:"to_lng"`
}
