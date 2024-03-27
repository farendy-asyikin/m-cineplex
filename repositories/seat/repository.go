package seatrepository

import "main.go/models"

func (r *seatRepository) ListUnbookedSeatByFilmIDAndLocationID(filmID, locationID uint64) ([]*models.Seat, error) {
	var unbookedSeats []*models.Seat

	err := r.db.
		Joins("JOIN films ON seats.location_id = films.location_id").
		Where("seats.location_id = ? AND films.id = ? AND seats.is_booked = ?", locationID, filmID, false).
		Find(&unbookedSeats).Error
	if err != nil {
		return nil, err
	}

	return unbookedSeats, nil
}
