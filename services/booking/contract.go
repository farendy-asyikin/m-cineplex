package bookingservice

import (
	"main.go/models"
	bookingrepository "main.go/repositories/booking"
	"main.go/repositories/is_booked"
	"main.go/schemas"
)

type BookingService interface {
	CreateBooking(request schemas.CreateBookingRequest) (*models.Booking, error)
	GetBookingByID(ID string) (*models.Booking, error)
}

type bookingService struct {
	bookingRepository  bookingrepository.BookingRepository
	isBookedRepository is_booked.IsBookedRepository
}

func NewBookingService(
	bookingRepository bookingrepository.BookingRepository,
	isBookedRepository is_booked.IsBookedRepository,
) *bookingService {
	return &bookingService{
		bookingRepository:  bookingRepository,
		isBookedRepository: isBookedRepository,
	}
}
