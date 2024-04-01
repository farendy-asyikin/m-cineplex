package bookingrepository

import (
	"errors"
	"gorm.io/gorm"
	"main.go/models"
)

func (r *bookingRepository) CreateBooking(booking models.Booking) (*models.Booking, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&booking).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	bookedSeat := models.BookedSeat{
		SeatID:     booking.SeatID,
		LocationID: booking.LocationID,
	}

	if err := tx.Create(&bookedSeat).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &booking, nil
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
