package locationservice

import (
	locationrepository "main.go/repositories/location"
	"main.go/schemas"
)

type LocationService interface {
	ListLocation() ([]*schemas.ListLocationResponse, error)
}

type locationService struct {
	locationRepository locationrepository.LocationRepository
}

func NewLocationService(
	locationRepository locationrepository.LocationRepository,
) *locationService {
	return &locationService{
		locationRepository: locationRepository,
	}
}
