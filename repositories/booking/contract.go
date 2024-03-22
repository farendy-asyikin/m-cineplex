package bookingrepository

import (
	"gorm.io/gorm"
	"main.go/models"
)

type BookingRepository interface {
	CreateBooking(booking models.Booking) (*models.Booking, error)
	GetBookingByID(ID string) (*models.Booking, error)
	UpdateBooking(booking models.Booking) (*models.Booking, error)
	DeleteBookingByID(ID string) error
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *bookingRepository {
	return &bookingRepository{db}
}
