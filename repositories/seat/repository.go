package seatrepository

import "main.go/models"

func (r *seatRepository) ListUnbookedSeatByFilmIDAndLocationID(filmID, locationID uint64) ([]*models.Seat, error) {
	var unbookedSeats []*models.Seat

	err := r.db.
		Joins("JOIN films ON seats.location_id = films.location_id").
		Joins("LEFT JOIN booked_seats ON seats.id = booked_seats.seat_id").
		Where("seats.location_id = ? AND films.id = ? AND booked_seats.id IS NULL", locationID, filmID).
		Order("seats.id").
		Preload("Location").
		Find(&unbookedSeats).Error
	if err != nil {
		return nil, err
	}

	return unbookedSeats, nil
}
