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

	var seat models.Seat
	if err := tx.Where("id = ? AND is_booked = ?", booking.SeatID, false).First(&seat).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("failed to book, seat is already booked")
	}

	if err := tx.Create(&booking).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(&seat).Update("is_booked", true).Error; err != nil {
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
