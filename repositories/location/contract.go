package locationrepository

import (
	"gorm.io/gorm"
	"main.go/models"
)

type LocationRepository interface {
	ListLocation() ([]*models.Location, error)
}

type locationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *locationRepository {
	return &locationRepository{db}
}
