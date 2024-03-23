package seatrepository

import "main.go/models"

func (r *seatRepository) ListNotBookedSeat() ([]*models.Seat, error) {
	var readySeat []*models.Seat

	err := r.db.Find(&readySeat).Where("is_booked", false).Error
	if err != nil {
		return nil, err
	}

	return readySeat, nil
}
