package locationrepository

import "main.go/models"

func (r *locationRepository) ListLocation() ([]*models.Location, error) {
	var location []*models.Location

	err := r.db.Find(&location).Error
	if err != nil {
		return nil, err
	}

	return location, nil
}
