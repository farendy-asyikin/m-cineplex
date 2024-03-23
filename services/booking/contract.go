package bookingservice

import (
	"main.go/models"
	bookingrepository "main.go/repositories/booking"
	"main.go/schemas"
)

type BookingService interface {
	CreateBooking(request schemas.CreateBookingRequest) (*models.Booking, error)
	GetBookingByID(ID string) (*models.Booking, error)
}

type bookingService struct {
	bookingRepository bookingrepository.BookingRepository
}

func NewBookingService(
	bookingRepository bookingrepository.BookingRepository,
) *bookingService {
	return &bookingService{
		bookingRepository: bookingRepository,
	}
}
