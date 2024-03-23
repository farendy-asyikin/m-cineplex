package bookingrepository

import (
	"errors"
	"gorm.io/gorm"
	"main.go/models"
)

func (r *bookingRepository) CreateBooking(booking models.Booking) (*models.Booking, error) {
	return nil, nil
}

func (r *bookingRepository) GetBookingByID(ID string) (*models.Booking, error) {
	var booking models.Booking
	err := r.db.First(&booking, ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("booking not found")
		}

		return nil, err
	}

	return &booking, nil
}
