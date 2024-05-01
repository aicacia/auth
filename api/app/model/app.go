package model

import "time"

type HealthST struct {
	DB   bool      `json:"db" validate:"required"`
	Date time.Time `json:"date" validate:"required" format:"date-time"`
} // @name Health

func (health HealthST) IsHealthy() bool {
	return health.DB
}

type VersionST struct {
	Version string `json:"version" validate:"required"`
	Build   string `json:"build" validate:"required"`
} // @name Version
